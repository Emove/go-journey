# 如何优雅的关闭Go channel

[原文地址](https://go101.org/article/channel-closing.html)

[译文地址](https://www.jianshu.com/p/d24dfbb33781?from=singlemessage)

几天前，我写了一篇文章来说明Golang中channel的使用规范，在reddit和HN，那篇文章收到了很多赞同，但是我也收到了下面几个关于Go channel设计和规范的批评：

1. 在不能更改channel状态的情况下，没有简单普遍的方式来检查channel是否已经关闭了
2. 关闭已经关闭的channel会导致panic，所以在closer（关闭者）不知道channel是否已经关闭的情况下去关闭channel是很危险的
3. 发送值到已经关闭的channel会导致panic，所以如果sender（发送者）在不知道channel是否已经关闭的情况下去向channel发送值是很危险的

那些批评看起来都很有道理（实际上并没有）。是的，没有一个内置函数可以检查一个channel是否已经关闭。如果你能确定不会向channel发送任何值，那么也确实需要一个简单的方法来检查channel是否已经关闭：

```go
package main

import "fmt"

type T int

func IsClosed(ch <-chan T) bool {
    select {
    case <-ch:
        return true
    default:
    }
    
    return false
}

func main() {
    c := make(chan T)
    fmt.Println(IsClosed(c)) // false
    close(c)
    fmt.Println(IsClosed(c)) // true
}
```

上面已经提到了，没有一种适用的方式来检查channel是否已经关闭了。但是，就算有一个简单的*closed(chan T) bool*函数来检查channel是否已经关闭，他的用处还是很有限的，就像内置的len函数用来检查缓冲channel中元素数量一样。原因就在于，已经检查过的channel的状态有可能在调用了类似的方法返回之后就修改了，因此返回来的值已经不能够正确反映刚才检查的channel的当前状态了。

尽管在调用*closed(ch)*返回true的情况下停止向channel发送值是可以的，但是如果调用*closed(ch)*返回false，那么关闭channel或者继续向channel发送值就不安全了（会panic）。

## The Channel Closing Principle

在使用Go channel的时候，一个适用的原则是*不要从接收端关闭channel，也不要关闭有多个并发发送者的channel*。换句话说，如果说，如果sender（发送者）只是唯一的sender或者是channel最后一个活跃的sender，那么你应该在sender的goroutine关闭channel，从而通知receivers（接收者们）已经没有值可以读了。维持这条原则将保证永远不会发送向一个已经关闭的channel发送值或者关闭一个已经关闭的channel。（下面，我们将会称上面的原则为channel closing principle）

## 打破channel closing principle的解决方案

如果你因为某种原因从接收端（receiver side）关闭channel或者在多个发送者中的一个关闭channel，那么你应该使用列在Golang panic/recover Use Cases的函数来安全地发送值到channel中，（假设channel的元素类型是T）

```go
func SafeSend(ch chan, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			// the return result can be altered
			// in a defer function call
			closed = true
		}
	}()
	
	ch <- value // panic if ch is closed
	return false // <=>closed = false; return
}
```

如果channel没有被关闭的话，那么这个函数的性能将和*ch <- value*接近，对于channel关闭的时候，*SafeSend*函数只会在每个sender goroutine中调用一次，因此程序不会有太大的性能损失。

同样的想法也可以用在从多个goroutine关闭channel中：

```go
func SafeClose(ch chan T) (justClosed bool) {
    defer func() {
        if recover() != nil {
            justClosed = false
        }
    }()
    
    // assume ch != nil here.
    close(ch) // panic if ch is closed
    return true
}
```

很多人喜欢用*sync.Once*来关闭channel：

```go
type MyChannel struct {
    C chan T
    once sync.Once
}

func NewMyChannel() *MyChannel {
    return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
    mc.once.Do(func() {
        close(mc.C))
    })
}
```

当然了，我们也可以用*sync.Mutex*来避免多次关闭channel：

```go
type MyChannel struct {
    C chan T
    closed bool
    mutex sync.Mutex
}

func NewMyChannel() *MyChannel {
    return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
    mc.mutex.Lock()
    defer mc.mutex.Unlock()
    if !mc.closed{
        close(mc.C)
        mc.closed = true
    }
}

func (mc *MyChannel) IsClosed() bool {
    mc.mutex.Lock()
    defer mc.mutex.Unlock()
    return mc.closed
}
```

我们应该要理解为什么Go不支持内置*SafeSend*和*SafeClose*函数，原因就在于并不推荐从接收端或者多个并发发送端关闭channel。Golang甚至禁止关闭只接收（receive-only）的channel。

## 保持channel closing principle的优雅方案

上面的*SafeSend*函数有一个缺点是，在select语句的*case*关键字后不能作为发送操作被调用（译者注：类似于*case SafeSend(ch, t)*）。另外一个缺点是，很多人，包括我自己都觉得上面用过使用*panic/recover*和*sync*包的方案不够优雅。针对各种场景，下面介绍不用使用*panic/recover*和*sync*包，纯粹是利用channel的解决方案。（在下面的例子中，*sync.WaitGroup*只是用来让例子完整的。它的使用在实践中不一定一直都有效）

### 1、M个Receiver，一个Sender

Sender通过关闭data channel说“不再发送”，这是最简单的场景了，就只是当Sender不想再发送的时候让Sender关闭data来关闭channel ：

```go
import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const MaxRandomNumber = 100000
const NumOfReceivers = 100

var wgReceivers = sync.WaitGroup{}

func MultiReceiverTest() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	wgReceivers.Add(NumOfReceivers)

	dataCh := make(chan int, 100)

	// the sender
	go func() {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0{
				// the only sender can close the channel safely
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()
	
	// receivers
	for i := 0; i < NumOfReceivers; i++ {
		go func() {
			defer wgReceivers.Done()
			
			// receive values until dataCh is closed and
			// the value buffer queue of dataCh is empty
			for value := range dataCh{
				log.Println(value)
			}
		}()
	}
	
	wgReceivers.Wait()
}
```

### 2、一个Receiver，N个Sender

Receiver通过关闭一个额外的signal channel说“请停止发送”，这种场景比上一个要复杂点。我们不能让receiver关闭data channel，因为这么做将会打破channel closing priciple。但是我们可以让receiver关闭一个额外的signal channel来通知Sender停止发送：

```go
import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const MaxRandomNumber = 100000
const NumOfReceivers = 100

var wgReceivers = sync.WaitGroup{}

func MultiSenderTest() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	wgReceivers.Add(1)

	dataCh := make(chan int, NumOfReceivers)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel
	// Its sender is the receiver of channel dataCh
	// Its receiver are the senders of channel dataCh

	// senders
	for i := 0; i < NumOfReceivers; i++ {
		go func(){
			for {
				value := rand.Intn(MaxRandomNumber)
				select {
				case <- stopCh:
					return
				case dataCh <- value:
				}
			}
		}()

		// the receiver
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				if value == MaxRandomNumber -1 {
					// the receiver of the dataCh channel is
					// also the sender if the stopCh channel
					// It is safe to close the stop channel here
					close(stopCh)
					return
				}
				log.Println(value)
			}
		}()
		
		wgReceivers.Wait()
	}
}
```

正如注释说的，对于额外的signal channel来说，它的sender是data channel的receiver。额外的signal channel被它唯一的sender关闭，遵守了channel closing principle。

### 3、M个Receiver，N个Sender

它们当中任意一个通过通知一个moderator（仲裁者）关闭额外的signal channel来说“让我们结束游戏吧”，这是最复杂的场景了。我们不能让任意的Receiver和Sender关闭data channel，也不能让任何一个Receiver通过关闭一个额外的signal channel：

```go
func MultiSenderAndReceiverTest() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const NumOfReceiver = 10
	const NumOfSender = 1000

	wgReceivers.Add(NumOfReceiver)

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	toStop := make(chan string, 1)
	// the channel toStop is used to notify the moderator
	// to close the additional signal channel (stopCh)
	// Its senders are any senders and receivers of dataCh
	// Its receiver is the moderator goroutine shown below

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		// part of the trick used to notify the moderator
		// to close the additional signal channel
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumOfSender; i++ {
		go func(id string) {
			for {
				value := rand.Intn(MaxRandomNumber)
				if value == 0 {
					// here, a trick is used to notify the moderator
					// to close the additional signal channel
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// the first select here is to try to exit
				// the goroutine as early as possible
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumOfReceiver; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// same as senders, the first select here is to
				// try to exit the goroutine as early as possible
				select {
				case <- stopCh:
					return
				default:
				}

				select {
				case <- stopCh:
					return
				case value := <-dataCh:
					if value == MaxRandomNumber - 1{
						// the same trick is used to notify the moderator
						// to close the additional signal channel
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}
					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}

```

在这个例子中，仍然遵守着channel closing principle

请注意channel *toStop*的缓冲大小是1，这是为了避免当*moderator* goroutine准备好之前第一个通知就已经发送了，导致丢失。

- 更多的场景？

  很多的场景的变体基本上是基于上面三种的。举个例子，一个基于最复杂情况的变体可能要求Receivers读取*buffer channel*中剩下的所有值。这个应该很容器处理，所有这篇文章也就不提了。

  尽管上面三种场景不能覆盖所有Go channel的使用场景，但它们是最基础的，实践中的大多数场景都可以分类到哪三种中

## 结论

这里没有一种场景要求你去打破channel closing principle。如果你遇到了这种场景，请思考以下你的设计并重写你的代码
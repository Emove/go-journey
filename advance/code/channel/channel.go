package channel

import (
	"fmt"
	"sync"
	"time"
)

func SendToNilChannel() {
	var ch chan int
	fmt.Printf("%v\n", ch)
	ch <- 1
	fmt.Println("send success")
}

func ReceiveFromNilChannel() {
	var ch chan int
	fmt.Println(<-ch)
}

func SendToClosedChannel() {
	ch := make(chan int, 1)
	close(ch)
	ch <- 1
}

func CloseAClosedChannel() {
	ch := make(chan int, 1)
	close(ch)
	close(ch)
}

func ReceiveFromEmptyChannel() {
	ch := make(chan int, 1)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- 1
	}()
	// read values from a empty channel will be blocked while writing values to channel
	fmt.Println(<-ch)

}

// NoBuffChanTest 没有缓存的通道
func NoBuffChanTest() {
	ch := make(chan int)
	go receive(ch)
	ch <- 1
	fmt.Println("发送成功")
}

func receive(in chan int) {
	for i := range in {
		fmt.Println("接收成功", i)
	}
}

func CloseBlockChannel() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		close(ch)
	}()

	for {
		select {
		case <-ch:
			fmt.Println("receive")
			return
		}
	}
}

// GetDataFromChannel test put data to ch1 then get the data from ch2
func GetDataFromChannel() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("ch1", i)
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	// 通道关闭后会退出for range循环
	for i := range ch2 {
		fmt.Println("ch2", i)
	}
}

// TestSingleTrackChannel 测试单向通道的应用
func TestSingleTrackChannel() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)

}

func IsChannelWaiterSeq() {
	ch := make(chan int, 1)
	seq := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func(id int) {
			wg.Done()
			select {
			case num := <-ch:
				ch <- num
				seq <- id
				break
			}
		}(i)
	}

	wg.Wait()
	ch <- 1
	for {
		select {
		case i := <-seq:
			fmt.Println(i)
		}
	}
}

func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
		time.Sleep(10 * time.Microsecond)
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

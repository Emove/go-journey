## Goroutine
### 1、介绍
Go语言中的goroutine实现了程序员只需要关心任务的创建，由系统将任务分配到CPU上实现并发执行的机制。goroutine的概念类似于线程，但goroutine是由Go的运行时（runtime）调度和管理的。Go程序会只能地将goroutine中的任务合理地分配给每个CPU。Go语言之所以被称为现代化编程语言，就是因为它在语言层面已经内置了调用和上下文切换的机制。
在Go语言编程中你不需要去自己写进程、线程、协，你的技能包里只有一个goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个共routine去执行这个函数就可以了。
### 2、使用goroutine
Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine
一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数
示例：
```go
func main() {
	go sayHello()
	fmt.Println("this is main func")
}

func sayHello() {
	fmt.Println("hello goroutine")
}
```
输出：
```go
this is main func
```
可以看到执行结果只打印了this is main func，并没有打印hello goroutine。为什么呢？
在程序启动时，Go程序就会为main函数创建一个默认的goroutine。
当main函数返回的时候该goroutine就结束了，所有在main函数中启动的goroutine会一同结束，main函数所在的goroutine就像时权力的游戏中的夜王，其他的goroutine都是异鬼，夜王一死他转化的那些异鬼也就全部GG了。
再试试用简单粗暴的```time.Sleep```来实现
```go
func main() {
	go sayHello()
	fmt.Println("this is main func")
	time.Sleep(time.Second)
}

func sayHello() {
	fmt.Println("hello goroutine")
}
```
输出：
```go
this is main func
hello goroutine
```
在这一次的执行结果中可以发现，先打印的this is main func，紧接着打印hello goroutine。
首先为什么先打印this is main func？是因为我们在创建新的goroutine的时候需要花费一些时间，而此时main函数所在的goroutine是继续执行的。
### 3、goroutine与线程
#### 1、可增长的栈
OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB），一个goroutine的栈在其声明周期开始时只有很小的栈（典型情况下2KB），goroutine的栈不是固定大小的，它可以按需增大或缩小，goroutine的栈大小限制可以达到1GB，虽然极少会用到这么大。所以在Go语言中一次创建十万左右的goroutine也是可以的
#### 2、goroutine调度
GPM是Go语言运行时（runtime）层面的实现，是Go语言自己实现的一眼调度系统。区别于操作系统调度OS线程
1. G很好理解，就是个goroutine的，里面除了存放本goroutine信息外，还有与所在P的绑定等信息
2. P管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了回去其他P队列里抢任务
3. M（machine）是Go运行时（runtime）对操作系统内核线程的虚拟，M与内核线程一般是一一映射的关系，一个goroutine最终是放在M上执行的

P与M一般也是一一对应的。它们的关系是：P管理着一组G挂载在M上运行。当一个G长久阻塞在一个M上时，runtime会新建一个M，阻塞G所在的P会把其他的G挂在到新建的M上。当旧的G阻塞完成或者认为其已经死掉时，回收旧的M
P的个数是通过runtime.GOMAXPROCS设定（最大256），Go1.5版本之后默认为物理线程数。在并发量大的时候会增加一些P和M，但不会太多，切换太频繁的话得不偿失

单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的，goroutine则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用一个称为m:n调度的技术（复用/调度m个goroutine到n个OS线程）。其一大特点是goroutine的调度是在用户态下完成的，不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护者一块大的内存池，不直接调度系统的malloc函数（除非内存池需要改变），成本比调度OS线程低的多。另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上，再加上本身goroutine的超轻量，以上种种保证了Go调度方法的性能。

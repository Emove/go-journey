package code

import (
	"fmt"
	"runtime"
	"time"
)

// TestReSchedule 测试重新分配任务
// 会先输出world，在输出sched
func TestReSchedule() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("sched")
	}
	// time.Sleep(time.Second)
}

// TestGoExit 测试退出当前协程
func TestGoExit() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 退出当前协程
			runtime.Goexit()
			// 退出之后下方代码不会输出
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		// 退出之后，下方代码不会输出
		fmt.Println("A")
	}()
	for {

	}
}

// TestSetMaxProcs 测试设置核心数
func TestSetMaxProcs(n int) {
	if n == 0 {
		n = 1
	}
	runtime.GOMAXPROCS(n)
	go a()
	go b()
	time.Sleep(time.Second)
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

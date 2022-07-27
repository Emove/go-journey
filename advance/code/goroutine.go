package code

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	// goroutine结束就登记-1
	defer wg.Done()
	fmt.Println("hello goroutine,", i)
}

// TestMultiGoroutine 测试启动多个goroutine实例
func TestMultiGoroutine() {
	for i := 0; i < 10; i++ {
		wg.Add(i)
		go hello(i)
	}
	wg.Wait()
}

package context

import (
	"context"
	"fmt"
	"time"
)

func DeadlineContext() {
	d := time.Now().Add(1 * time.Second)
	deadline, cancelFunc := context.WithDeadline(context.Background(), d)
	//defer fmt.Println(deadline.Value())
	defer func() {
		cancelFunc()
	}()

	value := 0
	select {
	case <-time.After(1 * time.Second):
		value++
	case val := <-deadline.Done():
		fmt.Println(val)
		fmt.Println(deadline.Err().Error())
		deadline.Value(value)
	}

}

func TimeoutContext() {
	timeout, _ := context.WithTimeout(context.Background(), 3*time.Second)

	go func() {
		for {
			select {
			case value := <-timeout.Done():
				fmt.Println(value)
				return
			default:
				//time.AfterFunc(1 * time.Second, func() {
				//	fmt.Printf("....")
				//})
				time.Sleep(1 * time.Second)
				fmt.Println("...")
			}
		}
	}()
	//time.Sleep(3 * time.Second)
	//cancelFunc()
}

func CancelBeforeDone(control context.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	child, _ := context.WithCancel(ctx)
	ch := make(chan string, 3)
	go func() {
		for {
			select {
			case <-child.Done():
				ch <- "closed by parent"
				return
			default:
			}
		}
	}()

	go func() {
		cancel()
		ch <- "cancel called."
	}()
	for {
		select {
		case log := <-ch:
			fmt.Println(log)
			time.Sleep(100 * time.Millisecond)
		case <-control.Done():
			fmt.Println("timeout.")
			return
		default:

		}
	}
}

func CancelContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们取完需要的整数后调用cancel

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

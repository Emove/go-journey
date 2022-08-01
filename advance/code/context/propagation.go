package context

import (
	"context"
	"fmt"
	"time"
)

func ContextPropagation() {
	cancel, cancelFunc := context.WithCancel(context.TODO())

	for i := 0; i < 10; i++ {

		withCancel, _ := context.WithCancel(cancel)
		go func(i int) {
			for {
				select {
				case <-withCancel.Done():
					fmt.Printf("child goroutine: %v, stop by cancel\n", i)
					return
				default:
				}
			}
		}(i)

	}

	time.Sleep(2 * time.Second)

	cancelFunc()

	time.Sleep(10 * time.Second)
}

func StopParallelGoroutine() {
	cancel, cancelFunc := context.WithCancel(context.TODO())

	for i := 0; i < 3; i++ {
		go func(i int) {
			for {
				select {
				case <-cancel.Done():
					fmt.Printf("goroutine: %v\n", i)
					return
				default:
				}
			}
		}(i)
	}

	time.Sleep(2 * time.Second)

	cancelFunc()

	time.Sleep(10 * time.Second)
}

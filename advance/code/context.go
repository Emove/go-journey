package code

import (
	"context"
	"time"
)

func DeadlineContext() {
	d := time.Now().Add(4 * time.Second)
	deadline, cancelFunc := context.WithDeadline(context.Background(), d)
	//defer fmt.Println(deadline.Value())
	defer cancelFunc()

	value := 0
	select {
	case <-time.After(1 * time.Second):
		value++
	case <-deadline.Done():
		deadline.Value(value)
	}
}

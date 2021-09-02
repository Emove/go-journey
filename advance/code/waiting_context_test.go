package code

import (
	"fmt"
	"testing"
	"time"
)

func TestNewWaitingContext(t *testing.T) {
	context := NewWaitingContext(8 * time.Second)
	go func() {
		time.Sleep(10000 * time.Millisecond)
		context.SetValue("after 1 second")
	}()
	context.Done()
	if context.IsTimeout() {
		fmt.Println("context timeout")
	} else {
		fmt.Println("get context value: ", context.GetValue())
	}
}

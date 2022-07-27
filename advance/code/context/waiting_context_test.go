package context

import (
	"fmt"
	"testing"
	"time"
)

func TestNewWaitingContext(t *testing.T) {
	context := NewWaitingContext(3 * time.Second)
	//context.Done()
	go func() {
		time.Sleep(3000 * time.Millisecond)
		context.SetValue("after 3 second")
	}()
	context.Done()
	if context.IsTimeout() {
		fmt.Println("context timeout")
	} else {
		fmt.Println("get context value: ", context.GetValue())
	}
}

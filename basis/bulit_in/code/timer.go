package code

import (
	"fmt"
	"time"
)

func TimeAfter() {
	time.AfterFunc(time.Millisecond, func() {
		fmt.Println("timeout")
	})

	time.Sleep(10 * time.Millisecond)
}

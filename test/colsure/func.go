package colsure

import (
	"fmt"
	"time"
)

func FuncClosure() {

	var fn func()
	cnt := 0
	fn = func() {
		if cnt < 2 {
			time.AfterFunc(time.Second, fn)
		}
		fmt.Println(cnt)
		cnt++
	}
	time.AfterFunc(time.Second, fn)

	select {}
}

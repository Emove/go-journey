package main

import (
	"advance/code"
	"time"
)

func main() {
	//code.CancelContext()
	//code.DeadlineContext()
	code.TimeoutContext()
	time.Sleep(5 * time.Second)
}

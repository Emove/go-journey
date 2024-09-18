package main

import (
	"fmt"
)

func main() {
	//code.CancelContext()
	//code.DeadlineContext()
	//code.TimeoutContext()
	//code.SetValue()
	//code.LogInfo()
	//time.Sleep(5 * time.Second)
	i := 10
	j := hello(&i)
	fmt.Println(i, j)
}
func hello(i *int) (j int) {
	defer func() {
		j = 19
	}()
	return *i
}

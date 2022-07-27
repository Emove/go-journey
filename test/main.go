package main

import "fmt"

func main() {

	go func() {
		panic("goroutine panic")
	}()

	fmt.Println("main goroutine")

	select {}
}

package main

import "fmt"

const X = 7.0

var x interface{} = X

func main() {
	if y, ok := x.(int); ok {
		fmt.Println(y)
	} else {
		fmt.Println(int(y))
	}
}

package oop

import "fmt"

func Convert() {
	fn := func(x, y int) int {
		return x + y
	}

	var i interface{}
	i = fn

	if f, ok := i.(func(x, y int) int); ok {
		fmt.Printf(" f: %d\n", f(1, 1))
	}

}

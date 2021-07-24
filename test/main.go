package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//fmt.Println(nil == &Test{
	//	Name: "1",
	//})
	//fmt.Println(CatchOrDropCheck(6, 1))
	for i := 0; i < 4; i++ {
		fmt.Println(rand.Intn(2))
	}
}

type Test struct {
	Name string
}

func CatchOrDropCheck(x, y int) bool {
	if x > 2 && x < 19 {
		// x坐标在存储区内
		if y >= 1 && y <= 20 {
			// y坐标在轨道范围内
			return true
		}
	}
	return false
}

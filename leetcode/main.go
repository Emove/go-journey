package main

import (
	"fmt"
	"leetcode/code/spiral_order_29"
)

func main() {
	array1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	order := spiral_order_29.SpiralOrder(array1)
	for index := range order {
		fmt.Printf("%d \t", order[index])
	}
}

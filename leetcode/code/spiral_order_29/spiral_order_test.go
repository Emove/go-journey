package spiral_order_29

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	//[1,2,3,6,9,8,7,4,5]
	array1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	order := SpiralOrder(array1)
	for index := range order {
		fmt.Printf("%d \t", order[index])
	}
}

func TestSpiralOrder1(t *testing.T) {
	// [1,2,3,4,8,12,11,10,9,5,6,7]
	array2 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	order := SpiralOrder(array2)
	for index := range order {
		fmt.Printf("%d \t", order[index])
	}
}

func TestSpiralOrder2(t *testing.T) {
	// 1 2 3 4 8 7 6 5
	array := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	order := SpiralOrder(array)
	for index := range order {
		fmt.Printf("%d \t", order[index])
	}
}

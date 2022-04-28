package _05_sort_array_by_parity

import (
	"fmt"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	fmt.Printf("%v\n", SortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Printf("%v\n", SortArrayByParity([]int{0}))
}

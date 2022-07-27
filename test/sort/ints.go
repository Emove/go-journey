package sort

import (
	"fmt"
	"sort"
)

func Ints() {
	ints := []int{34, 343221, 542, 43}
	//sort.Ints(ints)
	sort.Slice(ints, func(i, j int) bool {
		return ints[i] > ints[j]
	})
	for _, i := range ints {
		fmt.Println(i)
	}
}

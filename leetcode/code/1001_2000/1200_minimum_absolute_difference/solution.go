package _200_minimum_absolute_difference

import (
	"math"
	"sort"
)

func MinimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	min, ans := math.MaxInt32, make([][]int, 0)

	for i := 0; i < len(arr)-1; i++ {
		differ := arr[i+1] - arr[i]
		if differ < min {
			min = differ
			ans = [][]int{{arr[i], arr[i+1]}}
		} else if differ == min {
			ans = append(ans, []int{arr[i], arr[i+1]})
		}
	}
	return ans
}

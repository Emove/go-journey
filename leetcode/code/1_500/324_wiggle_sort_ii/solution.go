package _24_wiggle_sort_ii

import "sort"

func WiggleSort(nums []int) {
	n := len(nums)
	temp := make([]int, n)
	copy(temp, nums)
	sort.Ints(temp)
	mid := (n-1)/2 + 1
	for left, right, index := mid-1, n-1, 0; index < n; index += 2 {
		nums[index] = temp[left]
		if index < n-1 {
			nums[index+1] = temp[right]
		}
		left--
		right--
	}
}

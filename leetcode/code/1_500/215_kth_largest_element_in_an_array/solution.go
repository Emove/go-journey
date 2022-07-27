package _15_kth_largest_element_in_an_array

import (
	"math/rand"
	"time"
)

func FindKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, k int) int {
	rand.Seed(time.Now().UnixNano())
	p := partition(nums, left, right)
	if p == k {
		return nums[p]
	} else if p < k {
		return quickSelect(nums, p+1, right, k)
	}
	return quickSelect(nums, left, p-1, k)
}

func partition(nums []int, left, right int) int {
	pivot := rand.Int()%(right-left+1) + left
	swap := func(a, b int) {
		nums[a], nums[b] = nums[b], nums[a]
	}

	swap(pivot, right)
	index := left
	for i := left; i < right; i++ {
		if nums[i] < nums[right] {
			swap(index, i)
			index++
		}
	}
	swap(index, right)
	return index
}

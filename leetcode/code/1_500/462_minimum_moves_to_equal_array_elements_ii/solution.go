package _62_minimum_moves_to_equal_array_elements_ii

import (
	"math/rand"
	"sort"
	"time"
)

func MinMoves2(nums []int) int {
	sort.Ints(nums)
	cnt := 0
	for left, right := 0, len(nums)-1; left < right; {
		cnt += nums[right] - nums[left]
		left++
		right--
	}
	return cnt
}

func MinMoves2i(nums []int) int {
	mid, cnt := findMidst(nums), 0
	for i := 0; i < len(nums); i++ {
		cnt += abs(nums[i] - mid)
	}
	return cnt
}

func findMidst(nums []int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)/2)
}

func quickSelect(nums []int, left, right, k int) int {
	p := partition(nums, left, right)
	if p == k {
		return nums[p]
	} else if p < k {
		return quickSelect(nums, p+1, right, k)
	}
	return quickSelect(nums, left, p-1, k)
}

func partition(nums []int, left, right int) int {
	pivot := rand.Intn(right-left+1) + left
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

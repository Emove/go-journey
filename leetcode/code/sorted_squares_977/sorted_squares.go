package sorted_squares_977

import "math"

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left := 0
	right := len(nums) - 1
	index := right
	for left <= right {
		base := 0
		if math.Abs(float64(nums[left])) >= math.Abs(float64(nums[right])) {
			base = nums[left]
			left++
		} else {
			base = nums[right]
			right--
		}
		res[index] = base * base
		index--
	}
	return res
}

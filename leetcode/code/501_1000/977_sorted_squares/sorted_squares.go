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

func sortedSquares2(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	left := 0
	right := length - 1
	index := right
	l := math.Abs(float64(nums[left]))
	r := math.Abs(float64(nums[right]))
	for left <= right {
		base := 0
		if l >= r {
			base = nums[left]
			if left < right-1 {
				left++
				l = math.Abs(float64(nums[left]))
			}
		} else {
			base = nums[right]
			if right > left {

			}
			right--
			r = math.Abs(float64(nums[right]))
		}
		res[index] = base * base
		index--
	}
	return res
}

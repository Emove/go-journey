package _56_next_greater_element_iii

import (
	"math"
	"strconv"
)

func NextGreaterElement(args int) int {
	nums := []byte(strconv.Itoa(args))
	n := len(nums)
	i := n - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	if i < 0 {
		return -1
	}

	j := n - 1
	for ; j >= i && nums[j] <= nums[i]; j-- {

	}
	nums[i], nums[j] = nums[j], nums[i]

	for l, r := i+1, n-1; l < r; l++ {
		nums[l], nums[r] = nums[r], nums[l]
		r--
	}

	ans, _ := strconv.Atoi(string(nums))
	if ans > math.MaxInt32 {
		return -1
	}
	return ans
}

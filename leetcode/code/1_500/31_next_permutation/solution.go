package _1_next_permutation

// 两遍扫描
func NextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}

	if i >= 0 {
		j := n - 1
		for ; j > i && nums[j] <= nums[i]; j-- {

		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	for l, r := i+1, n-1; l <= r; l++ {
		nums[l], nums[r] = nums[r], nums[l]
		r--
	}

}

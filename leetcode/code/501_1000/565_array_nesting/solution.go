package _65_array_nesting

func ArrayNesting(nums []int) int {
	n, cnt, ans := len(nums), 0, 0
	for i := range nums {
		cnt = 0
		for nums[i] < n {
			i, nums[i] = nums[i], n
			cnt++
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}

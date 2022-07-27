package _98_house_robber

func Rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		temp := second
		second = max(second, first+nums[i])
		first = temp
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package _08_smallest_range_1

func SmallestRangeI(nums []int, k int) int {
	minNum, maxNum := nums[0], nums[0]
	for _, num := range nums[1:] {
		if num < minNum {
			minNum = num
		} else if num > maxNum {
			maxNum = num
		}
	}
	ans := maxNum - minNum - 2*k
	if ans > 0 {
		return ans
	}
	return 0
}

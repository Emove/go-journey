package search_insert_position_35

func searchInsert(nums []int, target int) int {
	left := 0
	ans := len(nums)
	right := ans - 1

	for left <= right {
		mid := ((right - left) >> 1) + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

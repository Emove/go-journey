package _4_find_first_and_last_position_of_element_in_sorted_array

func SearchRange(nums []int, target int) []int {
	l := search(nums, target)
	if l == len(nums) || nums[l] != target {
		return []int{-1, -1}
	}
	r := search(nums, target+1)
	return []int{l, r - 1}
}

func search(nums []int, target int) int {
	l, r, ans := 0, len(nums)-1, len(nums)
	for l <= r {
		mid := l + (r-l)>>1
		if target < nums[mid] || nums[mid] == target {
			r = mid - 1
			ans = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}

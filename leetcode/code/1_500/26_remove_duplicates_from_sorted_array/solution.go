package _6_remove_duplicates_from_sorted_array

func RemoveDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return n
	}
	left, right := 1, 1
	for right < n {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}

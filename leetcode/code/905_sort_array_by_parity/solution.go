package _05_sort_array_by_parity

func SortArrayByParity(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	left := 0
	right := length - 1
	for left < right {
		for nums[left]&1 == 0 && left < right {
			left++
		}
		for nums[right]&1 == 1 && left < right {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}

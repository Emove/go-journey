package rotate_array_189

func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	reverse(nums, 0, length-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, length-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

package move_zeroes_283

func moveZeroes(nums []int) {
	left := 0
	right := 0
	length := len(nums)
	for right < length {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

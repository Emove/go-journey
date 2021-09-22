package binary_search_704

func search(nums []int, target int) int {
	low, height := 0, len(nums)-1
	for low <= height {
		mid := (height-low)/2 + low
		num := nums[mid]
		if target == num {
			return mid
		} else if target > num {
			low = mid + 1
		} else {
			height = mid - 1
		}
	}
	return -1
}

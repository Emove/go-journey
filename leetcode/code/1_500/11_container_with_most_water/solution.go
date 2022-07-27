package _1_container_with_most_water

func MaxArea(height []int) (ans int) {
	for left, right := 0, len(height)-1; left < right; {
		width := right - left
		ans = max(ans, min(height[left], height[right])*width)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

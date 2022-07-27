package _21_shortest_distance_to_a_character

func ShortestToChar(s string, c byte) []int {
	length := len(s)
	ans := make([]int, length)
	left, right := 0, 0
	lastC := -1
	for right < length {
		if s[right] == c {
			if lastC == -1 {
				lastC = right
			}
			for left < right {
				ans[left] = min(abs(right-left), abs(left-lastC))
				left++
			}
			ans[right] = 0
			left = right + 1
			lastC = right
		}
		right++
	}
	for left < right {
		ans[left] = abs(left - lastC)
		left++
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

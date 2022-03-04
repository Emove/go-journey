package length_of_longest_substrings_3

func lengthOfLongestSubString(s string) int {
	length := len(s)
	ans := 0
	cmap := make([]int, 128)
	left := 0
	for right := 0; right < length; right++ {
		index := s[right]
		left = max(left, cmap[index])
		ans = max(ans, right-left+1)
		cmap[index] = right + 1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

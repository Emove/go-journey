package _68_binary_gap

func BinaryGap(n int) (ans int) {
	last, i := -1, 0
	for n > 0 {
		if n&1 == 1 {
			if last != -1 {
				ans = max(ans, i-last)
			}
			last = i
		}
		i++
		n >>= 1
	}

	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

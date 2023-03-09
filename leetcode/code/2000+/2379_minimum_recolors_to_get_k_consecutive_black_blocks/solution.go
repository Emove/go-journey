package _379_minimum_recolors_to_get_k_consecutive_black_blocks

func minimumRecolors(blocks string, k int) int {
	ans, whites := k, 0
	for left, right := 0, 0; right < len(blocks); right++ {
		if blocks[right] == 'W' {
			whites++
		}
		if right < k-1 {
			continue
		}
		ans = min(ans, whites)
		if blocks[left] == 'W' {
			whites--
		}
		left++
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

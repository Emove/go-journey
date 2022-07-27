package _26_flip_string_to_monotone_increasing

func MinFlipsMonoIncr(s string) int {
	dp, cnt := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			cnt++
		} else {
			dp = min(dp+1, cnt)
		}
	}
	return dp
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

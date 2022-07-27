package _67_unique_substrings_in_wraparound_string

func FindSubstringInWraproundString(p string) int {
	dp, cnt := make([]int, 26), 0
	for i, ch := range p {
		if i > 0 && (byte(ch)-p[i-1]+26)%26 == 1 {
			cnt++
		} else {
			cnt = 1
		}
		dp[ch-'a'] = max(dp[ch-'a'], cnt)
	}
	ans := 0
	for _, v := range dp {
		ans += v
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

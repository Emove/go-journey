package _42_di_string_match

// DiStringMatch 贪心算法
func DiStringMatch(s string) []int {
	n := len(s)
	min, max := 0, n
	perm := make([]int, n+1)
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			perm[i] = min
			min++
		} else {
			perm[i] = max
			max--
		}
	}
	perm[n] = min
	return perm
}

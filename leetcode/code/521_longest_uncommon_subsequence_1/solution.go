package _21_longest_uncommon_subsequence_1

func FindLUSlength(a string, b string) int {
	if a != b {
		n := len(a)
		m := len(b)
		if n > m {
			return n
		}
		return m
	}
	return -1
}

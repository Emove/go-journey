package kmp

func Kmp(s, p string) int {
	n := len(p)
	if n == 0 {
		return 0
	}
	j, next := 0, getNext(p)
	for i := 0; i < len(s); i++ {
		for j > 0 && s[i] != p[j] {
			j = next[j-1]
		}
		if s[i] == p[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

func getNext(p string) []int {
	n := len(p)
	j, next := 0, make([]int, n)
	for i := 0; i < n; i++ {
		for j > 0 && p[i] != p[j] {
			j = next[j-1]
		}
		if p[i] == p[j] {
			j++
		}
		next[i] = j
	}
	return next
}

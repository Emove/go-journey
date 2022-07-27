package _22_longest_uncommon_subsequence_ii

import "sort"

func FindLUSLength(strs []string) int {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})

	isSpecial := true
	for i, s1 := range strs {
		s1l := len(s1)
		isSpecial = true
		for j, s2 := range strs {
			if i == j {
				continue
			}
			if len(s2) < s1l {
				break
			}
			if isSub(s1, s2) {
				isSpecial = false
				break
			}
		}
		if isSpecial {
			return s1l
		}
	}
	return -1
}

func isSub(a, b string) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	return i == len(a)
}

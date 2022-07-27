package _4_longest_common_prefix

import "strings"

func LongestCommonPrefix(strs []string) string {
	n := len(strs)
	if n < 1 {
		return ""
	}
	prefix := strings.Builder{}
	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		for j := 1; j < n; j++ {
			if i == len(strs[j]) || strs[j][i] != ch {
				goto re
			}
		}
		prefix.WriteByte(ch)
	}
re:
	return prefix.String()
}

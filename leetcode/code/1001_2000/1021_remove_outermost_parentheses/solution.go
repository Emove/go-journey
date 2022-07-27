package _021_remove_outermost_parentheses

import "strings"

func RemoveOuterParentheses(s string) string {
	builder, cnt, begin := strings.Builder{}, 0, 0
	for i, ch := range s {
		if ch == '(' {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			builder.WriteString(s[begin+1 : i])
			begin = i + 1
		}
	}
	return builder.String()
}

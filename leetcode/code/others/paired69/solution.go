package paired69

import "strings"

func Paired69(S string) string {
	// write code here
	ans := strings.Builder{}
	cnt := 0
	prefix := strings.Builder{}
	for _, ch := range S {
		if ch == '6' {
			cnt++
			ans.WriteByte('6')
		} else {
			if cnt > 0 {
				cnt--
			} else {
				prefix.WriteByte('6')
			}
			ans.WriteByte('9')
		}
	}
	for i := 0; i < cnt; i++ {
		ans.WriteString("9")
	}
	return prefix.String() + ans.String()
}

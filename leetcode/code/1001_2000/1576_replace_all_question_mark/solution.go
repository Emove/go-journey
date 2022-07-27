package replace_all_question_mark_1576

func modifyString(s string) string {
	n := len(s)
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		ch := s[i]
		if ch == '?' {
			for c := 'a'; c <= 'c'; c++ {
				if (i > 0 && int32(result[i-1]) == c) || (i < n-1 && c == int32(s[i+1])) {
					continue
				}
				ch = byte(c)
				break
			}
		}
		result[i] = ch
	}
	return string(result)
}

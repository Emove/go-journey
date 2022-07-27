package _90_find_and_replace_pattern

func FindAndReplacePattern(words []string, pattern string) []string {
	ans := make([]string, 0)
	for _, word := range words {
		if match(word, pattern) {
			ans = append(ans, word)
		}
	}
	return ans
}

func match(word, pattern string) bool {
	m1, m2 := make([]rune, 26), make([]rune, 26)
	for i, ch1 := range word {
		ch2 := pattern[i]
		if m1[ch1-'a'] == 0 && m2[ch2-'a'] == 0 {
			m1[ch1-'a'] = rune(ch2)
			m2[ch2-'a'] = ch1
		} else if m1[ch1-'a'] != rune(ch2) || m2[ch2-'a'] != ch1 {
			return false
		}
	}
	return true
}

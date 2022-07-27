package _19_most_common_word

import (
	"strings"
)

func MostCommonWord(paragraph string, banned []string) string {
	bannedMap := make(map[string]struct{}, len(banned))
	for _, bannedWord := range banned {
		bannedMap[bannedWord] = struct{}{}
	}
	cnt := make(map[string]int)
	l := len(paragraph)
	for i := 0; i < l; {
		start := i
		if !isLetter(paragraph[start]) {
			i++
			continue
		}
		for i < l && isLetter(paragraph[i]) {
			i++
		}
		word := strings.ToLower(paragraph[start:i])
		if _, ok := bannedMap[word]; !ok {
			cnt[word] += 1
		}
	}
	m, mcw := 0, ""
	for s, v := range cnt {
		if v > m {
			mcw = s
			m = v
		}
	}
	return mcw
}

func isLetter(ch uint8) bool {
	return (ch >= 65 && ch <= 90) || (ch >= 97 && ch <= 122)
}

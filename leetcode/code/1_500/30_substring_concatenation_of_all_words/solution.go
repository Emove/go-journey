package _0_substring_concatenation_of_all_words

func FindSubstring(s string, words []string) []int {
	m, n, l := len(words), len(words[0]), len(s)
	cnt := make(map[string]int)
	for _, word := range words {
		cnt[word]++
	}
	substringLength := m * n
	ans := make([]int, 0)
LOOP:
	for i := 0; i <= l-substringLength; i++ {
		substring := s[i : i+substringLength]
		subCnt := make(map[string]int)
		for j := 0; j < m; j++ {
			w := substring[j*n : (j*n)+n]
			if cnt[w] == 0 || subCnt[w]+1 > cnt[w] {
				continue LOOP
			}
			subCnt[w]++
		}
		ans = append(ans, i)
	}
	return ans
}

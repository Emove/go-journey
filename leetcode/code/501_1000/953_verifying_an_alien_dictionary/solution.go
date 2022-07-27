package _53_verifying_an_alien_dictionary

func IsAlienSorted(words []string, order string) bool {
	dic := make([]int, 26)
	for i := range order {
		dic[order[i]-'a'] = i
	}
	n := len(words)
	for i := 0; i < n-1; i++ {
		word1, word2 := words[i], words[i+1]
		n1, n2 := len(word1), len(word2)
		for j := 0; j < n1; j++ {
			if j == n2 {
				return false
			}
			order1, order2 := dic[word1[j]-'a'], dic[word2[j]-'a']
			if order1 < order2 {
				break
			} else if order1 > order2 {
				return false
			}
		}
	}
	return true
}

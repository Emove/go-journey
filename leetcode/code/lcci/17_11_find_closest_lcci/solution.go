package _7_11_find_closest_lcci

import (
	"sync"
)

var INDEX map[string][]int
var once sync.Once

func FindClosest(words []string, word1, word2 string) int {
	n := len(words)
	ans, idx1, idx2 := n, -n, n
	for i, word := range words {
		if word == word1 {
			idx1 = i
			ans = min(ans, abs(idx1-idx2))
		} else if word == word2 {
			idx2 = i
			ans = min(ans, abs(idx1-idx2))
		}
	}
	return ans
}

func FindClosest1(words []string, word1, word2 string) int {
	once.Do(func() {
		INDEX = make(map[string][]int)
		for i, word := range words {
			idxs := INDEX[word]
			idxs = append(idxs, i)
			INDEX[word] = idxs
		}
	})
	w1s := INDEX[word1]
	w2s := INDEX[word2]
	ans := len(words)
	for i, j := 0, 0; i < len(w1s) && j < len(w2s); {
		ans = min(ans, abs(w1s[i]-w2s[j]))
		if w1s[i] <= w2s[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

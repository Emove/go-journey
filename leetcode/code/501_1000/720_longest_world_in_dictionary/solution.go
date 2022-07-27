package _20_longest_world_in_dictionary

import (
	trie2 "leetcode/code/alg/trie"
)

func LongestWorld(words []string) string {
	root := trie2.NewTrieNode()
	for _, word := range words {
		root.Insert(word)
	}
	longest := ""
	ln := len(longest)
	for _, word := range words {
		if root.Search(word) {
			wn := len(word)
			if wn > ln || (wn == ln && word < longest) {
				longest = word
				ln = wn
			}
		}
	}
	return longest
}

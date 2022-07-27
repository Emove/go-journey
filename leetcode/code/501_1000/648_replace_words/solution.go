package _48_replace_words

import (
	"strings"
)

func ReplaceWords(dictionary []string, sentence string) string {
	tree := BuildWithWords(dictionary)
	ans := strings.Builder{}
	for _, word := range strings.Split(sentence, " ") {
		if covered, path := tree.Cover(word); covered {
			ans.WriteString(path + " ")
		} else {
			ans.WriteString(word + " ")
		}
	}
	return ans.String()[:ans.Len()-1]
}

// TrieNode 字典树
type TrieNode struct {
	end      bool
	word     string
	children []*TrieNode
}

func BuildWithWords(words []string) (root *TrieNode) {
	root = NewTrieNode()
	for _, word := range words {
		root.Insert(word)
	}
	return root
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make([]*TrieNode, 26),
	}
}

func (root *TrieNode) Insert(word string) {
	node := root
	for _, s := range word {
		i := s - 'a'
		if node.children[i] == nil {
			node.children[i] = NewTrieNode()
		}
		node = node.children[i]
	}
	node.end = true
	node.word = word
}

func (root *TrieNode) Cover(word string) (bool, string) {
	node := root
	path := []rune{}
	for _, s := range word {
		i := s - 'a'
		if node.end {
			return true, string(path)
		}
		if node.children[i] == nil {
			return node.end, string(path)
		}
		path = append(path, s)
		node = node.children[i]
	}
	return true, string(path)
}

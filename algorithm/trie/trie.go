package trie

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

func (root *TrieNode) Search(word string) bool {
	node := root
	for _, s := range word {
		i := s - 'a'
		if node.children[i] == nil || !node.children[i].end {
			return false
		}
		node = node.children[i]
	}
	return true
}

func (root *TrieNode) StartWith(word string) bool {
	node := root
	for _, s := range word {
		i := s - 'a'
		if node.children[i] == nil {
			return false
		}
		node = node.children[i]
	}
	return true
}

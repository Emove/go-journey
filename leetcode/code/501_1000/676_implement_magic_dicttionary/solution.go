package _76_implement_magic_dicttionary

type TrieNode struct {
	end      bool
	children [26]*TrieNode
}

// 字典树
type MagicDictionary struct {
	root *TrieNode
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		root: &TrieNode{},
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		cur := this.root
		for _, c := range word {
			c -= 'a'
			if cur.children[c] == nil {
				cur.children[c] = &TrieNode{}
			}
			cur = cur.children[c]
		}
		cur.end = true
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return query(this.root, searchWord, 1)
}

func query(node *TrieNode, word string, limit int) bool {
	if len(word) == 0 {
		return node.end && limit == 0
	}
	ch := word[0] - 'a'
	if node.children[ch] != nil && query(node.children[ch], word[1:], 0) {
		return true
	}
	if limit == 1 {
		for i, child := range node.children {
			if i != int(ch) && child != nil && query(child, word[1:], 0) {
				return true
			}
		}
	}
	return false
}

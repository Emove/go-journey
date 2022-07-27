package _45_prefix_and_suffix_search

type WordFilter struct {
	pre  *TrieNode
	suff *TrieNode
}

func Constructor(words []string) WordFilter {
	pre, suff := newTrieNode(), newTrieNode()
	for i, word := range words {
		add(pre, word, i, false)
		add(suff, word, i, true)
	}
	return WordFilter{
		pre:  pre,
		suff: suff,
	}
}

func (wf *WordFilter) F(prefix, suffix string) int {
	pre, suff := wf.pre, wf.suff
	for _, ch := range prefix {
		l := ch - 'a'
		if pre.children[l] == nil {
			return -1
		}
		pre = pre.children[l]
	}

	for i := len(suffix) - 1; i >= 0; i-- {
		l := suffix[i] - 'a'
		if suff.children[l] == nil {
			return -1
		}
		suff = suff.children[l]
	}

	i, j := len(pre.idx)-1, len(suff.idx)-1
	for i >= 0 && j >= 0 {
		if pre.idx[i] > suff.idx[j] {
			i--
		} else if pre.idx[i] < suff.idx[j] {
			j--
		} else {
			return pre.idx[i]
		}
	}
	return -1
}

type TrieNode struct {
	idx      []int
	children []*TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		idx:      make([]int, 0),
		children: make([]*TrieNode, 26),
	}
}

func add(node *TrieNode, word string, idx int, isTurn bool) {
	n := len(word)
	if !isTurn {
		for i := 0; i < n; i++ {
			preCh := word[i] - 'a'
			if node.children[preCh] == nil {
				node.children[preCh] = newTrieNode()
			}
			node = node.children[preCh]
			node.idx = append(node.idx, idx)
		}
	} else {
		for i := n - 1; i >= 0; i-- {
			suffCh := word[i] - 'a'
			if node.children[suffCh] == nil {
				node.children[suffCh] = newTrieNode()
			}
			node = node.children[suffCh]
			node.idx = append(node.idx, idx)
		}
	}
}

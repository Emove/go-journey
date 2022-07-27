package _99_falling_squares

// FallingSquares 这个解法中的线段树采取了自顶向下动态扩展节点的做法
func FallingSquares(positions [][]int) []int {
	ans := make([]int, len(positions))
	maxHeight := 0
	for _, pos := range positions {
		maxHeight = max(maxHeight, pos[0]+pos[1]-1)
	}
	root := &Node{start: 0, end: maxHeight, val: 0}
	for i, pos := range positions {
		start, end := pos[0], pos[0]+pos[1]-1
		height := query(root, start, end)
		update(root, start, end, height+pos[1])
		ans[i] = root.val
	}
	return ans
}

type Node struct {
	val              int
	left, right      *Node
	start, end, maxR int
	modified         bool // 当该值被设置为true时，意味着当前节点的值变化了，pushdown时需要将当前节点的值更新到子节点
}

func update(node *Node, start, end int, value int) {
	if start <= node.start && node.end <= end {
		node.modified = true
		node.val = value
		return
	}
	pushdown(node)
	mid := (node.start + node.end) >> 1
	if start <= mid {
		update(node.left, start, end, value)
	}
	if end > mid {
		update(node.right, start, end, value)
	}
	pushup(node)
}

func query(node *Node, start, end int) int {
	if start <= node.start && node.end <= end {
		return node.val
	}
	pushdown(node)
	mid, ans := (node.start+node.end)>>1, 0
	if start <= mid {
		ans = query(node.left, start, end)
	}
	if end > mid {
		ans = max(ans, query(node.right, start, end))
	}
	return ans
}

func pushdown(node *Node) {
	if node.left == nil {
		node.left = &Node{start: node.start, end: (node.start + node.end) >> 1}
	}
	if node.right == nil {
		node.right = &Node{start: (node.start+node.end)>>1 + 1, end: node.end}
	}
	if !node.modified {
		return
	}
	node.left.val, node.left.modified = node.val, true
	node.right.val, node.right.modified = node.val, true
	node.modified = false
}

func pushup(node *Node) {
	node.val = max(node.left.val, node.right.val)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

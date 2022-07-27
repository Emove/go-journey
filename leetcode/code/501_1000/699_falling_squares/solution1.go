package _99_falling_squares

// 线段树
func FallingSquares1(positions [][]int) []int {
	ans := make([]int, len(positions))
	curMax, root := 0, (*Node)(nil)
	for i, pos := range positions {
		start, end := pos[0], pos[0]+pos[1]
		height := query1(root, start, end)
		root = update1(root, start, end, height+pos[1])
		curMax = max(curMax, height+pos[1])
		ans[i] = curMax
	}
	return ans
}

func query1(node *Node, start, end int) int {
	if node == nil || start > node.maxR {
		return 0
	}
	val := 0
	if end > node.start && start < node.end {
		// 存在交集
		val = node.val
	}
	val = max(val, query1(node.left, start, end))
	if end > node.start {
		val = max(val, query1(node.right, start, end))
	}
	return val
}

func update1(node *Node, start, end, val int) *Node {
	if node == nil {
		return &Node{start: start, end: end, val: val, maxR: end}
	}
	if start <= node.start {
		node.left = update1(node.left, start, end, val)
	} else {
		node.right = update1(node.right, start, end, val)
	}
	node.maxR = max(end, node.maxR)
	return node
}

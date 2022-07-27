package _29_N_ary_tree_level_order_traversal

import (
	"leetcode/code/structure"
)

func LevelOrder(root *structure.N_Node) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*structure.N_Node, 1)
	queue[0] = root
	for queue != nil {
		levelNodes := queue
		queue = nil
		nodes := make([]int, len(levelNodes))
		for i, n := range levelNodes {
			nodes[i] = n.Val
			for _, child := range n.Children {
				queue = append(queue, child)
			}
		}
		ans = append(ans, nodes)
	}
	return ans
}

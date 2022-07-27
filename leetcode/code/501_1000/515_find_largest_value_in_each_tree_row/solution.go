package _15_find_largest_value_in_each_tree_row

import (
	"leetcode/code/structure"
	"math"
)

func LargestValues(root *structure.TreeNode) (ans []int) {
	if root == nil {
		return
	}
	queue := []*structure.TreeNode{root}
	for queue != nil {
		q := queue
		queue = nil
		max := math.MinInt32
		for _, node := range q {
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, max)
	}
	return
}

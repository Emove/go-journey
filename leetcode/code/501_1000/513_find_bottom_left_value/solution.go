package _13_find_bottom_left_value

import (
	"leetcode/code/structure"
	"math"
)

func FindBottomLeftValue(root *structure.TreeNode) int {
	deepest := math.MaxInt32

	queue := []*structure.TreeNode{root}

	for queue != nil {
		q := queue
		queue = nil
		for i, node := range q {
			if i == 0 {
				deepest = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return deepest
}

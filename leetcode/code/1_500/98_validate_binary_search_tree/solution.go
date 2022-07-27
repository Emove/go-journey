package _8_validate_binary_search_tree

import (
	"leetcode/code/structure"
	"math"
)

func IsValidBST(root *structure.TreeNode) bool {
	node, mostRight, last := root, (*structure.TreeNode)(nil), int64(math.MinInt32-1)
	for node != nil {
		if node.Left != nil {
			mostRight = node.Left
			for mostRight.Right != nil && mostRight.Right != node {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = node
				node = node.Left
				continue
			}
			mostRight.Right = nil
		}
		if last >= int64(node.Val) {
			return false
		}
		last = int64(node.Val)
		node = node.Right
	}
	return true
}

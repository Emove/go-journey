package _01_symmetric_tree

import (
	"leetcode/code/structure"
)

// 对称二叉树
func IsSymmetric(root *structure.TreeNode) bool {
	//return isMirror(root, root)
	return isMirror2(root)
}

func isMirror(node1, node2 *structure.TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && isMirror(node1.Left, node2.Right) && isMirror(node2.Left, node1.Right)

}

func isMirror2(root *structure.TreeNode) bool {
	queue := []*structure.TreeNode{root, root}

	for len(queue) > 0 {
		node1, node2 := queue[0], queue[1]
		queue = queue[2:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}

		queue = append(queue, node1.Left)
		queue = append(queue, node2.Right)
		queue = append(queue, node1.Right)
		queue = append(queue, node2.Left)
	}

	return true
}

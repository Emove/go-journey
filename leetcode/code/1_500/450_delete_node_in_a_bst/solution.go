package _50_delete_node_in_a_bst

import (
	"leetcode/code/structure"
)

func deleteNode(root *structure.TreeNode, key int) *structure.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		node := root.Left
		for node.Right != nil {
			node = node.Right
		}
		node.Right = root.Right
		return root.Left
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

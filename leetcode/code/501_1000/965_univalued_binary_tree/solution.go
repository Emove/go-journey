package _65_univalued_binary_tree

import "leetcode/code/structure"

func IsUnivalTree(root *structure.TreeNode) bool {
	val, node, stk := root.Val, root, []*structure.TreeNode{root}
	for node != nil || len(stk) > 0 {
		for node != nil {
			stk = append(stk, node)
			node = node.Left
		}
		node = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if node.Val != val {
			return false
		}
		node = node.Right
	}
	return true
}

package _4_06_successor_lcci

import "leetcode/code/structure"

func InorderSuccessor(root, p *structure.TreeNode) *structure.TreeNode {
	node, successor := root, (*structure.TreeNode)(nil)
	if p.Right != nil {
		successor = p.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		return successor
	}
	for node != nil {
		if node.Val > p.Val {
			successor = node
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return successor
}

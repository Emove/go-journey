package _4_binary_tree_inorder_traversal

import (
	"leetcode/code/structure"
)

func InorderTraversal(root *structure.TreeNode) []int {
	res := make([]int, 0)
	var inorder func(node *structure.TreeNode)
	inorder = func(node *structure.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}

func InorderTraversal1(root *structure.TreeNode) []int {
	stack, ans, node := make([]*structure.TreeNode, 0), make([]int, 0), root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		n := len(stack) - 1
		node = stack[n]
		stack = stack[:n]
		ans = append(ans, node.Val)
		node = node.Right
	}
	return ans
}

// InorderTraversal2 Morris中序遍历
func InorderTraversal2(root *structure.TreeNode) []int {
	node := root
	var mostRight *structure.TreeNode
	ans := make([]int, 0)
	for node != nil {
		if node.Left != nil {
			mostRight = node.Left
			for mostRight.Right != nil && mostRight.Right != node {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = node
				node = node.Left
			} else {
				mostRight.Right = nil
				ans = append(ans, node.Val)
				node = node.Right
			}
		} else {
			ans = append(ans, node.Val)
			node = node.Right
		}
	}
	return ans
}

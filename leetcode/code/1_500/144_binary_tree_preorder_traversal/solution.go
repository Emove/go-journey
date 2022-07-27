package _44_binary_tree_preorder_traversal

import (
	"leetcode/code/structure"
)

func PreorderTraversal(root *structure.TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	var dfs func(node *structure.TreeNode)
	dfs = func(node *structure.TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func PreorderTraversal1(root *structure.TreeNode) []int {
	if root == nil {
		return nil
	}
	stack, ans := []*structure.TreeNode{root}, make([]int, 0)
	for len(stack) > 0 {
		n := len(stack) - 1
		node := stack[n]
		ans = append(ans, node.Val)
		stack = stack[:n]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return ans
}

func PreorderTraversal2(root *structure.TreeNode) []int {
	if root == nil {
		return nil
	}
	node, mostRight, ans := root, (*structure.TreeNode)(nil), make([]int, 0)
	for node != nil {
		if node.Left != nil {
			mostRight = node.Left
			for mostRight.Right != nil && mostRight.Right != node {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				ans = append(ans, node.Val)
				mostRight.Right = node
				node = node.Left
			} else {
				mostRight.Right = nil
				node = node.Right
			}
		} else {
			ans = append(ans, node.Val)
			node = node.Right
		}
	}
	return ans
}

package _45_binary_tree_postorder_traversal

import (
	"leetcode/code/structure"
)

func PostorderTraversal(root *structure.TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	var dfs func(node *structure.TreeNode)
	dfs = func(node *structure.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		ans = append(ans, node.Val)
	}
	dfs(root)
	return ans
}

func PostorderTraversal1(root *structure.TreeNode) []int {
	if root == nil {
		return nil
	}
	node, prev, stack, ans := root, (*structure.TreeNode)(nil), make([]*structure.TreeNode, 0), make([]int, 0)
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		n := len(stack) - 1
		node = stack[n]
		stack = stack[:n]
		if node.Right == nil || node.Right == prev {
			ans = append(ans, node.Val)
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}
	return ans
}

func PostorderTraversal2(root *structure.TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	reverseRightNode := func(node *structure.TreeNode) {
		start := len(ans)
		for node != nil {
			ans = append(ans, node.Val)
			node = node.Right
		}
		for left, right := start, len(ans)-1; left < right; {
			ans[left], ans[right] = ans[right], ans[left]
			left++
			right--
		}
	}
	node, mostRight := root, (*structure.TreeNode)(nil)
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
				reverseRightNode(node.Left)
				node = node.Right
			}
		} else {
			node = node.Right
		}
	}
	reverseRightNode(root)
	return ans
}

package _53_two_sum_4_input_is_a_BST

import (
	"leetcode/code/structure"
)

func FindTarget(root *structure.TreeNode, k int) bool {
	if root == nil {
		return false
	}
	m := make(map[int]struct{})
	var dfs func(node *structure.TreeNode) bool
	dfs = func(node *structure.TreeNode) bool {
		if node == nil {
			return false
		}
		if node.Val < k {
			if _, ok := m[node.Val]; ok {
				return true
			}
			temp := k - node.Val
			m[temp] = struct{}{}
		}
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}

// 双指针+中序遍历
func FindTarget2(root *structure.TreeNode, k int) bool {
	left, right := root, root
	leftStk := []*structure.TreeNode{left}
	rightStk := []*structure.TreeNode{right}
	for left.Left != nil {
		leftStk = append(leftStk, left.Left)
		left = left.Left
	}
	for right.Right != nil {
		rightStk = append(rightStk, right.Right)
		right = right.Right
	}

	for left != right {
		sum := left.Val + right.Val
		if sum == k {
			return true
		}
		if sum < k {
			l := len(leftStk) - 1
			left = leftStk[l]
			leftStk = leftStk[:l]
			for node := left.Right; node != nil; node = node.Left {
				leftStk = append(leftStk, node)
			}
		} else {
			r := len(rightStk) - 1
			right = rightStk[r]
			rightStk = rightStk[:r]
			for node := right.Left; node != nil; node = node.Right {
				rightStk = append(rightStk, node)
			}
		}
	}
	return false
}

package _14_flatten_binary_tree_to_linked_list

import "leetcode/code/structure"

func flatten(root *structure.TreeNode) {
	if root == nil {
		return
	}
	stack := []*structure.TreeNode{root}
	prev := (*structure.TreeNode)(nil)
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil {
			prev.Left = nil
			prev.Right = n
		}
		if n.Right != nil {
			stack = append(stack, n.Right)
		}
		if n.Left != nil {
			stack = append(stack, n.Left)
		}
		prev = n
	}
}

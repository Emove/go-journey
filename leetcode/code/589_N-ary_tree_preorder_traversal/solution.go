package _89_N_ary_tree_preorder_traversal

import (
	"leetcode/code/structure"
)

func preorder(root *structure.N_Node) []int {
	var ans []int
	if root == nil {
		return ans
	}
	st := []*structure.N_Node{root}
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		ans = append(ans, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			st = append(st, node.Children[i])
		}
	}
	return ans
}

func dfs(node *structure.N_Node) []int {
	if node == nil {
		return []int{}
	}
	result := []int{node.Val}
	for _, child := range node.Children {
		result = append(result, dfs(child)...)
	}
	return result
}

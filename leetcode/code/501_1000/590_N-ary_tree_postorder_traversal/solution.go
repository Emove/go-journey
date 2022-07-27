package _90_N_ary_tree_postorder_traversal

import "leetcode/code/structure"

func Postorder(root *structure.N_Node) []int {
	return postorderByDfs(root)
}

func postorderByDfs(root *structure.N_Node) []int {
	var ans []int
	var dfs func(node *structure.N_Node)
	dfs = func(node *structure.N_Node) {
		if node == nil {
			return
		}
		for _, n := range node.Children {
			dfs(n)
		}
		ans = append(ans, node.Val)
	}
	dfs(root)
	return ans
}

package _305_all_elements_in_two_binary_search_trees

import (
	"leetcode/code/structure"
)

func GetAllElements(root1, root2 *structure.TreeNode) []int {
	order1 := inOrder(root1)
	order2 := inOrder(root2)
	m, n := len(order1), len(order2)
	p, q := 0, 0
	res := make([]int, m+n)
	for p < m && q < n {
		if order1[p] <= order2[q] {
			res[p+q] = order1[p]
			p++
		} else {
			res[p+q] = order2[q]
			q++
		}
	}
	for p < m {
		res[p+q] = order1[p]
		p++
	}
	for q < n {
		res[p+q] = order2[q]
		q++
	}
	return res
}

func GetAllElements1(root1 *structure.TreeNode, root2 *structure.TreeNode) []int {
	res := make([]int, 0)
	var compareDfs func(node1, node2 *structure.TreeNode)
	compareDfs = func(node1, node2 *structure.TreeNode) {
		if node1.Val <= node2.Val {
			if node2.Left != nil {
				compareDfs(node1, node2.Left)
				res = append(res, node2.Val)
			} else if node1.Left != nil {
				compareDfs(node1.Left, node2)
				res = append(res, node1.Val)
				if node1.Right != nil {
					compareDfs(node1.Right, node2)
				}
			}
		}
	}
	compareDfs(root1, root2)
	return res
}

func inOrder(node *structure.TreeNode) []int {
	res := make([]int, 0)

	var dfs func(node *structure.TreeNode)
	dfs = func(node *structure.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}

	dfs(node)
	return res
}

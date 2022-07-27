package _08_most_frequent_subtree_sum

import (
	"leetcode/code/structure"
)

func FindFrequentTreeSum(root *structure.TreeNode) []int {
	cnt := make(map[int]int)
	max := 0
	var dfs func(node *structure.TreeNode) int
	dfs = func(node *structure.TreeNode) int {
		if node == nil {
			return 0
		}

		sum := node.Val + dfs(node.Left) + dfs(node.Right)

		cnt[sum]++
		count := cnt[sum]
		if count > max {
			max = count
		}
		return sum
	}

	dfs(root)

	ans := make([]int, 0)
	for k, v := range cnt {
		if v == max {
			ans = append(ans, k)
		}
	}
	return ans
}

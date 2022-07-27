package leaf_similar_tree_872

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	vals := []int{}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			vals = append(vals, node.Val)
			return
		}

		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
	}

	dfs(root1)
	val1 := append([]int{}, vals...)
	vals = []int{}
	dfs(root2)

	if len(vals) != len(val1) {
		return false
	}

	for i, v := range val1 {
		if v != vals[i] {
			return false
		}
	}
	return true
}

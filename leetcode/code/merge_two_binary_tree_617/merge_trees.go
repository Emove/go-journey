package merge_two_binary_tree_617

import "leetcode/code/structure"

func mergeTrees(root1 *structure.TreeNode, root2 *structure.TreeNode) *structure.TreeNode {
	return merge(root1, root2)
}

func merge(node1, node2 *structure.TreeNode) *structure.TreeNode {
	if node1 == nil && node2 == nil {
		return nil
	} else if node1 == nil {
		return node2
	} else if node2 == nil {
		return node1
	}

	node := &structure.TreeNode{
		Val: node1.Val + node2.Val,
	}

	node.Left = merge(node1.Left, node2.Left)
	node.Right = merge(node1.Right, node2.Right)
	return node
}

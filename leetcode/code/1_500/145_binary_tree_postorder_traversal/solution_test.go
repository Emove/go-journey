package _45_binary_tree_postorder_traversal

import (
	"fmt"
	"leetcode/code/structure"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	root := buildTreeNode()
	fmt.Printf("%v\n", PostorderTraversal(root))
	fmt.Printf("%v\n", PostorderTraversal1(root))
	fmt.Printf("%v\n", PostorderTraversal2(root))
}

// postorder: 4,5,2,6,7,3,1
func buildTreeNode() *structure.TreeNode {
	left := &structure.TreeNode{
		Val:   2,
		Left:  &structure.TreeNode{Val: 4},
		Right: &structure.TreeNode{Val: 5},
	}
	right := &structure.TreeNode{
		Val:   3,
		Left:  &structure.TreeNode{Val: 6},
		Right: &structure.TreeNode{Val: 7},
	}
	return &structure.TreeNode{
		Val:   1,
		Left:  left,
		Right: right,
	}
}

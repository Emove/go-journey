package _44_binary_tree_preorder_traversal

import (
	"fmt"
	"leetcode/code/structure"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	root := buildTreeNode()
	fmt.Printf("%v\n", PreorderTraversal(root))
	fmt.Printf("%v\n", PreorderTraversal1(root))
	fmt.Printf("%v\n", PreorderTraversal2(root))
}

// preorder: 1,2,4,5,3,6,7
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

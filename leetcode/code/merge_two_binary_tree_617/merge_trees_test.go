package merge_two_binary_tree_617

import (
	"leetcode/code/structure"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	root1 := structure.BuildTreeNode([]int{1, 3, 2, 5})
	root2 := structure.BuildTreeNode([]int{2, 1, 3, -1, 4, -1, 7})
	trees := mergeTrees(root1, root2)
	structure.PrintTreeNode(trees)
}

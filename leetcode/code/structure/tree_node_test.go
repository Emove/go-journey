package structure

import (
	"testing"
)

func TestBuildTreeNode(t *testing.T) {
	node := BuildTreeNode([]int{2, 1, 3, -1, 4, -1, 7})
	PrintTreeNode(node)
}

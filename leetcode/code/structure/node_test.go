package structure

import (
	"testing"
)

func TestBuildNode(t *testing.T) {
	root := BuildNode([]int{1, 2, 3, 4, 5, 6, 7})
	PrintNode(root)
}

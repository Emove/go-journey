package segment_tree

import (
	"fmt"
	"testing"
)

func TestBuildSegmentTree(t *testing.T) {
	tree := BuildSegmentTree([]int{1, 3, 5, 7, 9, 11})
	dfsPrint(tree.root)
	fmt.Println("----------------------------------")

	tree.update(4, 6)
	dfsPrint(tree.root)
	fmt.Println("----------------------------------")

	rangeSum := tree.query(2, 5)
	fmt.Println(rangeSum)

}

func dfsPrint(n *node) {
	if n == nil {
		return
	}
	fmt.Println(n.val)
	dfsPrint(n.left)
	dfsPrint(n.right)
}

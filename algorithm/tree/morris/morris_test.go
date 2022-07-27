package morris

import (
	"fmt"
	"testing"
)

func TestMorrisPre(t *testing.T) {
	node := buildTreeNode()
	pre := MorrisPre(node)
	fmt.Printf("\n%v\n", pre)
}

func TestMorrisIn(t *testing.T) {
	in := MorrisIn(buildTreeNode())
	fmt.Printf("\n%v\n", in)
}

func TestMorrisPost(t *testing.T) {
	post := MorrisPost(buildTreeNode())
	fmt.Printf("%v\n", post)
}

func TestReverseRightNode(t *testing.T) {
	node := &TreeNode{
		val: 1,
		right: &TreeNode{
			val: 2,
			right: &TreeNode{
				val: 3,
			},
		},
	}
	rightNode := ReverseRightNode(node)
	fmt.Printf("%v\n", rightNode)
}

// pre: 1,2,4,5,3,6,7
// in: 4,2,5,1,6,3,7
// post: 4,5,2,6,7,3,1
func buildTreeNode() *TreeNode {
	left := &TreeNode{
		val:   2,
		left:  &TreeNode{val: 4},
		right: &TreeNode{val: 5},
	}
	right := &TreeNode{
		val:   3,
		left:  &TreeNode{val: 6},
		right: &TreeNode{val: 7},
	}
	return &TreeNode{
		val:   1,
		left:  left,
		right: right,
	}
}

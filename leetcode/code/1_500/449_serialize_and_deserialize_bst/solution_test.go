package _49_serialize_and_deserialize_bst

import (
	"leetcode/code/structure"
	"testing"
)

func TestConstructor(t *testing.T) {
	codec := Constructor()
	// nums := []int{41,37,44,24,39,42,48,1,35,38,40,-1,43,46,49,0,2,30,36,-1,-1,-1,-1,-1,-1,45,47,-1,-1,-1,-1,-1,4,29,32,-1,-1,-1,-1,-1,-1,3,9,26,-1,31,34,-1,-1,7,11,25,27,-1,-1,33,-1,6,8,10,16,-1,-1,-1,28,-1,-1,5,-1,-1,-1,-1,-1,15,19,-1,-1,-1,-1,12,-1,18,20,-1,13,17,-1,-1,22,-1,14,-1,-1,21,23}
	serialize := codec.serialize(structure.BuildTreeNode([]int{2, 1, 3}))
	root := codec.deserialize(serialize)
	structure.PrintTreeNode(root)
}

package populating_next_right_pointers_in_each_node_116

import (
	"leetcode/code/structure"
	"testing"
)

func Test_connect(t *testing.T) {
	root := structure.BuildNode([]int{1, 2, 3, 4, 5, 6, 7})
	connect(root)
	structure.PrintNodeNext(root)
}

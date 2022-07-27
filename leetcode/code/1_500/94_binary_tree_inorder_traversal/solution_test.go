package _4_binary_tree_inorder_traversal

import (
	"leetcode/code/structure"
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	type args struct {
		root *structure.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{root: structure.BuildTreeNode([]int{1, -1, 2, -1, -1, 3})},
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InorderTraversal2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

package _8_validate_binary_search_tree

import (
	"leetcode/code/structure"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	type args struct {
		root *structure.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{root: structure.BuildTreeNode([]int{2, 1, 3})},
			want: true,
		},
		{
			name: "second",
			args: args{root: structure.BuildTreeNode([]int{5, 1, 4, -1, -1, 3, 6})},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBST(tt.args.root); got != tt.want {
				t.Errorf("IsValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

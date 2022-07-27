package _01_symmetric_tree

import (
	"leetcode/code/structure"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
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
			args: args{root: structure.BuildTreeNode([]int{1, 2, 2, 3, 4, 4, 3})},
			want: true,
		},
		{
			name: "second",
			args: args{root: structure.BuildTreeNode([]int{1, 2, 2, -1, 3, -1, 3})},
			want: false,
		},
		{
			name: "third",
			args: args{root: structure.BuildTreeNode([]int{1, 2, 2, 2, -1, 2})},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSymmetric(tt.args.root); got != tt.want {
				t.Errorf("IsSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

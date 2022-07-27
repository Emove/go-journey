package _022_sum_of_root_to_leaf_binary_numbers

import (
	"leetcode/code/structure"
	"testing"
)

func TestSumRootToLeaf(t *testing.T) {
	type args struct {
		root *structure.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{root: structure.BuildTreeNode([]int{1, 0, 1, 0, 1, 0, 1})},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumRootToLeaf(tt.args.root); got != tt.want {
				t.Errorf("RemoveOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}

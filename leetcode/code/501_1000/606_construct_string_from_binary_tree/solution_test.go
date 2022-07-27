package _06_construct_string_from_binary_tree

import (
	"leetcode/code/structure"
	"testing"
)

func Test_tree2str(t *testing.T) {
	type args struct {
		root *structure.TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{root: structure.BuildTreeNode([]int{1, 2, 3, 4})},
			want: "1(2(4))(3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree2str(tt.args.root); got != tt.want {
				t.Errorf("tree2str() = %v, want %v", got, tt.want)
			}
		})
	}
}

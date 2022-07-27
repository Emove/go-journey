package _13_find_bottom_left_value

import (
	"leetcode/code/structure"
	"testing"
)

func TestFindBottomLeftValue(t *testing.T) {
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
			args: args{root: structure.BuildTreeNode([]int{2, 1, 3})},
			want: 1,
		},
		{
			name: "second",
			args: args{root: structure.BuildTreeNode([]int{1, 2, 3, 4, -1, 5, 6, -1, -1, 7, 8})},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindBottomLeftValue(tt.args.root); got != tt.want {
				t.Errorf("FindBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

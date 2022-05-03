package _305_all_elements_in_two_binary_search_trees

import (
	"leetcode/code/structure"
	"reflect"
	"testing"
)

func TestGetAllElements(t *testing.T) {
	type args struct {
		root1 *structure.TreeNode
		root2 *structure.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{root1: structure.BuildTreeNode([]int{2, 1, 4}), root2: structure.BuildTreeNode([]int{1, 0, 3})},
			want: []int{0, 1, 1, 2, 3, 4},
		},
		{
			name: "second",
			args: args{root1: structure.BuildTreeNode([]int{1, -1, 8}), root2: structure.BuildTreeNode([]int{8, 1})},
			want: []int{1, 1, 8, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllElements(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

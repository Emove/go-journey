package _53_two_sum_4_input_is_a_BST

import (
	"leetcode/code/structure"
	"testing"
)

func TestFindTarget(t *testing.T) {
	type args struct {
		root *structure.TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{root: structure.BuildTreeNode([]int{5, 3, 6, 2, 4, -1, 7}), k: 9},
			want: true,
		},
		{
			name: "second",
			args: args{root: structure.BuildTreeNode([]int{5, 3, 6, 2, 4, -1, 7}), k: 28},
			want: false,
		},
	}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := FindTarget(tt.args.root, tt.args.k); got != tt.want {
	//			t.Errorf("FindTarget() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTarget2(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("FindTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

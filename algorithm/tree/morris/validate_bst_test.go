package morris

import "testing"

func TestVerifyBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{root: buildBST()},
			want: true,
		},
		{
			name: "second",
			args: args{root: buildNode()},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateBST(tt.args.root); got != tt.want {
				t.Errorf("ValidateBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildBST() *TreeNode {
	left := &TreeNode{
		val:   4,
		left:  &TreeNode{val: 2},
		right: &TreeNode{val: 6},
	}
	right := &TreeNode{
		val:   12,
		left:  &TreeNode{val: 10},
		right: &TreeNode{val: 14},
	}
	return &TreeNode{
		val:   8,
		left:  left,
		right: right,
	}
}

func buildNode() *TreeNode {
	left := &TreeNode{
		val:   4,
		left:  &TreeNode{val: 2},
		right: &TreeNode{val: 9},
	}
	right := &TreeNode{
		val:   12,
		left:  &TreeNode{val: 10},
		right: &TreeNode{val: 14},
	}
	return &TreeNode{
		val:   8,
		left:  left,
		right: right,
	}
}

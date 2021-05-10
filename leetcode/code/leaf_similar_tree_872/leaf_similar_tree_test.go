package leaf_similar_tree_872

import (
	"testing"
)

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{
				root1: toTree([]int{3, 5, 1, 6, 2, 9, 8, -1, -1, 7, 4}),
				root2: toTree([]int{3, 5, 1, 6, 7, 4, 2, -1, -1, -1, -1, -1, -1, 9, 8}),
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				root1: toTree([]int{1}),
				root2: toTree([]int{1}),
			},
			want: true,
		},
		{
			name: "third",
			args: args{
				root1: toTree([]int{1}),
				root2: toTree([]int{2}),
			},
			want: false,
		},
		{
			name: "forth",
			args: args{
				root1: toTree([]int{1, 2}),
				root2: toTree([]int{2, 2}),
			},
			want: true,
		},
		{
			name: "fifth",
			args: args{
				root1: toTree([]int{1, 3, 2}),
				root2: toTree([]int{1, 2, 3}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func toTree(nodeArray []int) *TreeNode {
	length := len(nodeArray)
	if length < 1 {
		return nil
	}
	root := &TreeNode{
		Val: nodeArray[0],
	}
	treeMap := make(map[int]*TreeNode)
	treeMap[0] = root
	for i, value := range nodeArray {
		if value == -1 || i == 0 {
			continue
		}
		// 计算父节点
		var parentNode *TreeNode
		cur := &TreeNode{
			Val: value,
		}
		if i%2 == 1 {
			// left
			parent := (i - 1) / 2
			parentNode = treeMap[parent]
			parentNode.Left = cur
		} else {
			parent := (i - 2) / 2
			parentNode = treeMap[parent]
			parentNode.Right = cur
		}
		treeMap[i] = cur
	}
	return root
}

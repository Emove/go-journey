package _4_06_successor_lcci

import (
	"leetcode/code/structure"
	"reflect"
	"testing"
)

func TestInorderSuccessor(t *testing.T) {
	type args struct {
		root *structure.TreeNode
		p    *structure.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *structure.TreeNode
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InorderSuccessor(tt.args.root, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

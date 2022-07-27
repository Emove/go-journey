package _29_N_ary_tree_level_order_traversal

import (
	"leetcode/code/structure"
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	type args struct {
		root *structure.N_Node
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

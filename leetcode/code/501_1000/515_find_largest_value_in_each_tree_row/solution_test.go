package _15_find_largest_value_in_each_tree_row

import (
	"leetcode/code/structure"
	"reflect"
	"testing"
)

func TestLargestValues(t *testing.T) {
	type args struct {
		root *structure.TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := LargestValues(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("LargestValues() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

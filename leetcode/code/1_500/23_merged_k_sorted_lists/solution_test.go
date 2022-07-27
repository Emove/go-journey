package _3_merged_k_sorted_lists

import (
	"leetcode/code/eq"
	"leetcode/code/structure"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	type args struct {
		lists []*structure.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structure.ListNode
	}{
		{
			name: "first",
			args: args{lists: []*structure.ListNode{structure.BuildListNode([]int{1, 4, 5}), structure.BuildListNode([]int{1, 3, 4}), structure.BuildListNode([]int{2, 6})}},
			want: structure.BuildListNode([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			name: "second",
			args: args{lists: []*structure.ListNode{}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeKLists1(tt.args.lists); !eq.EqualListNode(got, tt.want) {
				t.Errorf("MergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

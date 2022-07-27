package merge_two_sorted_list_21

import (
	"leetcode/code/structure"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *structure.ListNode
		l2 *structure.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structure.ListNode
	}{
		{
			name: "first",
			args: args{l1: structure.BuildListNode([]int{1, 2, 4}), l2: structure.BuildListNode([]int{1, 3, 4})},
			want: structure.BuildListNode([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "second",
			args: args{l1: structure.BuildListNode([]int{}), l2: structure.BuildListNode([]int{})},
			want: structure.BuildListNode([]int{}),
		},
		{
			name: "third",
			args: args{l1: structure.BuildListNode([]int{}), l2: structure.BuildListNode([]int{2})},
			want: structure.BuildListNode([]int{2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.l1, tt.args.l2)
			node := tt.want
			for node != nil && got != nil {
				if node.Val != got.Val {
					t.Errorf("test failed")
				}
				node = node.Next
				got = got.Next
			}
			if node != nil || got != nil {
				t.Errorf("test failed")
			}
		})
	}
}

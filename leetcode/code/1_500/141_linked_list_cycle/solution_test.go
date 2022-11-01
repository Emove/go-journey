package _41_linked_list_cycle

import (
	"leetcode/code/structure"
	"testing"
)

func TestHasCycle(t *testing.T) {
	type args struct {
		head *structure.ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{head: buildCycleListNode([]int{3, 2, 0, 4}, 1)},
			want: true,
		},
		{
			name: "second",
			args: args{head: buildCycleListNode([]int{1, 2}, 0)},
			want: true,
		},
		{
			name: "third",
			args: args{head: buildCycleListNode([]int{1}, -1)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCycle_1(tt.args.head); got != tt.want {
				t.Errorf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildCycleListNode(nums []int, pos int) *structure.ListNode {
	head := structure.BuildListNode(nums)

	if pos == -1 {
		return head
	}
	node := head
	for i := 0; i < pos && node != nil; i++ {
		node = node.Next
	}

	last := node
	for last.Next != nil {
		last = last.Next
	}

	last.Next = node

	return head
}

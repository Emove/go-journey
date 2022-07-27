package _9_ordered_circular_list

import (
	"leetcode/code/eq"
	"leetcode/code/structure"
	"testing"
)

func TestInsert(t *testing.T) {
	type args struct {
		head *structure.ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *structure.ListNode
	}{
		{
			name: "first",
			args: args{head: structure.BuildCircularListNode([]int{3, 4, 1}), x: 2},
			want: structure.BuildCircularListNode([]int{3, 4, 1, 2}),
		},
		{
			name: "second",
			args: args{head: structure.BuildCircularListNode([]int{}), x: 1},
			want: structure.BuildCircularListNode([]int{1}),
		},
		{
			name: "third",
			args: args{head: structure.BuildCircularListNode([]int{1}), x: 0},
			want: structure.BuildCircularListNode([]int{1, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.args.head, tt.args.x); !eq.EqualCircularListNode(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

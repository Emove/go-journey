package remove_nth_node_from_end_of_list_19

import (
	"leetcode/code/structure"
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *structure.ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *structure.ListNode
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{structure.BuildListNode([]int{1, 2, 3, 4, 5}), 2},
			want: structure.BuildListNode([]int{1, 2, 3, 5}),
		},
		{
			name: "second",
			args: args{structure.BuildListNode([]int{1}), 1},
			want: nil,
		},
		{
			name: "third",
			args: args{structure.BuildListNode([]int{1, 2}), 1},
			want: structure.BuildListNode([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthFromEnd(tt.args.head, tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

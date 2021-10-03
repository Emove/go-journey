package reverse_linked_list_206

import (
	"leetcode/code/structure"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *structure.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structure.ListNode
	}{
		{
			name: "first",
			args: args{head: structure.BuildListNode([]int{1, 2, 3, 4, 5})},
			want: structure.BuildListNode([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "second",
			args: args{head: structure.BuildListNode([]int{1, 2})},
			want: structure.BuildListNode([]int{2, 1}),
		},
		{
			name: "third",
			args: args{head: structure.BuildListNode([]int{})},
			want: structure.BuildListNode([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.args.head)
			node := tt.want
			for node != nil && got != nil {
				if node.Val != got.Val {
					t.Fatal("test failed")
				}
				node = node.Next
				got = got.Next
			}
			if node != nil || got != nil {
				t.Fatal("test failed")
			}
		})
	}
}

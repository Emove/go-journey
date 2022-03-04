package middle_of_the_linked_list_876

import (
	"leetcode/code/structure"
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *structure.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structure.ListNode
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{head: structure.BuildListNode([]int{1, 2, 3, 4, 5, 6})},
			want: structure.BuildListNode([]int{4, 5, 6}),
		},
		{
			name: "second",
			args: args{head: structure.BuildListNode([]int{1, 2, 3, 4, 5})},
			want: structure.BuildListNode([]int{3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

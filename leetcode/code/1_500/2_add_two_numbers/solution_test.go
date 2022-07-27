package __add_two_numbers

import (
	"leetcode/code/structure"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
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
			args: args{l1: structure.BuildListNode([]int{2, 4, 3}), l2: structure.BuildListNode([]int{5, 6, 4})},
			want: structure.BuildListNode([]int{7, 0, 8}),
		},
		{
			name: "second",
			args: args{l1: structure.BuildListNode([]int{0}), l2: structure.BuildListNode([]int{0})},
			want: structure.BuildListNode([]int{0}),
		},
		{
			name: "third",
			args: args{l1: structure.BuildListNode([]int{9, 9, 9, 9, 9, 9, 9}), l2: structure.BuildListNode([]int{9, 9, 9, 9})},
			want: structure.BuildListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(list2num(got), list2num(tt.want)) {
				t.Errorf("AddTwoNumbers() = %v, want %v", list2num(got), list2num(tt.want))
			}
		})
	}
}

func list2num(list *structure.ListNode) int {
	budiler := &strings.Builder{}
	for list != nil {
		budiler.WriteString(strconv.Itoa(list.Val))
		list = list.Next
	}
	atoi, _ := strconv.Atoi(budiler.String())
	return atoi
}

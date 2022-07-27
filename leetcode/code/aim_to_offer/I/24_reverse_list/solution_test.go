package _4_reverse_list

import (
	"fmt"
	"leetcode/code/structure"
	"testing"
)

func TestReverseList(t *testing.T) {
	head := &structure.ListNode{
		Val: 1,
		Next: &structure.ListNode{
			Val: 2,
			Next: &structure.ListNode{
				Val: 3,
			},
		},
	}
	list := ReverseList(head)
	for list != nil {
		fmt.Printf("%d\t", list.Val)
		list = list.Next
	}
}

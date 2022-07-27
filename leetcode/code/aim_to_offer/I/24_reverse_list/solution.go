package _4_reverse_list

import "leetcode/code/structure"

func ReverseList(head *structure.ListNode) *structure.ListNode {
	pre, node := (*structure.ListNode)(nil), head
	for node != nil {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}

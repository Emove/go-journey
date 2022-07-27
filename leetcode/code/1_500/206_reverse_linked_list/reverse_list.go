package reverse_linked_list_206

import "leetcode/code/structure"

func reverseList(head *structure.ListNode) *structure.ListNode {
	var prev *structure.ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

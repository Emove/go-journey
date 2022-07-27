package remove_nth_node_from_end_of_list_19

import "leetcode/code/structure"

func removeNthFromEnd(head *structure.ListNode, n int) *structure.ListNode {
	dummy := structure.NewListNode(-1)
	dummy.Next = head
	left := dummy
	right := dummy
	for i := 0; i < n+1; i++ {
		right = right.Next
	}
	for nil != right {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return dummy.Next
}

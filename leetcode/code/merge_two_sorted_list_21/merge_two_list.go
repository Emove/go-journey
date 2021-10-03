package merge_two_sorted_list_21

import "leetcode/code/structure"

func mergeTwoLists(l1 *structure.ListNode, l2 *structure.ListNode) *structure.ListNode {
	head := structure.NewListNode(-1)
	node := head
	for l1 != nil && l2 != nil {
		var val int
		if l1.Val <= l2.Val {
			val = l1.Val
			l1 = l1.Next
		} else {
			val = l2.Val
			l2 = l2.Next
		}
		node.Next = structure.NewListNode(val)
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return head.Next
}

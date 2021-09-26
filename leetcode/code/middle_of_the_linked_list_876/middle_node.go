package middle_of_the_linked_list_876

import "leetcode/code/structure"

func middleNode(head *structure.ListNode) *structure.ListNode {
	left := 1
	right := 1
	node := head
	for node.Next != nil {
		right++
		node = node.Next
	}
	mid := right / 2
	node = head
	for left <= mid {
		node = node.Next
		left++
	}
	return node
}

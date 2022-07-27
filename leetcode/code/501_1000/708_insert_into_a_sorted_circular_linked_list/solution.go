package _08_insert_into_a_sorted_circular_linked_list

import "leetcode/code/structure"

func Insert(head *structure.ListNode, x int) *structure.ListNode {
	target := &structure.ListNode{Val: x}
	if head == nil {
		head = target
		target.Next = head
		return head
	}

	if head.Next == head {
		head.Next = target
		target.Next = head
		return head
	}

	cur, next := head, head.Next
	for next != head {
		if x >= cur.Val && x <= next.Val {
			break
		}
		if cur.Val > next.Val {
			if x >= cur.Val || x <= next.Val {
				break
			}
		}
		cur = next
		next = next.Next
	}
	cur.Next = target
	target.Next = next
	return head
}

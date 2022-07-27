package __add_two_numbers

import "leetcode/code/structure"

func AddTwoNumbers(l1 *structure.ListNode, l2 *structure.ListNode) *structure.ListNode {
	adv, cur := 0, 0
	dummy := &structure.ListNode{}
	node := dummy

	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + adv
		adv, cur = 0, sum
		if cur >= 10 {
			adv, cur = sum/10, sum%10
		}
		node.Next = &structure.ListNode{}
		node = node.Next
		node.Val = cur
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum := l1.Val + adv
		adv, cur = 0, sum
		if cur >= 10 {
			adv, cur = sum/10, sum%10
		}
		node.Next = &structure.ListNode{}
		node = node.Next
		node.Val = cur
		l1 = l1.Next
	}
	for l2 != nil {
		sum := l2.Val + adv
		adv, cur = 0, sum
		if cur >= 10 {
			adv, cur = sum/10, sum%10
		}
		node.Next = &structure.ListNode{}
		node = node.Next
		node.Val = cur
		l2 = l2.Next
	}
	if adv != 0 {
		node.Next = &structure.ListNode{}
		node = node.Next
		node.Val = adv
	}
	return dummy.Next
}

package eq

import "leetcode/code/structure"

func EqualListNode(list1, list2 *structure.ListNode) bool {
	for list1 != nil && list2 != nil {
		if list1.Val != list2.Val {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}
	return list1 == nil && list2 == nil
}

func EqualCircularListNode(list1, list2 *structure.ListNode) bool {
	head1, head2 := list1, list2
	for list1 != head1 && list2 != head2 {
		if list1.Val != list2.Val {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}
	return list1 == head1 && list2 == head2
}

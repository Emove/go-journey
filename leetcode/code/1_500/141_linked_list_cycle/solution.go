package _41_linked_list_cycle

import "leetcode/code/structure"

func HasCycle(head *structure.ListNode) bool {
	m := make(map[*structure.ListNode]struct{})
	node := head
	for node != nil {
		if _, ok := m[node]; ok {
			return true
		} else {
			m[node] = struct{}{}
		}
		node = node.Next
	}
	return false
}

func HasCycle_1(head *structure.ListNode) bool {
	if head == nil {
		return false
	}
	l, r := head, head.Next
	for l != r {
		if r == nil || r.Next == nil {
			return false
		}
		l = l.Next
		r = r.Next.Next
	}
	return true
}

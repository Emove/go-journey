package _3_merged_k_sorted_lists

import (
	"leetcode/code/structure"
)

func MergeKLists(lists []*structure.ListNode) *structure.ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		ls := make([]*structure.ListNode, 0)
		n := len(lists)
		if n&1 == 1 {
			ls = append(ls, lists[n-1])
		}
		for i := 0; i < n-1; i += 2 {
			ls = append(ls, doMerge(lists[i], lists[i+1]))
		}
		lists = ls
	}
	return lists[0]
}

func MergeKLists1(lists []*structure.ListNode) *structure.ListNode {
	if len(lists) == 0 {
		return nil
	}
	type sortedListNode struct {
		node *structure.ListNode
		next *sortedListNode
	}
	// 在这里用前插法属实拉跨
	// 应该换成优先队列
	insert := func(head *sortedListNode, node *structure.ListNode) {
		if node == nil {
			return
		}
		for head.next != nil && head.next.node.Val < node.Val {
			head = head.next
		}
		head.next = &sortedListNode{
			node: node,
			next: head.next,
		}
	}
	head := &sortedListNode{}
	for _, list := range lists {
		insert(head, list)
	}
	dummy := &structure.ListNode{}
	tail := dummy
	for head.next != nil {
		node := head.next.node
		head.next = head.next.next

		tail.Next = node
		tail = tail.Next
		node = node.Next
		insert(head, node)
	}
	return dummy.Next
}

func doMerge(list1, list2 *structure.ListNode) *structure.ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	dummy := &structure.ListNode{}
	node := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}

	if list1 != nil {
		node.Next = list1
	}

	if list2 != nil {
		node.Next = list2
	}
	return dummy.Next
}

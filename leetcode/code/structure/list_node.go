package structure

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func BuildListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := NewListNode(nums[0])
	point := head
	length := len(nums)
	for i := 1; i < length; i++ {
		node := NewListNode(nums[i])
		point.Next = node
		point = node
	}
	return head
}

func BuildCircularListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := NewListNode(nums[0])
	point := head
	length := len(nums)
	for i := 1; i < length; i++ {
		node := NewListNode(nums[i])
		point.Next = node
		point = node
	}
	point.Next = head
	return head
}

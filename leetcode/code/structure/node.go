package structure

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func BuildNode(nums []int) *Node {
	length := len(nums)
	if length == 0 {
		return nil
	}

	root := &Node{
		Val: nums[0],
	}

	nodeMpa := make(map[int]*Node, length)
	nodeMpa[0] = root
	for i := 1; i < length; i++ {
		val := nums[i]
		if val == -1 {
			continue
		}

		var parent *Node
		var node *Node
		if i%2 == 1 {
			parent = nodeMpa[(i-1)/2]
			if parent == nil {
				continue
			}
			node = &Node{Val: val}
			parent.Left = node
		} else {
			parent = nodeMpa[(i/2)-1]
			if parent == nil {
				continue
			}
			node = &Node{Val: val}
			parent.Right = node
		}
		nodeMpa[i] = node
	}
	return root
}

func PrintNode(root *Node) {
	queue := make([]*Node, 1)
	queue[0] = root
	index := 0
	for index < len(queue) {
		node := queue[index]
		fmt.Printf("%d, ", node.Val)
		if node.Val == -1 || (node.Left == nil && node.Right == nil) {
			index++
			continue
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		} else {
			queue = append(queue, &Node{Val: -1})
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		} else {
			queue = append(queue, &Node{Val: -1})
		}
		index++
	}
}

func PrintNodeNext(root *Node) {
	if root == nil || root.Left == nil {
		return
	}

	layer := root
	for layer != nil {
		head := layer
		for head != nil {
			fmt.Printf("%d, ", head.Val)
			head = head.Next
		}
		fmt.Printf("#, ")
		layer = layer.Left
	}
}

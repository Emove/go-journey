package structure

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTreeNode(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}

	nodeMpa := make(map[int]*TreeNode, length)
	nodeMpa[0] = root
	for i := 1; i < length; i++ {
		val := nums[i]
		if val == -1 {
			continue
		}

		var parent *TreeNode
		var node *TreeNode
		if i%2 == 1 {
			parent = nodeMpa[(i-1)/2]
			if parent == nil {
				continue
			}
			node = &TreeNode{Val: val}
			parent.Left = node
		} else {
			parent = nodeMpa[(i/2)-1]
			if parent == nil {
				continue
			}
			node = &TreeNode{Val: val}
			parent.Right = node
		}
		nodeMpa[i] = node
	}
	return root
}

func PrintTreeNode(root *TreeNode) {
	queue := make([]*TreeNode, 1)
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
			queue = append(queue, &TreeNode{Val: -1})
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		} else {
			queue = append(queue, &TreeNode{Val: -1})
		}
		index++
	}

}

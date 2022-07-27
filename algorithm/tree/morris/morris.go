package morris

import (
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// Morris 莫里斯二叉树遍历算法
// node.left == nil
//		node = node.right
// node.left != nil
// 		find mostRight
//		mostRight.right == nil
//			mostRight.right = node
//			node = node.left
//		mostRight.right == node
//			mostRight.right = nil
//			node = node.right
func Morris(root *TreeNode) {
	node := root
	for node != nil {

		if node.left != nil {
			// I、左子树不为空
			// I-a、找到mostRight
			mostRight := node.left
			for mostRight.right != nil && mostRight.right != node {
				mostRight = mostRight.right
			}

			if mostRight.right == nil {
				// I-b、找到了该层的mostRight，将其右指针指向node
				//		node继续往下层左子树移动
				mostRight.right = node
				node = node.left
			} else {
				// I-c、mostRight不为空，该节点之前已经被遍历过一次了
				// 		第二次来到这个节点时，将该节点的右指针指向空
				//		并使node往右走
				mostRight.right = nil
				node = node.right
			}
		} else {
			// II、左子树为空、子树往右走
			node = node.right
		}
	}
}

// MorrisPre 前序遍历
// 在每次访问到一个未访问过的节点时，将其加入结果
func MorrisPre(root *TreeNode) []int {
	node := root
	ans := make([]int, 0)
	var mostRight *TreeNode
	for node != nil {
		if node.left != nil {
			mostRight = node.left
			for mostRight.right != nil && mostRight.right != node {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {

				ans = append(ans, node.val)
				fmt.Printf("$%d\t", node.val)

				mostRight.right = node
				node = node.left
			} else {
				mostRight.right = nil
				node = node.right
			}
		} else {
			ans = append(ans, node.val)
			fmt.Printf("#%d\t", node.val)
			node = node.right
		}
	}
	return ans
}

// MorrisIn 中序遍历
// 如果一个节点只会被访问一次，则将其加入结果
// 如果一个节点会被访问两次，则在第二次访问时将其加入结果
func MorrisIn(root *TreeNode) []int {
	node := root
	var mostRight *TreeNode
	ans := make([]int, 0)
	for node != nil {
		if node.left != nil {
			mostRight = node.left
			for mostRight.right != nil && mostRight.right != node {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = node
				node = node.left
				continue
			}
			mostRight.right = nil
		}
		ans = append(ans, node.val)
		fmt.Printf("#%d\t", node.val)
		node = node.right
	}
	return ans
}

// MorrisPost 后序遍历
func MorrisPost(root *TreeNode) []int {
	node := root
	var mostRight *TreeNode
	ans := make([]int, 0)
	reverseRightNode := func(node *TreeNode) {
		start := len(ans)
		for node != nil {
			ans = append(ans, node.val)
			node = node.right
		}
		for left, right := start, len(ans)-1; left < right; {
			ans[left], ans[right] = ans[right], ans[left]
			left++
			right--
		}
	}
	for node != nil {
		if node.left != nil {
			mostRight = node.left
			for mostRight.right != nil && mostRight.right != node {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = node
				node = node.left
			} else {
				mostRight.right = nil
				//ans = append(ans, ReverseRightNode(node.left)...)
				reverseRightNode(node.left)
				node = node.right
			}
		} else {
			node = node.right
		}
	}
	//ans = append(ans, ReverseRightNode(root)...)
	reverseRightNode(root)
	return ans
}

// ReverseRightNode 逆序右子节点
func ReverseRightNode(root *TreeNode) []int {
	if root.right == nil {
		return []int{root.val}
	}
	ans := make([]int, 0)
	node := root
	for node != nil {
		ans = append(ans, node.val)
		node = node.right
	}
	for left, right := 0, len(ans)-1; left < right; {
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}
	return ans
	//pre := (*TreeNode)(nil)
	//for node != nil {
	//	next := node.right
	//	node.right = pre
	//	pre = node
	//	node = next
	//}
	//node, pre = pre, nil
	//for node != nil {
	//	ans = append(ans, node.val)
	//	next := node.right
	//	node.right = pre
	//	pre = node
	//	node = next
	//}
	//return ans
}

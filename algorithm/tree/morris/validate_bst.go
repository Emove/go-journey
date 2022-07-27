package morris

import "math"

// ValidateBST 验证一棵树是否是搜索二叉树
// 利用搜索二叉树中序遍历序列的有序性
func ValidateBST(root *TreeNode) bool {
	node, mostRight, last := root, (*TreeNode)(nil), int64(math.MinInt32-1)
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
		if last > int64(node.val) {
			return false
		}
		last = int64(node.val)
		node = node.right
	}
	return true
}

package _022_sum_of_root_to_leaf_binary_numbers

import (
	"leetcode/code/structure"
	"strconv"
)

func SumRootToLeaf(root *structure.TreeNode) int {
	node, prev, ans := root, root, 0
	stk, binSeq := make([]*structure.TreeNode, 0), make([]byte, 0)
	for node != nil || len(stk) > 0 {
		for node != nil {
			stk = append(stk, node)
			binSeq = append(binSeq, byte(node.Val+48))
			node = node.Left
		}
		n := len(stk) - 1
		node = stk[n]
		stk = stk[:n]
		if node.Left == nil && node.Right == nil {
			val, _ := strconv.ParseInt(string(binSeq), 2, 0)
			ans += int(val)
		}
		binSeq = binSeq[:n]
		if node.Right == nil || node.Right == prev {
			prev = node
			node = nil
		} else {
			stk = append(stk, node)
			binSeq = append(binSeq, byte(node.Val+48))
			node = node.Right
		}
	}
	return ans
}

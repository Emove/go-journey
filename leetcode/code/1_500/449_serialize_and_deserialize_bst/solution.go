package _49_serialize_and_deserialize_bst

import (
	"leetcode/code/structure"
	"sort"
	"strconv"
	"strings"
)

const DELIMITER = ","

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *structure.TreeNode) string {
	builder := &strings.Builder{}
	var postorder func(node *structure.TreeNode)
	postorder = func(node *structure.TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		builder.WriteString(strconv.Itoa(node.Val))
		builder.WriteString(DELIMITER)
	}
	postorder(root)
	return builder.String()[:builder.Len()-1]
}

func (c *Codec) deserialize(data string) *structure.TreeNode {
	if len(data) == 0 {
		return nil
	}
	split := strings.Split(data, DELIMITER)
	post := make([]int, len(split))
	for i, v := range split {
		post[i], _ = strconv.Atoi(v)
	}
	mid := make([]int, len(split))
	copy(mid, post)
	sort.Ints(mid)
	var build func(mid, post []int) *structure.TreeNode
	build = func(mid, post []int) *structure.TreeNode {
		n := len(post)
		if n == 0 {
			return nil
		}
		node := &structure.TreeNode{
			Val: post[n-1],
		}
		if n == 1 {
			return node
		}
		index := 0
		for ; index < len(mid) && mid[index] != node.Val; index++ {
		}
		node.Left = build(mid[:index], post[:index])
		node.Right = build(mid[index+1:], post[index:n-1])
		return node
	}
	return build(mid, post)
}

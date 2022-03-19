package _06_construct_string_from_binary_tree

import (
	"leetcode/code/structure"
	"strconv"
	"strings"
)

func tree2str(root *structure.TreeNode) string {
	if root == nil {
		return ""
	}
	type N struct {
		node *structure.TreeNode
		flag int
	}
	ns := []*N{&N{root, 0}}
	builder := &strings.Builder{}
	//builder.WriteString(strconv.Itoa(root.Val))
	//if root.Left == nil && root.Right != nil {
	//	builder.WriteString("()")
	//}
	//if root.Right != nil {
	//	ns = append(ns, &N{root.Right, 0})
	//}
	//if root.Left != nil {
	//	ns = append(ns, &N{root.Left, 0})
	//}
	for len(ns) > 0 {
		last := len(ns) - 1
		n := ns[last]
		if n.flag == 0 {
			n.flag = 1
			ns[last] = n
			if n.node != root {
				builder.WriteByte('(')
			}
			builder.WriteString(strconv.Itoa(n.node.Val))
			if n.node.Left == nil && n.node.Right != nil {
				builder.WriteString("()")
			}
			if n.node.Right != nil {
				ns = append(ns, &N{n.node.Right, 0})
			}
			if n.node.Left != nil {
				ns = append(ns, &N{n.node.Left, 0})
			}
		} else {
			if n.node != root {
				builder.WriteByte(')')
			}
			ns = ns[:last]
		}
	}
	return builder.String()
}

func dfs(root *structure.TreeNode) string {
	if root == nil {
		return ""
	}
	node := root
	builder := strings.Builder{}
	builder.WriteString(strconv.Itoa(node.Val))
	if node.Left == nil && node.Right == nil {
		return builder.String()
	} else if node.Right == nil {
		builder.WriteByte('(')
		builder.WriteString(dfs(node.Left))
		builder.WriteByte(')')
	} else {
		builder.WriteByte('(')
		builder.WriteString(dfs(node.Left))
		builder.WriteByte(')')
		builder.WriteByte('(')
		builder.WriteString(dfs(node.Right))
		builder.WriteByte(')')
	}
	return builder.String()
}

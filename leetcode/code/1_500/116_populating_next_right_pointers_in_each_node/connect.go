package populating_next_right_pointers_in_each_node_116

import (
	"leetcode/code/structure"
)

func connect(node *structure.Node) *structure.Node {
	if node == nil {
		return nil
	}

	layer := node
	for layer.Left != nil {
		pointer := layer

		for pointer != nil {
			if pointer.Left != nil {
				pointer.Left.Next = pointer.Right
			}

			if pointer.Next != nil {
				pointer.Right.Next = pointer.Next.Left
			}

			pointer = pointer.Next
		}
		layer = layer.Left
	}
	return node
}

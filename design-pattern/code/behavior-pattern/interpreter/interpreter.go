package interpreter

import (
	"strconv"
	"strings"
)

// 解释器模式定义一套语言文法，并设计该语言解释器，使用户能使用特定文法控制解释器行为
// 解释器模式的意义在于，它分离多种复杂功能的实现，每个功能只需关注自身的解释
// 对于调用者不用关心内部的解释器的工作，只需要用简单的方式组合命令就可以

type Node interface {
	Interpret() int
}

type ValNode struct {
	val int
}

func (node *ValNode) Interpret() int {
	return node.val
}

type AddNode struct {
	left, right Node
}

func (node *AddNode) Interpret() int {
	return node.left.Interpret() + node.right.Interpret()
}

type MinNode struct {
	left, right Node
}

func (node *MinNode) Interpret() int {
	//if node.left.Interpret() < node.right.Interpret() {
	//	node.left.Interpret()
	//}
	return node.left.Interpret() - node.right.Interpret()
}

type Parser struct {
	exp   []string
	index int
	prev  Node
}

func (parser *Parser) Parse(exp string) {
	parser.exp = strings.Split(exp, " ")
	for {
		if parser.index >= len(parser.exp) {
			return
		}
		switch parser.exp[parser.index] {
		case "+":
			parser.prev = parser.newAddNode()
		case "-":
			parser.prev = parser.newMinNode()
		default:
			parser.prev = parser.newValNode()
		}
	}
}

func (parser *Parser) newAddNode() Node {
	parser.index++
	return &AddNode{
		left:  parser.prev,
		right: parser.newValNode(),
	}
}

func (parser *Parser) newMinNode() Node {
	parser.index++
	return &MinNode{
		left:  parser.prev,
		right: parser.newValNode(),
	}
}

func (parser *Parser) newValNode() Node {
	v, _ := strconv.Atoi(parser.exp[parser.index])
	parser.index++
	return &ValNode{
		val: v,
	}
}

func (parser *Parser) Result() Node {
	return parser.prev
}

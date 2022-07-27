///**
// * 双向跳表
// * @author Emove
// * @date 2021/1/19
// */
package skip_list

import (
	"fmt"
	"math/rand"
	"strconv"
)

type SkipNode struct {
	value    uint64
	isHeader bool
	pre      *SkipNode
	next     *SkipNode
	down     *SkipNode
}

func NewHeaderNode() *SkipNode {
	return &SkipNode{
		value:    0,
		isHeader: true,
	}
}

func NewNode(value uint64) *SkipNode {
	return &SkipNode{
		value:    value,
		isHeader: false,
	}
}

const MaxLevel = 32

type SkipList struct {
	header           *SkipNode
	currentHighLevel uint32
}

func NewSkipList() *SkipList {
	return &SkipList{
		header:           NewHeaderNode(),
		currentHighLevel: 0,
	}
}

// 精确查找目标值
func (list *SkipList) Search(value uint64) (bool, uint64) {
	temp := list.header
	for temp != nil {
		if temp.value == value {
			if temp.isHeader {
				return false, 0
			}
			return true, temp.value
		} else if temp.next == nil {
			// 没有下一个节点，只能下降
			temp = temp.down
		} else if temp.next.value > value {
			//需要下降去寻找
			temp = temp.down
		} else {
			temp = temp.next
		}
	}
	return false, 0
}

// 查询小于当前传入目标值的所有值
func (list *SkipList) SearchEqualOrLessThan(value uint64) (bool, []uint64) {
	temp := list.header
	result := make([]uint64, 0)
	for temp != nil {
		if temp.value == value {
			if temp.isHeader {
				return false, result
			}
			break
		} else if temp.next == nil {
			// 没有下一个节点，只能下降
			temp = temp.down
		} else if temp.next.value > value {
			// 如果下层不为空则到下层继续寻找
			if temp.down != nil {
				temp = temp.down
				continue
			}
			// 如果已经没有下层了，说明当前节点是小于目标值的最大节点
			break
		} else {
			temp = temp.next
		}
	}
	if temp != nil {
		for temp.pre != nil {
			result = append(result, temp.value)
			temp = temp.pre
		}
		return true, result
	}
	return false, result
}

func (list *SkipList) Delete(value uint64) {
	temp := list.header
	for temp != nil {
		if temp.next == nil {
			//右侧没有了，说明这一层找到，没有只能下降
			temp = temp.down
		} else if temp.next.value == value {
			if temp.next.next != nil {
				// 设置右侧节点的右侧节点的左节点为当前节点，双向链表是这样子的
				temp.next.next.pre = temp
			}
			//找到节点，右侧即为待删除节点
			//删除右侧节点
			temp.next = temp.next.next
			//向下继续查找删除
			temp = temp.down
		} else if temp.next.value > value {
			//右侧已经不可能了，向下
			temp = temp.down
		} else {
			//节点还在右侧
			temp = temp.next
		}
	}
}

func (list *SkipList) Add(value uint64) {
	if 0 == value {
		// 特殊情况，0被头节点占用了，为了避免引起混淆，不要用0
		return
	}
	//存储向下的节点，这些节点可能在右侧插入节点
	stack := NewStack()
	temp := list.header
	for temp != nil {
		// 进行查找操作
		if temp.next == nil {
			// 右侧没有了，只能下降
			// 记录曾经向下的节点
			stack.Push(temp)
			temp = temp.down
		} else if temp.next.value > value {
			// 需要下降去寻找
			// 记录曾经向下的节点
			stack.Push(temp)
			temp = temp.down
		} else {
			// 向右
			temp = temp.next
		}
	}

	//当前层数，从第一层添加(第一层必须添加，先添加再判断)
	level := uint32(1)
	//保持前驱节点(即down的指向，初始为null)
	var downNode *SkipNode
	for !stack.IsEmpty() {
		//在该层插入node
		//弹出待插入的左侧节点
		temp = stack.Pop()
		node := NewNode(value)
		//处理竖方向
		node.down = downNode
		//标记新的节点下次使用
		downNode = node
		if temp.next == nil {
			//右侧为null 说明插入在末尾
			temp.next = node
			node.pre = temp
		} else {
			//右侧还有节点，插入在两节点之间
			rightNode := temp.next
			node.next = rightNode
			rightNode.pre = node
			node.pre = temp
			temp.next = node
		}
		//考虑是否需要向上
		if level > MaxLevel {
			//已经到达最高级的节点啦
			return
		}
		num := rand.Intn(100) //[0-1]随机数
		if num > 50 {
			//运气不好结束
			return
		}

		level++
		if level > list.currentHighLevel {
			//比当前最大高度要高但是依然在允许范围内 需要改变head节点
			list.currentHighLevel = level
			//需要创建一个新的头节点
			highestNode := NewHeaderNode()
			highestNode.down = list.header
			//改变head
			list.header = highestNode
			// 将head压入栈中，等下一次弹出
			stack.Push(highestNode)
		}
	}
}

// for test
func (list *SkipList) printList() {
	teamNode := list.header
	index := 1
	last := teamNode
	for last.down != nil {
		last = last.down
	}
	for teamNode != nil {
		enumNode := teamNode.next
		enumLast := last.next
		fmt.Printf("%-8s", "head->")
		for enumLast != nil && enumNode != nil {
			if enumLast.value == enumNode.value {
				fmt.Printf("%-5s", strconv.FormatUint(enumLast.value, 10)+"->")
				enumLast = enumLast.next
				enumNode = enumNode.next
			} else {
				enumLast = enumLast.next
				fmt.Printf("%-5s", "")
			}
		}
		teamNode = teamNode.down
		index++
		fmt.Println()
	}
}

type Stack struct {
	size   int32
	values []*SkipNode
}

func NewStack() *Stack {
	return &Stack{
		size:   -1,
		values: make([]*SkipNode, 0),
	}
}

// 压入栈
func (stack *Stack) Push(node *SkipNode) {
	stack.values = append(stack.values, node)
	stack.size++
}

// 弹出
func (stack *Stack) Pop() *SkipNode {
	if stack.size == -1 {
		return nil
	}
	node := stack.values[len(stack.values)-1:][0]
	stack.values = stack.values[:len(stack.values)-1]
	stack.size--
	return node
}

// 是否为空
func (stack *Stack) IsEmpty() bool {
	return stack.size == -1
}

// 当前栈大小
func (stack *Stack) Size() int32 {
	return stack.size
}

///**
// * @author Emove
// * @date 2021/1/19
// */
package skip_list

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	for i := 0; i < 5; i++ {
		stack.Push(NewNode(uint64(i)))
	}

	for stack.size > 1 {
		fmt.Printf("stack size: %d, pop value: %d\n", stack.Size(), stack.Pop().value)
	}
	stack.Push(NewNode(9))
	for !stack.IsEmpty() {
		fmt.Printf("stack size: %d, pop value: %d\n", stack.Size(), stack.Pop().value)
	}
}

func TestSkipList_NewSkipList(t *testing.T) {
	list := NewSkipList()
	for i := 1; i < 20; i++ {
		list.Add(uint64(i))
	}
	list.printList()
}

func TestSkipList_Delete(t *testing.T) {
	list := NewSkipList()
	for i := 1; i < 10; i++ {
		list.Add(uint64(i))
	}
	list.printList()
	list.Delete(uint64(7))
	list.Delete(uint64(3))
	list.printList()
}

func TestSkipList_Search(t *testing.T) {
	list := NewSkipList()
	for i := 1; i < 20; i++ {
		list.Add(uint64(i))
	}
	fmt.Println(list.Search(15))
	// 如果查找节点为0，头节点会被返回
	fmt.Println(list.Search(0))
	fmt.Println(list.Search(22))
}

func TestSkipList_SearchEqualOrLessThan(t *testing.T) {
	list := NewSkipList()
	for i := 1; i < 20; i++ {
		if i == 15 {
			continue
		}
		list.Add(uint64(i))
	}
	ok, than := list.SearchEqualOrLessThan(15)
	if ok {
		for i := range than {
			fmt.Printf("%d \t", than[i])
		}
	}
}

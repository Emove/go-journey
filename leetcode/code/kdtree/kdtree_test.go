package kdtree

import (
	"fmt"
	"testing"
)

func TestKdTree(t *testing.T) {
	tree := NewKDTree(100)
	tree.Put("3,6", []int{3, 6})
	tree.Put("2,7", []int{2, 7})
	tree.Put("17,15", []int{17, 15})
	tree.Put("13,15", []int{13, 15})
	tree.Put("6,12", []int{6, 12})
	tree.Put("9,1", []int{9, 1})
	tree.Put("10,19", []int{10, 19})

	tree.Build()

	tree.print()

	result := tree.SearchRange([][]int{{8, 10}, {0, 2}})
	t.Logf("%v", result)
}

func TestKdTree2(t *testing.T) {
	tree := NewKDTree(1000)
	for x := 1; x <= 10; x++ {
		for y := 1; y <= 10; y++ {
			tree.Put(fmt.Sprintf("%d,%d", x, y), []int{x, y})
		}
	}

	tree.Build()

	tree.print()

	result := tree.SearchRange([][]int{{8, 10}, {0, 2}})
	t.Logf("%v", result)
}

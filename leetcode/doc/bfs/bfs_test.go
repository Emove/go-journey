package bfs

import (
	"leetcode/code/structure"
	"testing"
)

func TestBFS(t *testing.T) {
	BFS(structure.BuildNode([]int{1, 3, 5, 7, 9, 11}))
}

func TestBFSEdges(t *testing.T) {
	//BFSByEdges(4, [][]int{{1,0},{1,2},{1,3}})
	BFSEdges(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}})
}

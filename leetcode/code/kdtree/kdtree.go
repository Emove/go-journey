package kdtree

import (
	"fmt"
)

/**
K Dimensional Tree
任意给个矩形，检索里面有几个node
*/

const (
	K int = 2
)

type KDNode struct {
	nodeID      string
	point       []int
	distance    float64
	left, right *KDNode
}
type KDTree struct {
	root *KDNode
	sort []*KDNode
	len  int
}

func NewKDTree(sz int) (t *KDTree) {
	t = new(KDTree)
	t.sort = make([]*KDNode, sz)
	t.len = 0
	t.root = nil
	return
}

func newKdNode(nodeID string, p []int) (n *KDNode) {
	n = &KDNode{
		nodeID: nodeID,
		//x分量加点权重
		distance: float64(p[0]*p[0])*1.000001 + float64(p[1]*p[1]),
		point:    make([]int, len(p)),
		left:     nil,
		right:    nil,
	}
	//assert len(point) == K
	copy(n.point, p)
	return
}

func kdInsertRec(n *KDNode, m *KDNode, depth int) *KDNode {
	if n == nil {
		return m
	}
	cd := depth % K
	if m.point[cd] < n.point[cd] {
		n.left = kdInsertRec(n.left, m, depth+1)
	} else {
		n.right = kdInsertRec(n.right, m, depth+1)
	}
	return n
}

func (t *KDTree) print() {
	printTree(t.root)
}

func printTree(root *KDNode) {
	q := []*KDNode{root}
	for len(q) > 0 {
		q1 := q[:]
		for _, n := range q1 {
			fmt.Printf(" %s(%d,%d)", n.nodeID, n.point[0], n.point[1])
			if n.left != nil {
				q = append(q, n.left)
			}
			if n.right != nil {
				q = append(q, n.right)
			}
		}
		q = q[len(q1):]
		fmt.Println()
	}
}

func isPointSame(p1, p2 []int) bool {
	return p1[0] == p2[0] && p1[1] == p2[1]
}

func kdSearchRec(n *KDNode, point []int, depth int) string {
	if n == nil {
		return ""
	}
	if isPointSame(n.point, point) {
		return n.nodeID
	}
	cd := depth % K
	if point[cd] < n.point[cd] {
		return kdSearchRec(n.left, point, depth+1)
	}
	return kdSearchRec(n.right, point, depth+1)
}

func (t *KDTree) SearchOne(p []int) string {
	return kdSearchRec(t.root, p, 0)
}

/**
 * 查找一个范围
 */
func kdSearchRangeR(n *KDNode, r [][]int, depth int, nodeIDList []string) []string {
	if n == nil {
		return nodeIDList
	}
	cd := depth % K
	//在内部
	if n.point[cd] <= r[cd][1] && n.point[cd] >= r[cd][0] {
		other := (cd + 1) % K
		if n.point[other] <= r[other][1] && n.point[other] >= r[other][0] {
			//添加点
			nodeIDList = append(nodeIDList, n.nodeID)
		}
		//递归左右子树
		return kdSearchRangeR(n.right, r, depth+1, kdSearchRangeR(n.left, r, depth+1, nodeIDList))
	}

	if n.point[cd] > r[cd][1] {
		//递归左子树
		return kdSearchRangeR(n.left, r, depth+1, nodeIDList)
	}
	//递归右子树
	return kdSearchRangeR(n.right, r, depth+1, nodeIDList)
}

func (t *KDTree) SearchRange(r [][]int) []string {
	return kdSearchRangeR(t.root, r, 0, []string{})
}

func (t *KDTree) Put(nodeID string, p []int) {
	n := newKdNode(nodeID, p)

	//按到原点的距离排序
	var i int
	for i = 0; i < t.len; i++ {
		if n.distance < t.sort[i].distance {
			for j := t.len; j > i; j-- {
				t.sort[j] = t.sort[j-1]
			}
			break
		}
	}
	t.sort[i] = n
	t.len++
	return
}

func (t *KDTree) Build() {
	var q [][]int
	if t.len > 0 {
		q = append(q, []int{0, t.len})
	}
	for len(q) > 0 {
		be := q[0]
		q = q[1:]
		m := (be[1] + be[0]) / 2
		t.root = kdInsertRec(t.root, t.sort[m], 0)
		if m > be[0] {
			q = append(q, []int{be[0], m})
		}
		if be[1] > m+1 {
			q = append(q, []int{m + 1, be[1]})
		}
	}
}

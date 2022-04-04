package segment_tree

type SegmentTree struct {
	root *node
}

type node struct {
	val   int
	left  *node
	right *node
	start int
	end   int
}

func BuildSegmentTree(nums []int) *SegmentTree {
	length := len(nums)
	if length == 0 {
		return nil
	}

	root := &node{
		val:   nums[0],
		start: 0,
		end:   length - 1,
	}
	build(nums, root, 0, length-1)
	return &SegmentTree{
		root: root,
	}
}

func build(nums []int, n *node, start, end int) {
	if start == end {
		n.val = nums[start]
		return
	}

	mid := (start + end) / 2
	n.left = &node{
		start: start,
		end:   mid,
	}
	n.right = &node{
		start: mid + 1,
		end:   end,
	}

	build(nums, n.left, start, mid)
	build(nums, n.right, mid+1, end)

	n.val = n.left.val + n.right.val
}

func (tree *SegmentTree) update(index, val int) {
	if index >= tree.root.start && index <= tree.root.end {
		tree.root.doUpdate(index, val)
	}
}

func (tree *SegmentTree) query(left, right int) int {
	root := tree.root
	if right < root.start || left > root.end {
		return 0
	}
	return root.queryRange(left, right)
}

func (n *node) queryRange(left, right int) int {
	if n == nil || right < n.start || left > n.end {
		return 0
	} else if left <= n.start && n.end <= right {
		return n.val
	}
	return n.left.queryRange(left, right) + n.right.queryRange(left, right)
}

func (n *node) doUpdate(index, val int) {
	if n == nil {
		return
	}
	if n.start == n.end && n.start == index {
		n.val = val
		return
	}
	mid := (n.start + n.end) / 2
	if index >= n.start && index <= mid {
		n.left.doUpdate(index, val)
	} else {
		n.right.doUpdate(index, val)
	}
	n.val = n.left.val + n.right.val
}

package _29_my_calendar_i

// 线段树
type MyCalendar struct {
	tree, lazy map[int]bool
}

func Constructor() MyCalendar {
	return MyCalendar{tree: map[int]bool{}, lazy: map[int]bool{}}
}

func (c *MyCalendar) query(start, end, l, r, idx int) bool {
	if end < l || r < start {
		return false
	}
	if c.lazy[idx] {
		return true
	}

	if start <= l && r <= end {
		return c.tree[idx]
	}

	mid := (l + r) >> 1
	return c.query(start, end, l, mid, 2*idx) || c.query(start, end, mid+1, r, 2*idx+1)
}

func (c *MyCalendar) update(start, end, l, r, idx int) {
	if end < l || r < start {
		return
	}
	if start <= l && r <= end {
		c.tree[idx] = true
		c.lazy[idx] = true
	} else {
		mid := (l + r) >> 1
		c.update(start, end, l, mid, 2*idx)
		c.update(start, end, mid+1, r, 2*idx+1)
		c.tree[idx] = true
		if c.lazy[2*idx] && c.lazy[2*idx+1] {
			c.lazy[idx] = true
		}
	}
}

func (c *MyCalendar) Book(start int, end int) bool {
	if c.query(start, end-1, 0, 1e9, 1) {
		return false
	}
	c.update(start, end-1, 0, 1e9, 1)
	return true
}

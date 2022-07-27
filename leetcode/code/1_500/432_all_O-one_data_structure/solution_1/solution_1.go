package solution_1

type AllOne struct {
	head, tail *node
	m          map[string]*node // 所有key所在node的索引
}

type node struct {
	cnt        int            // 当前node所代表的计数
	prev, next *node          // 前后节点
	keys       []string       // 计数为cnt的key
	keyMap     map[string]int // key在keys中的下标，用于加速删除
}

func newNode(cnt int) *node {
	return &node{
		cnt:    cnt,
		keyMap: make(map[string]int),
	}
}

func newNodeWithPrevAndNex(cnt int, prev, next *node) *node {
	return &node{
		cnt:    cnt,
		prev:   prev,
		next:   next,
		keyMap: make(map[string]int),
	}
}

func (n *node) add(key string) {
	length := len(n.keys)
	n.keys = append(n.keys, key)
	n.keyMap[key] = length
}

func (n *node) delete(key string) {
	if index, ok := n.keyMap[key]; ok {
		delete(n.keyMap, key)
		nl := len(n.keys)
		if nl == index {
			n.keys = n.keys[:index]
		} else if nl == 0 {
			n.keys = n.keys[1:]
			for k, v := range n.keyMap {
				n.keyMap[k] = v - 1
			}
		} else {
			n.keys = append(n.keys[:index], n.keys[index+1:]...)
			for i := index; i < len(n.keys); i++ {
				n.keyMap[n.keys[i]]--
			}
		}
	}
}

func (n *node) clear() {
	if len(n.keys) == 0 {
		n.prev.next = n.next
		n.next.prev = n.prev
	}
}

func Constructor() AllOne {
	head := newNode(-1)
	tail := newNode(-1)
	head.next = tail
	tail.prev = head
	return AllOne{
		head: head,
		tail: tail,
		m:    make(map[string]*node),
	}
}

func (ao *AllOne) Inc(key string) {
	if n, ok := ao.m[key]; ok {
		n.delete(key)
		cnt := n.cnt
		var next *node
		if n.next.cnt == cnt+1 {
			next = n.next
		} else {
			next = newNodeWithPrevAndNex(cnt+1, n, n.next)
			n.next.prev = next
			n.next = next
		}
		next.add(key)
		ao.m[key] = next
		n.clear()
	} else {
		var node *node
		if ao.head.next.cnt == 1 {
			node = ao.head.next
		} else {
			node = newNodeWithPrevAndNex(1, ao.head, ao.head.next)
			ao.head.next.prev = node
			ao.head.next = node
		}
		node.add(key)
		ao.m[key] = node
	}
}

func (ao *AllOne) Dec(key string) {
	n := ao.m[key]
	n.delete(key)
	cnt := n.cnt
	if cnt == 1 {
		delete(ao.m, key)
	} else {
		var prev *node
		if n.prev.cnt == cnt-1 {
			prev = n.prev
		} else {
			prev = newNodeWithPrevAndNex(cnt-1, n.prev, n)
			n.prev.next = prev
			n.prev = prev
		}
		prev.add(key)
		ao.m[key] = prev
	}
	n.clear()
}

func (ao *AllOne) GetMaxKey() string {
	prev := ao.tail.prev
	if len(prev.keys) > 0 {
		return prev.keys[0]
	}
	return ""
}

func (ao *AllOne) GetMinKey() string {
	next := ao.head.next
	if len(next.keys) > 0 {
		return next.keys[0]
	}
	return ""
}

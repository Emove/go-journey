package _32_all_O_one_data_structure

const (
	MAX string = "max"
	MIN string = "min"
)

type AllOne struct {
	index     map[string]*Node
	maxAndMin map[string]*Node
}

type Node struct {
	key string
	cnt int
}

func Constructor() AllOne {
	return AllOne{
		index:     make(map[string]*Node),
		maxAndMin: make(map[string]*Node),
	}
}

func (this *AllOne) Inc(key string) {
	var node *Node
	ok := false
	if node, ok = this.index[key]; ok {
		node.cnt++
		n := this.maxAndMin[MAX]
		if node.cnt > n.cnt {
			this.maxAndMin[MAX] = node
		}
		if n, ok := this.maxAndMin[MIN]; ok && n.key == node.key {
			minKeyCnt := n.cnt
			for _, n1 := range this.index {
				if n1.cnt < minKeyCnt {
					minKeyCnt = n.cnt
					n = n1
				}
			}
			this.maxAndMin[MIN] = n
		}
	} else {
		node := &Node{key: key, cnt: 1}
		this.index[key] = node
		if _, ok := this.maxAndMin[MAX]; !ok {
			this.maxAndMin[MAX] = node
		}
		if n, ok := this.maxAndMin[MIN]; ok {
			if n.cnt > node.cnt {
				this.maxAndMin[MIN] = node
			}
		} else {
			this.maxAndMin[MIN] = node
		}
	}
}

func (this *AllOne) Dec(key string) {
	if node, ok := this.index[key]; ok {
		node.cnt--
		if node.cnt == 0 {
			delete(this.index, key)
		}
		indexLen := len(this.index)
		if indexLen == 0 {
			this.maxAndMin[MAX] = nil
			this.maxAndMin[MIN] = nil
		}
		maxKeyNode := this.maxAndMin[MAX]
		maxKeyCnt := maxKeyNode.cnt
		minKeyNode := maxKeyNode
		minKeyCnt := maxKeyCnt
		for _, n := range this.index {
			if n.cnt > maxKeyCnt {
				maxKeyCnt = n.cnt
				maxKeyNode = n
			}
			if n.cnt < minKeyCnt {
				minKeyCnt = n.cnt
				minKeyNode = n
			}
		}
		this.maxAndMin[MAX] = maxKeyNode
		this.maxAndMin[MIN] = minKeyNode
	}
}

func (this *AllOne) GetMaxKey() string {
	if n, ok := this.maxAndMin[MAX]; ok {
		return n.key
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if n, ok := this.maxAndMin[MIN]; ok {
		return n.key
	}
	return ""
}

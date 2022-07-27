package solution_3

import (
	"math"
)

type AllOne struct {
	index map[string]int
}

func Constructor() AllOne {
	return AllOne{
		index: make(map[string]int),
	}
}

func (this *AllOne) Inc(key string) {
	if _, ok := this.index[key]; ok {
		this.index[key]++
	} else {
		this.index[key] = 1
	}
}

func (this *AllOne) Dec(key string) {
	if _, ok := this.index[key]; ok {
		this.index[key]--
		if this.index[key] == 0 {
			delete(this.index, key)
		}
	}
}

func (this *AllOne) GetMaxKey() string {
	cnt := 0
	res := ""
	for k, v := range this.index {
		if v > cnt {
			cnt = v
			res = k
		}
	}
	return res
}

func (this *AllOne) GetMinKey() string {
	cnt := math.MaxInt32
	res := ""
	for k, v := range this.index {
		if v < cnt {
			cnt = v
			res = k
		}
	}
	return res
}

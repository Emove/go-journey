package _80_insert_delete_get_random_O_one

import "math/rand"

type RandomizedSet struct {
	nums  []int
	index map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:  []int{},
		index: make(map[int]int),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.index[val]; ok {
		return false
	}
	rs.index[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	i, ok := rs.index[val]
	if !ok {
		return false
	}
	last := len(rs.nums) - 1
	rs.nums[i] = rs.nums[last]
	rs.index[rs.nums[i]] = i
	delete(rs.index, val)
	rs.nums = rs.nums[:last]
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}

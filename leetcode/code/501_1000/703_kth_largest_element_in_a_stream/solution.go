package _03_kth_largest_element_in_a_stream

import (
	"container/heap"
	"sort"
)

type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for i := range nums {
		kl.Add(nums[i])
	}
	return kl
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Pop() interface{} {
	slice := kl.IntSlice
	v := slice[len(slice)-1]
	kl.IntSlice = slice[:len(slice)-1]
	return v
}

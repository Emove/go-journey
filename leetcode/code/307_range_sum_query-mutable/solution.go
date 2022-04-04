package _07_range_sum_query_mutable

// NumArray 树状数组解法
type NumArray struct {
	nums []int
	t    []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	ins := NumArray{
		nums: nums,
		t:    make([]int, n+1),
		n:    n,
	}
	for i := 0; i < n; i++ {
		ins.add(i+1, nums[i])
	}
	return ins
}

func (this *NumArray) Update(index int, val int) {
	this.add(index+1, val-this.nums[index])
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.query(right+1) - this.query(left)
}

func lowbit(x int) int {
	return x & -x
}

func (this *NumArray) query(index int) (ans int) {
	for i := index; i > 0; i -= lowbit(i) {
		ans += this.t[i]
	}
	return
}

func (this *NumArray) add(index, k int) {
	for i := index; i <= this.n; i += lowbit(i) {
		this.t[i] += k
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

package permutation_46

var res [][]int
var path []int
var used []bool

func permute(nums []int) [][]int {
	n := len(nums)
	res = make([][]int, 0)
	used = make([]bool, n)
	path = make([]int, n)
	dfs(nums, n, 0)
	return res
}

func dfs(nums []int, n, k int) {
	if k == n {
		item := make([]int, n)
		copy(item, path)
		res = append(res, item)
		return
	}

	for i := 0; i < n; i++ {
		if !used[i] {
			path[k] = nums[i]
			used[i] = true
			dfs(nums, n, k+1)
			used[i] = false
		}
	}
}

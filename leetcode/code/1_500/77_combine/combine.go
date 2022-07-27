package combine_77

var res [][]int
var path []int

func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	if k <= 0 || n < k {
		return res
	}
	path = make([]int, k)
	dfs(n, k, 0, 1)
	return res
}

func dfs(n, k, index, curr int) {
	if k == 0 {
		item := make([]int, k+index)
		copy(item, path)
		res = append(res, item)
		return
	}

	if curr > n-k+1 {
		return
	}

	// 不选curr
	dfs(n, k, index, curr+1)

	path[index] = curr
	dfs(n, k-1, index+1, curr+1)
}

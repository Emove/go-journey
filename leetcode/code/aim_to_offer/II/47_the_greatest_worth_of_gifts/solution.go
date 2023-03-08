package _7_the_greatest_worth_of_gifts

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	w := make([][]int, m)
	for i := 0; i < m; i++ {
		w[i] = make([]int, n)
	}

	for i, g := range grid {
		for j, v := range g {
			if i > 0 {
				w[i][j] = max(w[i][j], w[i-1][j])
			}
			if j > 0 {
				w[i][j] = max(w[i][j], w[i][j-1])
			}
			w[i][j] += v
		}
	}
	return w[m-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package _17_pacific_atlantic_water_flow

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func PacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	pacificQueue, atlanticQueue := make([][]int, m+n-1), make([][]int, m+n-1)
	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)

		pacific[i][0] = true
		pacificQueue[i] = []int{i, 0}

		atlantic[i][n-1] = true
		atlanticQueue[i] = []int{i, n - 1}
	}
	if m >= 1 && n >= 1 {
		atlantic[m-1][0] = true
		if m < m+n-1 {
			atlanticQueue[m] = []int{m - 1, 0}
		}
		pacific[0][n-1] = true
		pacificQueue[m+n-2] = []int{0, n - 1}
	}
	for j := 1; j < n-1; j++ {
		pacific[0][j] = true
		pacificQueue[m+j-1] = []int{0, j}

		atlantic[m-1][j] = true
		atlanticQueue[m+j] = []int{m - 1, j}
	}
	bfs(m, n, heights, pacificQueue, pacific)
	bfs(m, n, heights, atlanticQueue, atlantic)
	res := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func bfs(m, n int, grid [][]int, queue [][]int, res [][]bool) {
	for queue != nil {
		q := queue
		queue = nil
		for _, cell := range q {
			for _, dir := range dirs {
				newRow, newColumn := cell[0]+dir[0], cell[1]+dir[1]
				if newRow < m && newRow >= 0 && newColumn < n && newColumn >= 0 && !res[newRow][newColumn] && grid[newRow][newColumn] >= grid[cell[0]][cell[1]] {
					res[newRow][newColumn] = true
					queue = append(queue, []int{newRow, newColumn})
				}
			}
		}
	}
}

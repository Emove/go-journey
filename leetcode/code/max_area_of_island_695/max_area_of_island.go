package max_area_of_island_695

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	depth := len(grid)
	width := len(grid[0])

	for i := 0; i < depth; i++ {
		for j := 0; j < width; j++ {
			ans = max(ans, dfs(grid, i, j, depth, width))
		}
	}

	return ans
}

func dfs(grid [][]int, x, y, depth, width int) int {
	if x < 0 || y < 0 || x == depth || y == width || grid[x][y] == 0 {
		return 0
	}
	grid[x][y] = 0
	ans := 1
	ans += dfs(grid, x+1, y, depth, width)
	ans += dfs(grid, x-1, y, depth, width)
	ans += dfs(grid, x, y+1, depth, width)
	ans += dfs(grid, x, y-1, depth, width)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

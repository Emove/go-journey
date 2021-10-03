package rotting_oranges_994

func orangesRotting(grid [][]int) int {
	depth := len(grid)
	width := len(grid[0])

	count := 0
	queue := make([][]int, 0)

	for i := 0; i < depth; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == 1 {
				count++
			} else if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	round := 0
	for count > 0 && len(queue) > 0 {
		round++
		n := len(queue)
		for k := 0; k < n; k++ {
			cell := queue[0]
			queue = queue[1:]
			x := cell[0]
			y := cell[1]
			if x-1 >= 0 && grid[x-1][y] == 1 {
				grid[x-1][y] = 2
				count--
				queue = append(queue, []int{x - 1, y})
			}
			if x+1 < depth && grid[x+1][y] == 1 {
				grid[x+1][y] = 2
				count--
				queue = append(queue, []int{x + 1, y})
			}
			if y-1 >= 0 && grid[x][y-1] == 1 {
				grid[x][y-1] = 2
				count--
				queue = append(queue, []int{x, y - 1})
			}
			if y+1 < width && grid[x][y+1] == 1 {
				grid[x][y+1] = 2
				count--
				queue = append(queue, []int{x, y + 1})
			}
		}
	}

	if count > 0 {
		return -1
	}
	return round
}

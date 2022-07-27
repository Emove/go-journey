package _75_cut_off_trees_for_golf_event

import "sort"

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func CutOffTree(forest [][]int) int {
	if forest[0][0] == 0 {
		return -1
	}
	n, m := len(forest), len(forest[0])
	rankQueue := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if forest[i][j] > 1 {
				rankQueue = append(rankQueue, []int{forest[i][j], i, j})
			}
		}
	}
	sort.Slice(rankQueue, func(i, j int) bool {
		return rankQueue[i][0] < rankQueue[j][0]
	})

	x, y, ans := 0, 0, 0

	for _, cell := range rankQueue {
		dis := bfs(forest, n, m, x, y, cell[1], cell[2])
		if dis == -1 {
			return -1
		}
		ans += dis
		x, y = cell[1], cell[2]
	}
	return ans
}

func bfs(forest [][]int, n, m, x, y, tx, ty int) int {
	if x == tx && y == ty {
		return 0
	}
	queue, visited, ans := [][]int{{x, y}}, make([][]bool, n), 0
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	visited[x][y] = true
	for queue != nil {
		q := queue
		queue = nil
		for _, point := range q {
			for _, dir := range dirs {
				nx, ny := point[0]+dir[0], point[1]+dir[1]
				if nx < 0 || nx >= n || ny < 0 || ny >= m {
					continue
				}
				if forest[nx][ny] == 0 || visited[nx][ny] == true {
					continue
				}
				if nx == tx && ny == ty {
					return ans + 1
				}
				queue = append(queue, []int{nx, ny})
				visited[nx][ny] = true
			}
		}
		ans++
	}
	return -1
}

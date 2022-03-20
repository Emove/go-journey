package _039_the_time_when_the_network_becomes_idle

func NetworkBecomesIdle(edges [][]int, patience []int) (ans int) {
	n := len(patience)
	adj := make([][]int, n)
	// 构造连通图
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	queue := []int{0}
	visit := make([]bool, n)
	visit[0] = true
	for dist := 1; queue != nil; dist++ {
		q := queue
		queue = nil
		for _, node := range q {
			for _, v := range adj[node] {
				if visit[v] {
					continue
				}
				visit[v] = true
				queue = append(queue, v)
				ans = max(ans, patience[v]*((2*dist-1)/patience[v])+2*dist+1)
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

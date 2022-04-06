package _10_minimum_height_trees

func FindMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adj := make([][]int, n)

	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	parents := make([]int, n)
	bfs := func(start int) (end int) {
		visited := make([]bool, n)
		visited[start] = true
		queue := []int{start}
		for len(queue) > 0 {
			end, queue = queue[0], queue[1:]
			for _, y := range adj[end] {
				if !visited[y] {
					parents[y] = end
					queue = append(queue, y)
					visited[y] = true
				}
			}
		}
		return
	}

	// 找到离0最远的节点x
	x := bfs(0)
	// 找到离x最远的节点y
	y := bfs(x)

	path := []int{}
	parents[x] = -1
	for y != -1 {
		path = append(path, y)
		y = parents[y]
	}

	l := len(path)
	if l&1 == 0 {
		return []int{path[l/2-1], path[l/2]}
	}
	return []int{path[l/2]}
}

func FindMinHeightTrees1(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adj := make([][]int, n)

	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	parents := make([]int, n)
	maxDepth, node := 0, -1
	var dfs func(curr, parent, depth int)
	dfs = func(curr, parent, depth int) {
		if depth > maxDepth {
			maxDepth, node = depth, curr
		}
		parents[curr] = parent
		for _, y := range adj[curr] {
			if y != parent {
				dfs(y, curr, depth+1)
			}
		}
	}

	dfs(0, -1, 1)
	maxDepth = 0
	dfs(node, -1, 1)

	path := []int{}
	for node != -1 {
		path = append(path, node)
		node = parents[node]
	}

	l := len(path)
	if l&1 == 0 {
		return []int{path[l/2-1], path[l/2]}
	}
	return []int{path[l/2]}
}

func FindMinHeightTrees2(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adj := make([][]int, n)
	deg := make([]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
		deg[edge[0]]++
		deg[edge[1]]++
	}

	queue := []int{}
	for i, d := range deg {
		if d == 1 {
			queue = append(queue, i)
		}
	}
	remainder := n
	for remainder > 2 {
		remainder -= len(queue)
		temp := queue
		queue = nil
		for _, x := range temp {
			for _, y := range adj[x] {
				deg[y]--
				if deg[y] == 1 {
					queue = append(queue, y)
				}
			}
		}
	}
	return queue
}

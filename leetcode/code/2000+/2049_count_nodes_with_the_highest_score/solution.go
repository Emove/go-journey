package _049_count_nodes_with_the_highest_score

func CountHighestScoreNodes(parents []int) int {
	n := len(parents)
	children := make([][]int, n)
	for node := 1; node < n; node++ {
		p := parents[node]
		children[p] = append(children[p], node)
	}

	maxScore := 0
	ans := 0
	var dfs func(node int) int
	dfs = func(node int) int {
		score, childSize := 1, 1
		for _, ch := range children[node] {
			remain := dfs(ch)
			// 此处score等于子树的分数
			score *= remain
			// 累加当前节点的子节点的大小
			childSize += remain
		}

		if node > 0 {
			// 计算移除当前节点后的总分数 score等于子树分数乘积、n - childSize等于移除当前节点后剩下的节点数
			score *= n - childSize
		}
		if score == maxScore {
			ans++
		} else if score > maxScore {
			maxScore = score
			ans = 1
		}
		// 以当前节点为根节点的子树大小
		return childSize
	}
	dfs(0)
	return ans
}

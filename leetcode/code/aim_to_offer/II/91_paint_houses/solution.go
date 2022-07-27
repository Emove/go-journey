package _1_paint_houses

func MinCost(costs [][]int) int {
	n := len(costs)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 3)
	}
	dp[0][0], dp[0][1], dp[0][2] = costs[0][0], costs[0][1], costs[0][2]
	for i := 1; i < n; i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		dp[i][2] = min(dp[i-1][1], dp[i-1][0]) + costs[i][2]
	}
	return min(dp[n-1][0], min(dp[n-1][1], dp[n-1][2]))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

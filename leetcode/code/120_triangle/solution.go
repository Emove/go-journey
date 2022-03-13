package _20_triangle

// 自底向上
func MinimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 1 {
		return triangle[0][0]
	}
	dp := triangle[n-1]
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func DPMinimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		curr := i & 1
		prev := 1 - curr
		// 每一层最左一位，只能从上一层的最左位移动下来
		dp[curr][0] = dp[prev][0] + triangle[i][0]

		for j := 1; j < i; j++ {
			dp[curr][j] = min(dp[prev][j-1], dp[prev][j]) + triangle[i][j]
		}

		// 同理，每一层最右一位只能从上一层最后一位移动下来
		dp[curr][i] = dp[prev][i-1] + triangle[i][i]
	}
	last := (n - 1) & 1
	sum := dp[last][0]
	for i := 1; i < n; i++ {
		sum = min(sum, dp[last][i])
	}
	return sum
}

func OptimizeDp(triangle [][]int) int {
	n := len(triangle)
	if n == 1 {
		return triangle[0][0]
	}
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		// 先算最右
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}
	sum := dp[0]
	for i := 1; i < n; i++ {
		sum = min(sum, dp[i])
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

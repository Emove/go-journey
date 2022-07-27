package _73_length_of_longest_fibonacci_subsequence

func LenLongestFibSubseq(arr []int) int {
	ans, n := 0, len(arr)
	index := make(map[int]int, n)
	for i := 0; i < n; i++ {
		index[arr[i]] = i
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0 && arr[j]*2 > arr[i]; j-- {
			if k, ok := index[arr[i]-arr[j]]; ok {
				dp[j][i] = max(dp[k][j]+1, 3)
				ans = max(ans, dp[j][i])
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

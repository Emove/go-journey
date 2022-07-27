package _09_fibonacci_number

func Fib(n int) int {
	dp := []int{0, 1}
	for i := 2; i <= n; i++ {
		dp[i&1] = dp[0] + dp[1]
	}
	return dp[n&1]
}

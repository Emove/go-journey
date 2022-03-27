package _028_find_missing_observations

func MissingRolls(rolls []int, mean, n int) []int {
	sum := mean * (n + len(rolls))
	for _, v := range rolls {
		sum -= v
	}
	if sum < n || sum > n*6 {
		return nil
	}
	quotient, remainder := sum/n, sum%n
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = quotient
		if ans[i]+remainder < 6 {
			ans[i]++
		}
	}
	return ans
}

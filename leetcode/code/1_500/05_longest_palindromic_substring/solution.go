package _5_longest_palindromic_substring

// LongestPalindrome 双指针，从中心往左右两边扩展
func LongestPalindrome(s string) string {
	n := len(s)
	if n < 1 {
		return ""
	}
	start, end := 0, 0
	cal := func(left, right int) (int, int) {
		for left >= 0 && right < n {
			if s[left] == s[right] {
				left--
				right++
			} else {
				break
			}
		}
		return left + 1, right - 1
	}
	for i := 0; i < n-1; i++ {
		left, right := cal(i, i)
		if right-left > end-start {
			start, end = left, right
		}
		left, right = cal(i, i+1)
		if right-left > end-start {
			start, end = left, right
		}
	}
	return s[start : end+1]
}

// LongestPalindrome1 动态规划
func LongestPalindrome1(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	max, start := 1, 0
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	// L表示子串长度
	for L := 2; L <= n; L++ {
		// i表示子串起始，j表示子串结束
		for i := 0; i < n; i++ {
			j := L + i - 1

			if j >= n {
				break
			}

			if s[i] == s[j] {
				if j-i <= 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}

			if dp[i][j] && j-i+1 > max {
				max = j - i + 1
				start = i
			}
		}

	}
	return s[start : start+max]
}

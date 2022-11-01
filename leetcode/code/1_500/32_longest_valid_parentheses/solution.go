package _2_longest_valid_parentheses

func LongestValidParentheses(s string) (ans int) {

	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				ans = max(ans, i-stack[len(stack)-1])
			} else {
				stack = append(stack, i)
			}
		}
	}
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

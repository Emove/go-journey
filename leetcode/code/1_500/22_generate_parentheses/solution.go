package _2_generate_parentheses

func GenerateParenthesis(n int) (ans []string) {
	s, length := make([]byte, 0), n*2
	var backtrack func(left, right int)
	backtrack = func(left, right int) {
		if left+right == length {
			ans = append(ans, string(s))
			return
		}
		if left < n {
			s = append(s, '(')
			backtrack(left+1, right)
			s = s[:len(s)-1]
		}
		if left > right && left > 0 {
			s = append(s, ')')
			backtrack(left, right+1)
			s = s[:len(s)-1]
		}
	}
	backtrack(0, 0)
	return ans
}

package _41_different_ways_to_add_parentheses

import (
	"strconv"
	"unicode"
)

func DiffWaysToCompute(expression string) []int {
	n := len(expression)
	memo := make(map[string][]int)

	var dfs func(left, right int) []int
	dfs = func(left, right int) []int {
		if val, ok := toDigit(expression[left : right+1]); ok {
			return []int{val}
		}

		if val, ok := memo[expression[left:right+1]]; ok {
			return val
		}
		ans := make([]int, 0)
		for i := left; i <= right; i++ {
			ch := expression[i]
			if unicode.IsDigit(rune(ch)) {
				continue
			}
			ls, rs := dfs(left, i-1), dfs(i+1, right)
			val := 0
			for _, l := range ls {
				for _, r := range rs {
					switch ch {
					case '+':
						val = l + r
					case '-':
						val = l - r
					case '*':
						val = l * r
					}
					ans = append(ans, val)
				}
			}
		}
		memo[expression[left:right+1]] = ans
		return ans
	}

	return dfs(0, n-1)
}

func toDigit(exp string) (int, bool) {
	atoi, err := strconv.Atoi(exp)
	if err != nil {
		return -1, false
	}
	return atoi, true
}

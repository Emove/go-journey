package _7_letter_combination_of_a_phone_number

import (
	"strings"
)

func LetterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	dic, n, ans, builder := getDic(), len(digits), make([]string, 0), strings.Builder{}
	var backtrack func(index int)
	backtrack = func(index int) {
		if index == n {
			ans = append(ans, builder.String())
		} else {
			chs := dic[digits[index]-48]
			for _, v := range chs {
				builder.WriteRune(v)
				backtrack(index + 1)
				temp := builder.String()[:builder.Len()-1]
				builder.Reset()
				builder.WriteString(temp)
			}
		}
	}
	backtrack(0)
	return ans
}

func LetterCombinations1(digits string) []string {
	if digits == "" {
		return []string{}
	}
	dic, n, ans, stk := getDic(), len(digits), make([]string, 0), []string{}
	for _, v := range dic[digits[0]-48] {
		stk = append(stk, string(v))
	}
	for len(stk) > 0 {
		sn := len(stk) - 1
		prefix := stk[sn]
		stk = stk[:sn]
		pn := len(prefix)
		if pn == n {
			ans = append(ans, prefix)
			continue
		}
		for _, v := range dic[digits[pn]-48] {
			if len(prefix) == n-1 {
				ans = append(ans, prefix+string(v))
			} else {
				stk = append(stk, prefix+string(v))
			}
		}
	}
	return ans
}

func LetterCombinations2(digits string) []string {
	if digits == "" {
		return []string{}
	}
	dic, n := getDic(), len(digits)
	dp := make([][]string, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]string, 0)
	}
	for _, v := range dic[digits[0]-48] {
		dp[0] = append(dp[0], string(v))
	}
	for i := 1; i < n; i++ {
		for _, prefix := range dp[i-1] {
			for _, v := range dic[digits[i]-48] {
				dp[i] = append(dp[i], prefix+string(v))
			}
		}
	}
	return dp[n-1]
}

func getDic() []string {
	dic := make([]string, 10)
	dic[2] = "abc"
	dic[3] = "def"
	dic[4] = "ghi"
	dic[5] = "jkl"
	dic[6] = "mno"
	dic[7] = "pqrs"
	dic[8] = "tuv"
	dic[9] = "wxyz"
	return dic
}

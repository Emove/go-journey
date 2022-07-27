package _8_string_to_integer_atoi

import "math"

func MyAtoi(s string) int {
	n := len(s)
	neg, ans := 1, 0
	i := 0
	for ; i < n && s[i] == ' '; i++ {
	}
	if i < n {
		if s[i] == '-' {
			neg = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}
	up, bot := math.MaxInt32/10, math.MinInt32/10
	for ; i < n && isNum(s[i]); i++ {
		cur := int(s[i] - 48)
		if neg == 1 {
			if ans > up || (ans == up && cur > 7) {
				return math.MaxInt32
			}
		} else {
			if -ans < bot || (-ans == bot && cur > 8) {
				return math.MinInt32
			}
		}
		ans = ans*10 + cur
	}
	return ans * neg
}

func isNum(ch uint8) bool {
	return ch >= 48 && ch <= 57
}

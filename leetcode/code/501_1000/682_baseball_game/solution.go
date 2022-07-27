package _82_baseball_game

import (
	"strconv"
)

func CalPoints(ops []string) int {
	n := len(ops)
	cur, _ := strconv.Atoi(ops[0])
	res := []int{cur}
	last, ans := 0, cur
	for i := 1; i < n; i++ {
		v := ops[i]
		switch v {
		case "+":
			res = append(res, res[last]+res[last-1])
			ans += res[last] + res[last-1]
			last++
		case "D":
			res = append(res, res[last]*2)
			ans += res[last] * 2
			last++
		case "C":
			ans -= res[last]
			res = res[:last]
			last--
		default:
			cur, _ := strconv.Atoi(v)
			res = append(res, cur)
			ans += cur
			last++
		}
	}
	return ans
}

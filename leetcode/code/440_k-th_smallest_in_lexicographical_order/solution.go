package _40_k_th_smallest_in_lexicographical_order

import (
	"math"
	"strconv"
)

func FindKthNumber(n, k int) int {
	ans := 1
	cnt := 0
	for k > 1 {
		cnt = getCnt(ans, n)
		if cnt < k {
			k -= cnt
			ans++
		} else {
			k--
			ans *= 10
		}
	}
	return ans
}

func getCnt(x, limit int) int {
	a, b := strconv.Itoa(x), strconv.Itoa(limit)
	n := len(a)
	k := len(b) - n
	res := 0
	for i := 0; i < k; i++ {
		res += int(math.Pow(float64(10), float64(i)))
	}
	u, _ := strconv.Atoi(b[:n])
	if u > x {
		res += int(math.Pow(float64(10), float64(k)))
	} else if u == x {
		res += limit - x*int(math.Pow(float64(10), float64(k))) + 1
	}
	return res
}

func FindKthNumber1(n, k int) int {
	ans := 1
	k--
	for k > 0 {
		steps := getSteps(ans, n)
		if steps <= k {
			k -= steps
			ans++
		} else {
			k--
			ans *= 10
		}
	}
	return ans
}

func getSteps(x, limit int) int {
	steps := 0
	first, last := x, x
	for first <= limit {
		steps += min(last, limit) - first + 1
		first *= 10
		last = first + 9
	}
	return steps
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

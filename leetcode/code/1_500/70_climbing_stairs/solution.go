package _0_climbing_stairs

import "math"

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	p, q, r := 0, 1, 2
	for i := 3; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func ClimbStairs1(n int) int {
	sqrt5 := math.Sqrt(5)
	pow1 := math.Pow((1+sqrt5)/2, float64(n+1))
	pow2 := math.Pow((1-sqrt5)/2, float64(n+1))
	return int(math.Round((pow1 - pow2) / sqrt5))
}

package _0_powx_n

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return mul(x, n)
	}
	return 1.0 / mul(x, -n)
}

func mul(x float64, n int) float64 {
	ans := 1.0
	for n > 0 {
		if n%2 == 1 {
			ans *= x
		}
		x *= x
		n /= 2
	}
	return ans
}

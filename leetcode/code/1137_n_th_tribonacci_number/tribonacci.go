package n_th_tribonacci_number_1137

func Tribonacci(n int) int {
	pres := []int{0, 1, 1}
	if n <= 2 {
		return pres[n]
	}
	for i := 3; i < n; i++ {
		pres[i%3] = pres[0] + pres[1] + pres[2]
	}
	return pres[0] + pres[1] + pres[2]
}

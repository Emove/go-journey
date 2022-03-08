package _055_plates_between_candles

func PlatesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	q, p := -1, -1
	l := make([]int, n)
	r := make([]int, n)
	sum := make([]int, n+1)
	for i, j := 0, n-1; i < n; i++ {
		if s[i] == '|' {
			p = i
		}
		if s[j] == '|' {
			q = j
		}
		l[i] = p
		r[j] = q
		if s[i] == '*' {
			sum[i+1] = sum[i] + 1
		} else {
			sum[i+1] = sum[i]
		}
		j--
	}

	qn := len(queries)
	ans := make([]int, qn)
	for i := 0; i < qn; i++ {
		x := r[queries[i][0]]
		y := l[queries[i][1]]
		if x != -1 && x <= y {
			ans[i] = sum[y+1] - sum[x]
		}

	}

	return ans
}

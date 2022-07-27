package _68_kth_smallest_number_in_multiplication_table

import "sort"

func FindKthNumber(m, n, k int) int {
	_m, _n := min(m, n), max(m, n)
	left, right := 1, _m*_n
	for left < right {
		mid := (left + right) >> 1
		cnt := 0
		for i := 1; i <= n; i++ {
			if i*m <= mid {
				cnt += m
			} else {
				cnt += mid / i
			}
		}
		if cnt >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

//getCnt := func(mid int) int {
//	ans := 0
//	for i := 1; i <= n; i++ {
//		if i * m <= mid {
//			ans += m
//		} else {
//			ans += mid/i
//		}
//	}
//	return ans
//}

func FindKthNumber1(m, n, k int) int {
	return sort.Search(m*n, func(x int) bool {
		count := x / n * n
		for i := x/n + 1; i <= m; i++ {
			count += x / i
		}
		return count >= k
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package _29_consecutive_numbers_sum

func ConsecutiveNumbersSum(n int) (ans int) {
	for k := 1; k*(k+1) <= n*2; k++ {
		if isKConsecutive(n, k) {
			ans++
		}
	}
	return
}

func isKConsecutive(n, k int) bool {
	if k%2 == 1 {
		return n%k == 0
	}
	return n%k != 0 && 2*n%k == 0
}

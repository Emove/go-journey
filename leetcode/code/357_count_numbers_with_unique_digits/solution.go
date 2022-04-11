package _57_count_numbers_with_unique_digits

var t = []int{1, 10, 91, 739, 5275, 32491, 168571, 712891, 2345851, 5611771, 8877691}

func countNumbersWithUniqueDigits(n int) int {
	//return t[n]
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	ans, cur := 10, 9
	for i := 0; i < n-1; i++ {
		cur *= 9 - i
		ans += cur
	}
	return ans
}

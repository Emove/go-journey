package _28_self_dividing_number

func SelfDividingNumbers(left, right int) []int {
	ans := make([]int, 0)
	index := left
	for i := left; i < 10 && i <= right; i++ {
		ans = append(ans, i)
		index = i
	}
	for i := index + 1; i <= right; i++ {
		if isSelfDividingNumber(i) {
			ans = append(ans, i)
		}
	}
	return ans
}

func isSelfDividingNumber(i int) bool {
	for num := i; num > 0; num /= 10 {
		if div := num % 10; div == 0 || i%div != 0 {
			return false
		}
	}
	return true
}

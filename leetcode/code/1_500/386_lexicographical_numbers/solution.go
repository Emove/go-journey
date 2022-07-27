package _86_lexicographical_numbers

func LexicalOrder(n int) []int {
	res := make([]int, n)

	num := 1
	for i := range res {
		res[i] = num

		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}

	return res
}

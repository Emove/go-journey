package _93_binary_number_with_alternating_bits

func HasAlternatingBits(n int) bool {
	for pre := 2; n != 0; n /= 2 {
		cur := n % 2
		if cur == pre {
			return false
		}
		pre = cur
	}
	return true
}

func solution2(n int) bool {
	// 当且仅当n的二进制数为交替位是，a全为1
	a := n ^ (n >> 1)
	return (a & (a + 1)) == 0
}

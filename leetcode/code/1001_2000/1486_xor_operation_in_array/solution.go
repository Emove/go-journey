package xor_operation_in_array_1486

func xorOperation(n int, start int) int {
	//var result int
	//for i := 0; i < n; i++ {
	//	result ^= start + 2*i
	//}
	//return result
	b0 := n & start & 1
	s := start / 2
	res := computeXor(s-1) ^ computeXor(s+n-1)
	return res<<1 + b0
}

func computeXor(n int) int {
	switch n % 4 {
	case 0:
		return n
	case 1:
		return 1
	case 2:
		return n + 1
	}
	return 0
}

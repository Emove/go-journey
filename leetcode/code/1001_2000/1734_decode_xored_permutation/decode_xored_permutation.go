package decode_xored_permutation_1734

func decode(encoded []int) []int {
	length := len(encoded) + 1
	result := make([]int, length)
	total := 0
	odd := 0
	sign := 0
	go func() {
		for i := 0; i < length-1; i += 2 {
			odd ^= encoded[i]
		}
		sign = 1
	}()
	for i := 1; i <= length; i++ {
		total ^= i
	}
	for sign != 1 {

	}
	result[length-1] = total ^ odd
	for i := length - 2; i >= 0; i-- {
		result[i] = encoded[i] ^ result[i+1]
	}
	return result
}

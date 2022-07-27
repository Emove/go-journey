package decode_xored_array_1720

func decode(encoded []int, first int) []int {
	length := len(encoded)
	result := make([]int, length+1)
	result[0] = first

	for i := 0; i < length; i++ {
		result[i+1] = result[i] ^ encoded[i]
	}

	return result
}

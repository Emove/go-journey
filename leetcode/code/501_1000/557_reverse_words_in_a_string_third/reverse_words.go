package reverse_words_in_a_string_third_557

func reverseWords(s string) string {
	left := 0
	right := 0
	length := len(s)
	bytes := []byte(s)
	for right < length-1 {
		for bytes[right+1] != ' ' {
			right++
			if right == length-1 {
				break
			}
		}
		rightLastIndex := right
		for left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
		right = rightLastIndex + 2
		left = right
	}
	return string(bytes)
}

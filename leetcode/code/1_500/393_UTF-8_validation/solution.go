package _93_UTF_8_validation

const MASK1 = 1 << 7      // 128
const MASK2 = 1<<7 | 1<<6 // 192

func ValidUtf8(data []int) bool {
	n := len(data)
	for i := 0; i < n; {
		length := getBytes(data[i])
		if length < 0 || i+length > n {
			return false
		}
		for _, v := range data[i+1 : i+length] {
			if v&MASK2 != MASK1 {
				return false
			}
		}
		i += length
	}
	return true
}

func getBytes(num int) (n int) {
	if num&MASK1 == 0 {
		return 1
	}
	for mask := MASK1; mask&num != 0; mask >>= 1 {
		n++
		if n > 4 {
			return -1
		}
	}
	if n >= 2 {
		return n
	}
	return -1
}

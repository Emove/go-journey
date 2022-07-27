package _3_roman_to_integer

func RomanToInt(s string) int {
	sum, prev := 0, getInt(s[0])
	for i := 1; i < len(s); i++ {
		num := getInt(s[i])
		if prev < num {
			sum -= prev
		} else {
			sum += prev
		}
		prev = num
	}
	sum += prev
	return sum
}

func getInt(ch byte) int {
	switch ch {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

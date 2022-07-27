package _04_base_7

func ConvertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	negative := num < 0
	if negative {
		num = -num
	}
	var result []byte
	for num > 0 {
		result = append(result, '0'+byte(num%7))
		num /= 7
	}
	if negative {
		result = append(result, '-')
	}
	for i, n := 0, len(result); i < n/2; i++ {
		result[i], result[n-1-i] = result[n-1-i], result[i]
	}
	return string(result)
}

package _108_defanging_an_ip_address

func DefangIPaddr(address string) string {
	//return strings.ReplaceAll(address, ".", "[.]")
	ans, offset := make([]byte, len(address)+6), 0
	for _, ch := range address {
		if ch == '.' {
			ans[offset] = '['
			offset++
			ans[offset] = '.'
			offset++
			ans[offset] = ']'
		} else {
			ans[offset] = byte(ch)
		}
		offset++
	}
	return string(ans)
}

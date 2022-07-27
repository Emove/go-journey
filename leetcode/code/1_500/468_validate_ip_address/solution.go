package _68_validate_ip_address

func ValidIPAddress(queryIP string) string {
	n := len(queryIP)
	ipv4 := true
	for i := 0; i < n; i++ {
		if queryIP[i] >= ':' {
			ipv4 = false
			break
		}
	}
	if ipv4 {
		if validateIPV4(queryIP) {
			return "IPv4"
		}
	} else {
		if validateIPV6(queryIP) {
			return "IPv6"
		}
	}
	return "Neither"
}

func validateIPV4(queryIP string) bool {
	n, cnt, sum := len(queryIP), 0, 0
	if n < 7 {
		return false
	}
	for i := 0; i < n; i++ {
		if queryIP[i] == '.' {
			if i == n-1 || queryIP[i+1] == '.' {
				return false
			}
			cnt++
			sum = 0
			continue
		}
		if sum == 0 && queryIP[i] == '0' && i < n-1 && queryIP[i+1] != '.' {
			return false
		}
		val := int(queryIP[i] - 48)
		if 0 > val || val > 9 {
			return false
		}
		sum = sum*10 + val
		if sum > 255 {
			return false
		}
	}
	return cnt == 3
}

func validateIPV6(queryIP string) bool {
	n, cnt, interval := len(queryIP), 0, 0
	if n < 15 {
		return false
	}
	for i := 0; i < n; i++ {
		ch := queryIP[i]
		if ch == ':' {
			if i == n-1 || queryIP[i+1] == ':' {
				return false
			}
			cnt++
			interval = 0
			continue
		}
		interval++
		if interval < 1 || interval > 4 {
			return false
		}
		if ch < '0' {
			return false
		}
		if ch > '9' && ch < 'A' {
			return false
		}
		if ch > 'F' && ch < 'a' {
			return false
		}
		if ch > 'f' {
			return false
		}
	}
	return cnt == 7
}

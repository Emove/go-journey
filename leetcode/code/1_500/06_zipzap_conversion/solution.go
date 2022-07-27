package zipzap_conversion_06

func Convert(s string, numRows int) string {
	n := len(s)
	r := numRows
	if r == 1 || r >= n {
		return s
	}
	t := 2*r - 2
	ans := make([]byte, n)
	//index:=0
	for i := 0; i < r; i++ {
		for j := 0; j+i < n; j += t {
			ans = append(ans, s[j+i])
			//ans[index] = s[j+i]
			//index++
			if i > 0 && i < r-1 && j+t-i < n {
				//ans[index] = s[j+t-i]
				ans = append(ans, s[j+t-i])
				//index++
			}
		}
	}
	return string(ans)
}

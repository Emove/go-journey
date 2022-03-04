package letter_case_permutation_784

var res []string

func letterCasePermutation(s string) []string {
	n := len(s)
	chs := []byte(s)
	res = make([]string, 0)
	dfs(chs, n, 0)
	return res
}

func dfs(chs []byte, n, k int) {
	res = append(res, string(chs))
	for i := k; i < n; i++ {
		temp := chs[i]
		if isLetter(temp) {
			if temp >= 97 {
				chs[i] = temp - 32
			} else {
				chs[i] = temp + 32
			}
			dfs(chs, n, i+1)
			chs[i] = temp
		}

	}
}

func isLetter(ch uint8) bool {
	return (ch >= 65 && ch <= 90) ||
		(ch >= 97 && ch <= 122)
}

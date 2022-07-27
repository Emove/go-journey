package permutation_in_string_567

// 差分+计数
func checkInclusion(s1, s2 string) bool {
	n1 := len(s1)
	// s1中的字符计数
	cnt := 0
	feq := make([]int, 26)
	for _, c := range s1 {
		index := c - 'a'
		if feq[index] == 0 {
			cnt++
		}
		feq[index]++
	}
	n2 := len(s2)
	for i := 0; i < n2; i++ {
		index := s2[i] - 'a'
		feq[index]--
		if feq[index] == 0 {
			cnt--
		}
		if i >= n1 {
			theFirstIndexOfSubstring := s2[i-n1] - 'a'
			if feq[theFirstIndexOfSubstring] == 0 {
				// 如果最左边的左移一位是s1中的字符
				// 说明当前i-n1所构成的子串不符合s1的排列
				cnt++
			}
			feq[theFirstIndexOfSubstring]++
		}
		if cnt == 0 {
			return true
		}
	}
	return false
}

package _1_05_one_away_lcci

func OneEditAway(first, second string) bool {
	n, m := len(first), len(second)
	// 让n恒大于m，first恒长于second
	if n < m {
		n, m = m, n
		first, second = second, first
	}
	if n-m > 1 {
		return false
	}
	cnt := 0
	for i, j := 0, 0; i < n && j < m; {
		if first[i] == second[j] {
			i++
			j++
		} else {
			if n == m {
				i++
				j++
			} else {
				i++
			}
			cnt++
			if cnt > 1 {
				return false
			}
		}
	}
	return cnt <= 1
}

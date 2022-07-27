package _893_check_if_all_the_integers_in_a_range_are_covered

func isCovered(ranges [][]int, left int, right int) bool {
	// 差分数组
	diff := make([]int, 52, 52)
	for _, ran := range ranges {
		diff[ran[0]]++
		diff[ran[1]+1]--
	}
	// 计算差分数组前缀和
	for i := 1; i <= 51; i++ {
		diff[i] = diff[i-1] + diff[i]
	}
	for i := left; i <= right; i++ {
		if diff[i] <= 0 {
			return false
		}
	}
	return true
}

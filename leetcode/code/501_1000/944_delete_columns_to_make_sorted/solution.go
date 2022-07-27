package _44_delete_columns_to_make_sorted

func MinDeletionSize(strs []string) int {
	ans := 0
	rows, columns := len(strs), len(strs[0])

	for i := 0; i < columns; i++ {
		for j := 1; j < rows; j++ {
			if strs[j][i] < strs[j-1][i] {
				ans++
				break
			}
		}
	}

	return ans
}

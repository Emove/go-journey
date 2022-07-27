package _252_cells_with_odd_values_in_a_matrix

func OddCells(m int, n int, indices [][]int) int {
	rows := make([]int, m)
	cols := make([]int, n)
	for _, index := range indices {
		rows[index[0]]++
		cols[index[1]]++
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans += (rows[i] + cols[j]) & 1
		}
	}

	return ans
}

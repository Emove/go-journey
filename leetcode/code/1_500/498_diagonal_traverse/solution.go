package _98_diagonal_traverse

func FindDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	idx, ans := 0, make([]int, n*m)
	row, col := 0, 0
	for i := 0; i < m+n-1; i++ {
		if i&1 == 0 {
			row, col = i, 0
			if i >= m {
				row = m - 1
				col = i - m + 1
			}
			for row >= 0 && col < n {
				ans[idx] = mat[row][col]
				idx++
				row--
				col++
			}
		} else {
			row, col = 0, i
			if i >= n {
				col = n - 1
				row = i - n + 1
			}
			for col >= 0 && row < m {
				ans[idx] = mat[row][col]
				idx++
				row++
				col--
			}
		}
	}
	return ans
}

package matrix_01_542

import "math"

func UpdateMatrix(mat [][]int) [][]int {
	depth := len(mat)
	width := len(mat[0])

	dist := make([][]int, depth)
	for i := 0; i < depth; i++ {
		dist[i] = make([]int, width)
	}
	for i := 0; i < depth; i++ {
		for j := 0; j < width; j++ {
			if mat[i][j] != 0 {
				dist[i][j] = math.MaxUint32
			}
		}
	}

	// 从左上角开始计算，每个点只计算水平左方向和垂直上方向
	for i := 0; i < depth; i++ {
		for j := 0; j < width; j++ {
			if i-1 >= 0 {
				dist[i][j] = min(dist[i][j], dist[i-1][j]+1)
			}
			if j-1 >= 0 {
				dist[i][j] = min(dist[i][j], dist[i][j-1]+1)
			}
		}
	}

	// 从右下角开始计算，每个点只计算水平右方向和垂直下方向
	for i := depth - 1; i >= 0; i-- {
		for j := width - 1; j >= 0; j-- {
			if i+1 < depth {
				dist[i][j] = min(dist[i][j], dist[i+1][j]+1)
			}
			if j+1 < width {
				dist[i][j] = min(dist[i][j], dist[i][j+1]+1)
			}
		}
	}
	return dist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

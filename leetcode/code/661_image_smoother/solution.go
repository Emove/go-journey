package _61_image_smoother

func ImageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			// 以点[i,j]为中心
			// 平移矩阵中的有效点数和有效点数的值和
			cnt, sum := 0, 0
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x >= 0 && x < m && y >= 0 && y < n {
						sum += img[x][y]
						cnt++
					}
				}
			}
			ans[i][j] = sum / cnt
		}
	}
	return ans
}

func ImageSmoother1(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	ans := make([][]int, m)
	for i := range img {
		ans[i] = make([]int, n)
		for j := range img[1] {
			cnt, sum := 0, 0
			for _, row := range img[max(i-1, 0):min(i+2, m)] {
				for _, v := range row[max(j-1, 0):min(j+2, n)] {
					sum += v
					cnt++
				}
			}
			ans[i][j] = sum / cnt
		}
	}
	return ans
}

func ImageSmoother2(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	ans, sum := make([][]int, m), make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		sum[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + img[i-1][j-1]
		}
	}

	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			leftTopX, leftTopY := max(0, i-1), max(0, j-1)
			rightBottomX, rightBottomY := min(m-1, i+1), min(n-1, j+1)
			cnt := (rightBottomX - leftTopX + 1) * (rightBottomY - leftTopY + 1)
			total := sum[rightBottomX+1][rightBottomY+1] - sum[leftTopX][rightBottomY+1] - sum[rightBottomX+1][leftTopY] + sum[leftTopX][leftTopY]
			ans[i][j] = total / cnt
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

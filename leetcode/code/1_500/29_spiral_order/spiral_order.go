package spiral_order_29

// 顺时针打印矩阵
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows := len(matrix)
	columns := len(matrix[0])
	start := 0
	result := make([]int, 0)
	for rows > start<<1 && columns > start<<1 {
		appendMatrixInCircle(matrix, result, columns, rows, start)
		start++
	}
	return result
}

func appendMatrixInCircle(matrix [][]int, result []int, columns, rows, start int) {
	// 从左到右打印
	endX := columns - 1 - start
	endY := rows - 1 - start
	if start <= endX {
		for i := start; i <= endX; i++ {
			result = append(result, matrix[start][i])
		}
	}
	// 从上到下打印
	if start <= endY {
		for i := start + 1; i <= endY; i++ {
			result = append(result, matrix[i][endX])
		}
	}
	// 从右到左打印
	if endX > start && endY > start {
		for i := endX - 1; i >= start; i-- {
			result = append(result, matrix[endY][i])
		}
	}
	// 从下到上打印
	if endX > start && endY > start+1 {
		for i := endY - 1; i > start; i-- {
			result = append(result, matrix[i][start])
		}
	}
}

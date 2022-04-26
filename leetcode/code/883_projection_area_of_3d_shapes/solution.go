package _83_projection_area_of_3d_shapes

func ProjectionArea(grid [][]int) int {
	var xyArea, yzArea, zxArea int
	for i, row := range grid {
		yzHeight, zxHeight := 0, 0
		for j, v := range row {
			if v > 0 {
				xyArea++
			}
			yzHeight = max(yzHeight, grid[j][i])
			zxHeight = max(zxHeight, v)
		}
		yzArea += yzHeight
		zxArea += zxHeight
	}
	return xyArea + yzArea + zxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

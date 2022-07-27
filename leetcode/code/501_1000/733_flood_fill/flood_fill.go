package flood_fill_733

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor != newColor {
		dfs(image, sr, sc, newColor, oldColor, len(image), len(image[0]))
	}
	return image
}

func dfs(image [][]int, sr, sc, newColor, oldColor, depth, width int) {
	if sr < 0 || sc < 0 || sr == depth || sc == width || image[sr][sc] != oldColor {
		return
	}

	image[sr][sc] = newColor
	dfs(image, sr+1, sc, newColor, oldColor, depth, width)
	dfs(image, sr-1, sc, newColor, oldColor, depth, width)
	dfs(image, sr, sc+1, newColor, oldColor, depth, width)
	dfs(image, sr, sc-1, newColor, oldColor, depth, width)
}

package _12_largest_triangle_area

import (
	"math"
)

func LargestTriangleArea(points [][]int) float64 {
	ans, n := float64(0), len(points)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := i + 2; k < n; k++ {
				ans = max(ans, calculateArea(points[i], points[j], points[k]))
			}
		}
	}
	return ans
}

func calculateArea(p1, p2, p3 []int) float64 {
	return 0.5 * math.Abs(float64(p1[0]*p2[1]+p2[0]*p3[1]+p3[0]*p1[1]-p1[0]*p3[1]-p2[0]*p1[1]-p3[0]*p2[1]))
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

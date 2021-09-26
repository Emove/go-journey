package number_of_boomerangs_447

func numberOfBoomerangs(points [][]int) int {
	length := len(points)
	ans := 0
	for _, point := range points {
		cnt := map[int]int{}
		for i := 0; i < length; i++ {
			x := point[0] - points[i][0]
			y := point[1] - points[i][1]
			dis := x*x + y*y
			cnt[dis]++
		}
		for _, m := range cnt {
			ans += m * (m - 1)
		}
	}
	return ans
}

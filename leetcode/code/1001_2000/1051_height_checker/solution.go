package _051_height_checker

func HeightChecker(heights []int) int {
	cnt := [101]int{}
	for _, height := range heights {
		cnt[height]++
	}
	ans, idx := 0, 0
	for height, count := range cnt {
		for ; count > 0; count-- {
			if heights[idx] != height {
				ans++
			}
			idx++
		}
	}
	return ans
}

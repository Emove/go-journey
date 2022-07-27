package _35_asteroid_collision

func AsteroidCollision(asteroids []int) []int {
	ans, alive := make([]int, 0), true
	for _, aster := range asteroids {
		alive = true
		if aster < 0 {
			for alive && len(ans) > 0 && ans[len(ans)-1] > 0 {
				lastOne := ans[len(ans)-1]
				if lastOne > -aster {
					alive = false
				} else if lastOne == -aster {
					alive = false
					ans = ans[:len(ans)-1]
				} else {
					ans = ans[:len(ans)-1]
				}
			}
		}
		if alive {
			ans = append(ans, aster)
		}
	}
	return ans
}

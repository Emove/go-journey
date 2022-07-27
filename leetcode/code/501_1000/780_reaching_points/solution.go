package _80_reaching_points

// ReachingPoints https://leetcode-cn.com/problems/reaching-points/
func ReachingPoints(sx, sy, tx, ty int) bool {
	for sx < tx && sy < ty {
		if tx < ty {
			ty %= tx
		} else {
			tx %= ty
		}
	}

	if tx < sx || ty < sy {
		return false
	}

	if sx == tx {
		return (ty-sy)%tx == 0
	}
	return (tx-sx)%ty == 0
}

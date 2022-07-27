package _217_minimum_cost_to_move_chips_to_the_same_position

func MinCostToMoveChips(position []int) int {
	cnt := make([]int, 2)
	for _, pos := range position {
		cnt[pos&1]++
	}

	if cnt[0] < cnt[1] {
		return cnt[0]
	}
	return cnt[1]
}

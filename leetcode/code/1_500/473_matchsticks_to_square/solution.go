package _73_matchsticks_to_square

import (
	"sort"
)

func MakeSquare(matchsticks []int) bool {
	n := len(matchsticks)
	if n < 4 {
		return false
	}
	sum, max := 0, 0
	for i := 0; i < n; i++ {
		sum += matchsticks[i]
		if matchsticks[i] > max {
			max = matchsticks[i]
		}
	}
	edgeLength := sum / 4
	if sum%4 != 0 || max > edgeLength {
		return false
	}

	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})

	edges := make([]int, 4)

	var backtrack func(index int) bool
	backtrack = func(index int) bool {
		if index == n {
			return true
		}

		for i := 0; i < len(edges); i++ {
			edges[i] += matchsticks[index]
			if edges[i] <= edgeLength && backtrack(index+1) {
				return true
			}
			edges[i] -= matchsticks[index]
		}
		return false
	}

	return backtrack(0)
}

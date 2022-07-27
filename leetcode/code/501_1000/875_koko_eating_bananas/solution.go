package _75_koko_eating_bananas

func MinEatingSpeed(piles []int, h int) int {
	maxVal := 1
	for _, pile := range piles {
		maxVal = max(maxVal, pile)
	}
	left, right := 1, maxVal
	count := func(speed int) int {
		sum := 0
		for _, pile := range piles {
			sum += (pile + speed - 1) / speed
		}
		return sum
	}
	for left < right {
		mid := left + (right-left)>>1
		if count(mid) > h {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

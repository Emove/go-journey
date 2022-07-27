package _44_find_smallest_letter_than_target

import (
	"sort"
)

func NextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	if target >= letters[right] {
		return letters[0]
	}
	for left < right {
		mid := (left + right) >> 1
		if letters[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return letters[right]
}

func nextGreatestLetter(letters []byte, target byte) byte {
	if target >= letters[len(letters)-1] {
		return letters[0]
	}
	i := sort.Search(len(letters)-1, func(i int) bool { return letters[i] > target })
	return letters[i]
}

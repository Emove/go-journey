package _36_find_right_interval

import "sort"

func FindRightInterval(intervals [][]int) []int {
	n := len(intervals)
	clone := make([][]int, n)
	for i := 0; i < n; i++ {
		clone[i] = make([]int, 2)
		clone[i][0] = intervals[i][0]
		clone[i][1] = i
	}

	sort.Slice(clone, func(i, j int) bool {
		return clone[i][0] < clone[j][0]
	})

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		interval := intervals[i]
		left, right := 0, n-1
		for left < right {
			mid := (left + right) >> 1
			if clone[mid][0] >= interval[1] {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if clone[right][0] >= interval[1] {
			ans[i] = clone[right][1]
		} else {
			ans[i] = -1
		}
	}
	return ans
}

func FindRightInterval1(intervals [][]int) []int {
	n := len(intervals)
	startIntervals, endIntervals := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		interval := intervals[i]
		startIntervals[i] = []int{interval[0], i}
		endIntervals[i] = []int{interval[1], i}
	}

	sort.Slice(startIntervals, func(i, j int) bool {
		return startIntervals[i][0] < startIntervals[j][0]
	})
	sort.Slice(endIntervals, func(i, j int) bool {
		return endIntervals[i][0] < endIntervals[j][0]
	})

	ans := make([]int, n)
	for left, right := 0, 0; left < n; left++ {
		for right < n && endIntervals[left][0] > startIntervals[right][0] {
			right++
		}
		if right < n {
			ans[endIntervals[left][1]] = startIntervals[right][1]
		} else {
			ans[endIntervals[left][1]] = -1
		}
	}
	return ans
}

package _71_minimum_number_of_refueling_stops

import (
	"container/heap"
	"fmt"
	"sort"
)

func MinRefuelStops(target int, startFuel int, stations [][]int) int {
	n := len(stations)
	dp := make([]int, n+1)
	dp[0] = startFuel

	for i, station := range stations {
		for j := i; j >= 0; j-- {
			if dp[j] >= station[0] {
				dp[j+1] = max(dp[j+1], dp[j]+station[1])
				fmt.Printf("%d, %d, %d\n", i, j, dp[j+1])
			}
		}
	}

	for i, v := range dp {
		if v >= target {
			return i
		}
	}

	return -1
}

func MinRefuelStopsGreedy(target int, startFuel int, stations [][]int) int {
	fuel, prev, h := startFuel, 0, &hp{}
	ans := 0
	for i, n := 0, len(stations); i <= n; i++ {
		curr := target
		if i < n {
			curr = stations[i][0]
		}
		fuel -= curr - prev
		for fuel < 0 && h.Len() > 0 {
			fuel += heap.Pop(h).(int)
			ans++
		}
		if fuel < 0 {
			return -1
		}
		if i < n {
			heap.Push(h, stations[i][1])
			prev = curr
		}
	}
	return ans
}

type hp struct {
	sort.IntSlice
}

func (hp *hp) Less(i, j int) bool {
	return hp.IntSlice[i] > hp.IntSlice[j]
}

func (hp *hp) Push(val interface{}) {
	hp.IntSlice = append(hp.IntSlice, val.(int))
}

func (hp *hp) Pop() interface{} {
	n := len(hp.IntSlice) - 1
	v := hp.IntSlice[n]
	hp.IntSlice = hp.IntSlice[:n]
	return v
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

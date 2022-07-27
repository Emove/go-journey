package _99_minimum_index_sum_of_two_lists

import (
	"math"
)

func FindRestaurant(list1, list2 []string) (ans []string) {
	n1, n2 := len(list1), len(list2)
	var m map[string]int
	var foreachList []string
	if n1 <= n2 {
		m = make(map[string]int, n1)
		foreachList = list2
		for i, v := range list1 {
			m[v] = i
		}
	} else {
		m = make(map[string]int, n2)
		foreachList = list1
		for i, v := range list2 {
			m[v] = i
		}
	}

	sum := math.MaxInt32
	for i, v := range foreachList {

		if i > sum {
			break
		}

		// 存在且索引和小于等于sum
		if index, ok := m[v]; ok {
			temp := index + i
			if temp == sum {
				ans = append(ans, v)
			} else if temp < sum {
				sum = temp
				ans = []string{v}
			}
		}

	}

	return
}

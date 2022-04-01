package _54_array_of_doubled_pairs

import "sort"

func CanRecorderDoubled(arr []int) bool {
	n := len(arr)
	index := make(map[int]int, n)
	sort.Slice(arr, func(i, j int) bool {
		return abs(arr[i]) < abs(arr[j])
	})
	zeros, count := 0, 0
	for i := 0; i < n; i++ {
		cur := arr[i]
		if cur == 0 {
			zeros++
			continue
		}
		req := index[cur]
		if req > 0 {
			index[cur]++
			count++
		} else {
			div := cur >> 1
			mul := cur << 1
			if cnt, ok := index[mul]; ok && cnt > 0 {
				index[mul]--
				count--
			} else if cur&1 == 0 {
				// 偶数
				if cnt, ok := index[div]; ok && cnt > 0 {
					index[div]--
					count--
				} else {
					index[cur]++
					count++
				}
			} else {
				index[cur]++
				count++
			}
		}
	}
	return zeros&1 == 0 && count == 0
}

func CanRecorderDoubled1(arr []int) bool {
	cnt := make(map[int]int, len(arr))
	for _, x := range arr {
		cnt[x]++
	}
	if cnt[0]&1 == 1 {
		return false
	}

	// 排序
	vals := make([]int, 0, len(cnt))
	for x := range cnt {
		vals = append(vals, x)
	}
	sort.Slice(vals, func(i, j int) bool { return abs(vals[i]) < abs(vals[j]) })

	for _, x := range vals {
		if cnt[2*x] < cnt[x] { // 无法找到足够的 2x 与 x 配对
			return false
		}
		cnt[2*x] -= cnt[x]
	}
	return true
}

// continue有性能消耗
func ImplWithContinue(arr []int) bool {
	n := len(arr)
	index := make(map[int]int, n)
	sort.Slice(arr, func(i, j int) bool {
		return abs(arr[i]) < abs(arr[j])
	})
	zeros, count := 0, 0
	for i := 0; i < n; i++ {
		cur := arr[i]
		if cur == 0 {
			zeros++
			continue
		}
		req := index[cur]
		if req > 0 {
			index[cur]++
			count++
		} else {
			div := cur >> 1
			mul := cur << 1
			if cnt, ok := index[mul]; ok && cnt > 0 {
				index[mul]--
				count--
				continue
			} else if cur&1 == 0 {
				// 偶数
				if cnt, ok := index[div]; ok && cnt > 0 {
					index[div]--
					count--
					continue
				}
			}
			index[cur]++
			count++
		}
	}
	return zeros&1 == 0 && count == 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

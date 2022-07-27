package _001_good_day_to_rob_bank

func GoodDayToRobBank(security []int, time int) []int {
	n := len(security)
	left := make([]int, n)
	right := make([]int, n)
	left[0], right[0] = 0, 0
	for i := 1; i < n; i++ {
		if security[i-1] >= security[i] {
			// 非递增
			left[i] = left[i-1] + 1
		}
		if security[n-i-1] <= security[n-i] {
			// 从后往前算
			right[n-i-1] = right[n-i] + 1
		}
	}
	ans := make([]int, 0)
	for i := 0; i < n-time; i++ {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}
	return ans
}

func GoodDayToRobBank1(security []int, time int) []int {
	n := len(security)
	if time == 0 {
		ans := make([]int, n)
		for i := 0; i < n; i++ {
			ans[i] = i
		}
		return ans
	}
	left, right := 0, 0
	ans := make([]int, 0)
	for i := 1; i < n-time; i++ {
		if security[i-1] >= security[i] {
			left++
		} else {
			left = 0
		}
		if security[i+time-1] <= security[i+time] {
			right++
		} else {
			right = 0
		}
		if left >= time && right >= time {
			ans = append(ans, i)
		}
	}
	return ans
}

package cost

func Cost(arr []int) int {
	n := len(arr)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		cost[i] = 1
	}

	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] && cost[i] <= cost[i-1] {
			cost[i] = cost[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] && cost[i] <= cost[i+1] {
			cost[i] = cost[i+1] + 1
		}
	}
	ans := 0
	for _, val := range cost {
		ans += val
	}
	return ans
}

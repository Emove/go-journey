package _672_richest_customer_wealth

func MaximumWealth(accounts [][]int) int {
	max := 0
	for i := range accounts {
		sum := 0
		for j := range accounts[i] {
			sum += accounts[i][j]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

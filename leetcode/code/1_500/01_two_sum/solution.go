package _1_two_sum

func TwoSum(nums []int, target int) []int {
	n := len(nums)
	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		v := nums[i]
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return []int{}
}

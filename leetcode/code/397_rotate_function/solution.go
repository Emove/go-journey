package _97_rotate_function

func MaxRotateFunction(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	sum, res := 0, make([]int, n)
	for i, num := range nums {
		sum += num
		res[0] += i * num
	}
	max := res[0]
	for i := 1; i < n; i++ {
		res[i] = res[i-1] + sum - n*nums[n-i]
		if res[i] > max {
			max = res[i]
		}
	}

	return max
}

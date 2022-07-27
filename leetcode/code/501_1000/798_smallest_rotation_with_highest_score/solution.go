package _98_smallest_rotation_with_highest_score

func BestRotation(nums []int) int {
	n := len(nums)
	diff := make([]int, n+1, n+1)
	for i := 0; i < n; i++ {
		left := (i - (n - 1) + n) % n
		right := (i - nums[i] + n) % n
		if left <= right {
			diff[left] += 1
			diff[right+1] -= 1
		} else {
			diff[0] += 1
			diff[right+1] -= 1
			diff[left] += 1
			diff[n] -= 1
		}
	}
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}
	ans := 0
	for i := 1; i < n; i++ {
		if diff[i] > diff[ans] {
			ans = i
		}
	}
	return ans
}

func BestRotation1(nums []int) int {
	n := len(nums)
	score := make([]int, n)
	for i := 0; i < n; i++ {
		score[(i-nums[i]+1+n)%n] -= 1
	}
	maxIdx := 0
	for i := 1; i < n; i++ {
		score[i] += score[i-1] + 1
		if score[i] > score[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}

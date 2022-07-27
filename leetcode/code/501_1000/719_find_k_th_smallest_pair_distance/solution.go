package _19_find_k_th_smallest_pair_distance

import "sort"

// SmallestDistancePair 排序+二分+双指针
func SmallestDistancePair(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	left, right := 0, nums[n-1]-nums[0]
	for left <= right {
		mid := left + (right-left)/2
		i, cnt := 0, 0
		for j := 0; j < n; j++ {
			for nums[j]-nums[i] > mid {
				i++
			}
			cnt += j - i
		}
		if cnt >= k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

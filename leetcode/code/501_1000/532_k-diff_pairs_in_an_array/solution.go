package _32_k_diff_pairs_in_an_array

import (
	"sort"
)

func FindPairs(nums []int, k int) int {
	sort.Ints(nums)
	n, ans := len(nums), 0
	target, left, right := 0, 0, n-1
	for i := 0; i < n; i++ {
		target, left, right = k+nums[i], i+1, n-1
		for left <= right {
			mid := (right + left) >> 1
			if nums[mid] == target {
				ans++
				break
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}

func FindPairs1(nums []int, k int) int {
	res, visited := map[int]struct{}{}, map[int]struct{}{}
	for _, num := range nums {
		if _, ok := visited[num-k]; ok {
			res[num-k] = struct{}{}
		}
		if _, ok := visited[num+k]; ok {
			res[num] = struct{}{}
		}
		visited[num] = struct{}{}
	}
	return len(res)
}

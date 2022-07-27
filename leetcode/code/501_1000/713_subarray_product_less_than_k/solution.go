package _13_subarray_product_less_than_k

func NumSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	n, ans, cur := len(nums), 0, 1
	for i, j := 0, 0; j < n; j++ {
		cur *= nums[j]
		for ; cur >= k; i++ {
			cur /= nums[i]
		}
		ans += j - i + 1
	}
	return ans
}

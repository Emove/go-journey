package _4_median_of_two_sorted_arrays

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	size := n + m
	mid := size / 2
	k1, k2 := 0, 0
	for left, right, i := 0, 0, 0; i <= mid; i++ {
		temp := k2
		if left < n && right < m {
			if nums1[left] <= nums2[right] {
				k2 = nums1[left]
				left++
			} else {
				k2 = nums2[right]
				right++
			}
		} else if left < n {
			k2 = nums1[left]
			left++
		} else if right < m {
			k2 = nums2[right]
			right++
		}
		k1 = temp
	}
	if size&1 == 1 {
		return float64(k2)
	}
	return (float64(k1) + float64(k2)) / 2
}

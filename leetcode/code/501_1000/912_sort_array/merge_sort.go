/**
 * @author Emove
 * @date 2021/1/28
 */
package sort_array_912

const insertionSortThreshold = 7

// 合并排序 O(NlogN) O(N)
func MergeSort(nums []int) []int {
	length := len(nums)
	temp := make([]int, length)
	return mergeSort(nums, temp, 0, length-1)
}

func mergeSort(nums, temp []int, left, right int) []int {
	if right-left < insertionSortThreshold {
		// 数量较少，使用插入排序
		return insertionSort(nums)
	}

	mid := left + (right-left)/2

	mergeSort(nums, temp, left, mid)
	mergeSort(nums, temp, mid+1, right)
	if nums[mid] <= nums[mid+1] {
		return nums
	}
	mergeTwoSortedArray(nums, temp, left, mid, right)
	return nums
}

// 合并两个有序数组
func mergeTwoSortedArray(arr1, arr2 []int, left, mid, right int) {
	i := left
	j := mid + 1
	for k := left; k <= right; k++ {
		if i == mid+1 {
			arr1[k] = arr2[j]
		} else if j == right+1 {
			arr1[k] = arr2[i]
			i++
		} else if arr2[i] <= arr2[j] {
			arr1[k] = arr2[i]
			i++
		} else {
			arr1[k] = arr2[j]
			j++
		}
	}
}

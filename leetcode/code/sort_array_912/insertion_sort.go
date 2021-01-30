/**
 * @author Emove
 * @date 2021/1/28
 */
package sort_array_912

// 插入排序 O(N^2) O(1) 估计这种排序也会直接超出时间限制
func insertionSort(nums []int) []int {
	length := len(nums)
	for i := 1; i < length; i++ {
		current := nums[i]
		index := i
		for index > 0 && nums[index-1] > current {
			nums[index] = nums[index-1]
			index--
		}
		nums[index] = current
	}
	return nums
}

/**
 * @author Emove
 * @date 2021/1/28
 */
package sort_array_912

// 选择排序 O(N^2) O(1) 好家伙，选择排序直接时间超出限制
// 每次从未排序的数组中选择最大的数
func selectionSort(nums []int) []int {
	length := len(nums)
	for i := range nums {
		min := nums[i]
		minIndex := i
		for j := i + 1; j < length; j++ {
			// 获取最大的数
			if min > nums[j] {
				minIndex = j
				min = nums[j]
			}
		}
		if minIndex != i {
			swap(nums, i, minIndex)
		}
	}
	return nums
}

func swap(nums []int, left, right int) {
	temp := nums[right]
	nums[right] = nums[left]
	nums[left] = temp
}

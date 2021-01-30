/**
 * @author Emove
 * @date 2021/1/28
 */
package sort_array_912

// 排序数组
// 我都不想用go写题解了，leetcode的go测试服务器肯定不咋地
// 同样的代码效率比Java差了几十倍，真离谱
func sortArray(nums []int) []int {
	return selectionSort(nums)
}

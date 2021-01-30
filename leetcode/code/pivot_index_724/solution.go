/**
 * @author Emove
 * @date 2021/1/28
 */
package pivot_index_724

func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	sum := 0
	for i, v := range nums {
		sum += v
		if sum == total {
			return i
		}
		total = total - v
	}
	return -1
}

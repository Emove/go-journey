/**
 * @author Emove
 * @date 2021/1/28
 */
package sort_array_912

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortArray(t *testing.T) {
	nums := []int{-1, 2, -8, -10}
	sort.Ints(nums)
	fmt.Print(nums)
	fmt.Print(MergeSort(nums))
}

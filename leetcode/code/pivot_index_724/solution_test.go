/**
 * @author Emove
 * @date 2021/1/28
 */
package pivot_index_724

import (
	"fmt"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	nums1 := []int{1, 7, 3, 6, 5, 6}
	nums2 := []int{1, 2, 3}
	nums3 := []int{-1, -1, -1, -1, -1, -1}
	nums4 := []int{-1, -1, -1, -1, -1, 0}

	fmt.Println(pivotIndex(nums1))
	fmt.Println(pivotIndex(nums2))
	fmt.Println(pivotIndex(nums3))
	fmt.Println(pivotIndex(nums4))
}

package cost

import (
	"fmt"
	"testing"
)

func TestCost(t *testing.T) {
	fmt.Println(Cost([]int{1, 0, 2}))
	fmt.Println(Cost([]int{3, 1, 0, 2, 4}))
}

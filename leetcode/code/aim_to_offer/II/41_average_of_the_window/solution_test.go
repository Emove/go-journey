package _1_average_of_the_window

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	movingAverage := Constructor(3)
	fmt.Println(movingAverage.Next(1))
	fmt.Println(movingAverage.Next(10))
	fmt.Println(movingAverage.Next(3))
	fmt.Println(movingAverage.Next(5))
}

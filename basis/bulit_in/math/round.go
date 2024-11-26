package math

import (
	"fmt"
	"math"
)

func Round() {
	//fmt.Println(uint32(float64(1.5)))
	fmt.Println(math.Round(0.5))
}

func Div(num int) {
	fmt.Println(float64(num / 2))
}

func Ceil(num int) {
	fmt.Println(math.Ceil(float64(num) / 2))
}

package code

import (
	"fmt"
	"strconv"
	"strings"
)

func convertDoubleInt2float32(x, z int) {
	fmt.Println(float32(x) + (float32(z) / float32(100)))
}

func convertFloat322DoubleInt(num float32) {
	x := int(num) % 32
	atoi, _ := strconv.Atoi(strings.Split(fmt.Sprintf("%v", num), ".")[1])
	z := atoi
	fmt.Printf("x: %d, z: %d \n", x, z)
}

func FormatFloat(num int) {
	fmt.Println(strconv.FormatFloat(float64(num/1000), 'f', 1, 64))
}

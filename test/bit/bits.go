package bit

import (
	"fmt"
	"strconv"
)

func Lowest() {
	n := 8
	fmt.Println(n & -n)
	fmt.Println(strconv.FormatInt(int64(-n), 2))
	fmt.Println(n & (n - 1))
}

func FormatInt() {
	fmt.Println(strconv.FormatInt(int64(7), 2))
	fmt.Println(strconv.FormatInt(int64(-7), 2))
}

func Odd() {
	for i := 0; i < 5; i++ {
		fmt.Println(i & 1)
	}
}

func ToByte(n int) {
	fmt.Println(strconv.FormatInt(int64(n), 2))
}

func XOR() {
	fmt.Println(0 ^ 0)
	fmt.Println(1 ^ 1)
	fmt.Println(1 ^ 2)
	fmt.Println(0 ^ 1)
	fmt.Println(2 ^ 0)
	fmt.Println(2 ^ 1)
	fmt.Println(2 ^ 2)
	fmt.Println(3 ^ 1)
	fmt.Println(3 ^ 2)
}

func AND() {
	fmt.Println(3 & 2)
	fmt.Println(3 & 1)
	fmt.Println(2 & 2)
	fmt.Println(2 & 1)
	fmt.Println(1 & 2)
	fmt.Println(1 & 1)
	fmt.Println(0 & 0)
	fmt.Println(0 & 1)
	fmt.Println(0 & 2)
}

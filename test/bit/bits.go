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

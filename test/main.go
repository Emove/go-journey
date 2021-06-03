package main

import "fmt"

func main() {
	fmt.Println(CatchOrDropCheck(6, 1))
}

func CatchOrDropCheck(x, y int) bool {
	if x > 2 && x < 19 {
		// x坐标在存储区内
		if y >= 1 && y <= 20 {
			// y坐标在轨道范围内
			return true
		}
	}
	return false
}

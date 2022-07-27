package _return

import (
	"fmt"
)

/**
 * @author Emove
 * @date 2021/1/8
 */

func ReturnInForSelect() {
	defer fmt.Println("function finish")
	i := 0
	for {
		select {
		default:
			fmt.Println(i)
			if i == 5 {
				return
			}
			i++
		}
	}
}

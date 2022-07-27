package code

import "fmt"

func BreakSelectTest(ch chan int) {
	defer fmt.Println("exit")
	for {
		select {
		case i := <-ch:
			if i == 2 {
				break
			} else if i == 4 {
				return
			} else {
				fmt.Println(i)
			}
		default:
		}
	}
}

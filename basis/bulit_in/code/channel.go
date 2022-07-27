/**
 * @author Emove
 * @date 2021/1/22
 */
package code

import (
	"fmt"
	"time"
)

func LenOfChannel() {
	ch := make(chan int, 20)
	fmt.Println(len(ch))
	for i := 0; i < 3; i++ {
		ch <- i
	}
	fmt.Println(len(ch))
	_ = <-ch
	fmt.Println(len(ch))
}

func BlockChannel() {
	ch := make(chan int, 1)

	stop := make(chan int, 1)

	go func() {
		for {
			select {
			case i := <-ch:
				if i == 2 {
					stop <- 1
				}
				time.Sleep(time.Second)
			}
		}
	}()

	for i := 0; i < 3; i++ {
		cur := time.Now().Second()
		ch <- i
		fmt.Println("wait-time: ", time.Now().Second()-cur)
	}

	select {
	case <-stop:
		return
	}
}

/**
 * @author Emove
 * @date 2021/1/22
 */
package code

import "fmt"

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

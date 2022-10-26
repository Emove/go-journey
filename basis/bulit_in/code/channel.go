/**
 * @author Emove
 * @date 2021/1/22
 */
package code

import (
	"fmt"
	"sync"
	"sync/atomic"
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

func IsSelectBlocking() {

	timer := time.NewTimer(time.Second)

	for i := 0; i < 5; i++ {

		select {
		case <-timer.C:
			return
		default:
		}

		fmt.Println(i)
		time.Sleep(200 * time.Millisecond)
	}

}

func ClosingByDefault() {

	ch := make(chan int, 10)

	go func() {
		time.Sleep(1 * time.Millisecond)
		for i := 0; i < 10; i++ {
			go func(i int) {
				ch <- i
			}(i)
		}
	}()

	closed := int32(0)
	go func() {
		time.Sleep(1 * time.Second)
		atomic.StoreInt32(&closed, 1)
	}()

	for {
		select {
		case num := <-ch:
			fmt.Println(num)
			time.Sleep(200 * time.Millisecond)
		default:
			if atomic.LoadInt32(&closed) == 1 && len(ch) == 0 {
				close(ch)
				return
			}
		}
	}

}

func ForeachChannel() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	for len(ch) > 0 {
		fmt.Println(<-ch)
	}

}

func LenTheClosedChannel() {
	var ch chan int

	fmt.Printf("len of a nil channel: %d\n", len(ch))

	ch = make(chan int, 10)

	go func() {
		for len(ch) > 0 {
			fmt.Println(<-ch)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	time.Sleep(time.Second)
}

func ReadFromClosedChannel() {
	ch := make(chan int, 10)

	go func() {
		for {
			i, ok := <-ch
			fmt.Println(i, "--", ok)
			time.Sleep(20 * time.Millisecond)
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	time.Sleep(time.Second)
}

func ConvertFromInterface() {
	m := sync.Map{}

	ch := make(chan int)

	m.Store("ch", ch)

	load, _ := m.Load("ch")
	ok := false
	ch, ok = load.(chan int)
	fmt.Printf("ch: %v, ok: %v", ch, ok)
}

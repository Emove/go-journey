package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Add() {
	i := int64(0)
	wg := &sync.WaitGroup{}
	wg.Add(100)
	nums := make([]bool, 100)
	for j := 0; j < 100; j++ {
		go func() {
			nums[atomic.AddInt64(&i, 1)-1] = true
			wg.Done()
		}()
	}

	wg.Wait()
	cnt := 0
	for _, v := range nums {
		if v == false {
			cnt++
		}
	}

	fmt.Println(cnt)
}

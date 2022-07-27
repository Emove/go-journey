package _map

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestIdempotent(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(5)
	key := "test"
	for i := 0; i < 5; i++ {
		go checkIdempotent(key)
		wg.Done()
	}
	wg.Wait()
	time.Sleep(3 * time.Second)
}

func TestForEachMap(t *testing.T) {
	m := make(map[int]string)
	for i := 0; i < 5; i++ {
		m[i] = strconv.Itoa(i)
	}
	for i := range m {
		fmt.Println(i)
	}
}

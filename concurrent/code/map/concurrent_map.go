package _map

import (
	"fmt"
	"sync"
)

var idempotentMap = make(map[string]bool)

var lock sync.Mutex

func checkIdempotent(key string) {
	result := idempotentMap[key]
	if !result {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("come in")
		b := idempotentMap[key]
		if b {
			fmt.Println(true)
			return
		}
		idempotentMap[key] = true
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}
}

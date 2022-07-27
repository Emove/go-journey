package decorator

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewCache(t *testing.T) {
	cache := NewCache()
	for i := 0; i <= 5; i++ {
		str := strconv.Itoa(i)
		cache.put(str, str)
	}
	for i := 0; i <= 5; i++ {
		key := strconv.Itoa(i)
		value := cache.get(key)
		if "" != value {
			fmt.Printf("key: %v -- value: %v\n", key, value)
		}
	}
}

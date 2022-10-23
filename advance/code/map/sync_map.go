package _map

import (
	"fmt"
	"sync"
)

func LoadOrStore() {
	m := sync.Map{}

	m.Store("first_key", "first_value_1")

	store, loaded := m.LoadOrStore("first_key", "first_value_2")
	fmt.Printf("loaded: %v, value: %v\n", loaded, store)

	store, loaded = m.LoadOrStore("second_key", "second_value")
	fmt.Printf("loaded: %v, value: %v\n", loaded, store)

}

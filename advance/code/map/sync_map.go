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

	// invalid
	type s struct {
		v string
	}
	var v *s
	store, loaded = m.LoadOrStore("third_key", v)
	fmt.Printf("loaded: %v, value: %v\n", loaded, store)
	v = &s{v: "third_value"}
	store, loaded = m.Load("third_key")
	fmt.Printf("loaded: %v, value: %v\n", loaded, store)

	// work
	type wrapper struct {
		delegate *s
	}
	w := &wrapper{}
	store, loaded = m.LoadOrStore("forth_key", w)
	fmt.Printf("loaded: %v, value: %v\n", loaded, store)
	w.delegate = v
	store, loaded = m.Load("forth_key")
	fmt.Printf("loaded: %v, value: %v\n", loaded, store)

}

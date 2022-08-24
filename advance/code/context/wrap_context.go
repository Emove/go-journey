package context

import (
	"context"
	"fmt"
)

type key1 struct {
}

type key2 struct {
}

func WrapContext() {
	background := context.Background()

	k1 := context.WithValue(background, key1{}, "key1")

	k2 := context.WithValue(k1, key2{}, "key2")

	fmt.Printf("%v\n", k2.Value(key1{}))
	fmt.Printf("%v\n", k2.Value(key2{}))
}

package color

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	for i := 0; i < 200; i++ {
		c := Random()
		fmt.Println(c)
	}
}

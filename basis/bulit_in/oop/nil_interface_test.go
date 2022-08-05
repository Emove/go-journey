package oop

import (
	"fmt"
	"testing"
)

func TestReturnNilInterface(t *testing.T) {
	if nil != ReturnNilInterface() {
		fmt.Println("hell")
	}
}

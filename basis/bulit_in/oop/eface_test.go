package oop

import (
	"fmt"
	"testing"
)

func TestNewNilNode(t *testing.T) {
	var resource Resource
	resource = NewNilNode()
	if resource == nil {
		fmt.Println("resource is nil")
	} else {
		fmt.Println("resource is not nil")
	}
}

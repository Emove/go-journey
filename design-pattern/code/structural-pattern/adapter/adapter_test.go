package adapter

import (
	"fmt"
	"testing"
)

func TestAdapter_Request(t *testing.T) {
	adaptee := NewAdaptee()
	adapter := NewAdapter(adaptee)
	request := adapter.Request()
	fmt.Println(request)
}

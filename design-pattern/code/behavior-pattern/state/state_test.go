package state

import (
	"testing"
)

func TestNewDayContext(t *testing.T) {
	context := NewDayContext()
	todayAndNext := func() {
		context.Today()
		context.Next()
	}

	for i := 0; i < 12; i++ {
		todayAndNext()
	}
}

package builder

import "testing"

func TestBuilder(t *testing.T) {
	builder := &ComputerBuilder{}
	director := NewDirector(builder)
	director.Constructor()
	if !builder.finished {
		t.Fatal("something wrong")
	}
}

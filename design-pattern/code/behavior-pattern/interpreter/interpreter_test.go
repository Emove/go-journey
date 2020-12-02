package interpreter

import (
	"testing"
)

func TestInterpret(t *testing.T) {
	p := &Parser{}
	p.Parse("1 + 2 + 3 - 4 + 5 - 6")
	result := p.Result().Interpret()
	expect := 1
	if result != expect {
		t.Fatalf("expect %d got %d", expect, result)
	}
}

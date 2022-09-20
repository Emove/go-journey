package reflect

import "testing"

func TestNewByType(t *testing.T) {
	NewByType()
}

func TestByJSON(t *testing.T) {
	ByJSON()
}

func BenchmarkNewByType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		NewByType()
	}
}

func BenchmarkByJSON(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ByJSON()
	}
}

func TestPrint(t *testing.T) {
	Print()
}

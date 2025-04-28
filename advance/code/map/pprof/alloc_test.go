package main

import (
	"testing"
)

func Benchmark_Write(b *testing.B) {
	Write(200000)
}

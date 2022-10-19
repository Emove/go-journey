package _map

import (
	"math/rand"
	"testing"
	"time"
)

const IN = 100000000

// 195 ns/op
func BenchmarkWriteRWM(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for pb.Next() {
			k := r.Intn(IN)
			WriteRWM(k)
		}
	})
}

// 48 ns/op
func BenchmarkReadRWM(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for pb.Next() {
			k := r.Intn(IN)
			ReadRWM(k)
		}
	})
}

// 58 ns/op
func BenchmarkDeleteRWM(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for pb.Next() {
			k := r.Intn(IN)
			DeleteRWM(k)
		}
	})
}

// ---------------------------------------  total in average: 195 + 48 + 58 = 301ns -------------------------------------------- //

// 290 ns/op
func BenchmarkWriteSyncM(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for pb.Next() {
			k := r.Intn(IN)
			WriteSyncM(k)
		}
	})
}

// 2 ns/op
func BenchmarkReadSyncM(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for pb.Next() {
			k := r.Intn(IN)
			ReadSyncM(k)
		}
	})
}

// 2 ns/op
func BenchmarkDeleteSyncM(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for pb.Next() {
			k := r.Intn(IN)
			DeleteSyncM(k)
		}
	})
}

// ---------------------------------------  total in average: 290 + 2 + 2 = 294ns -------------------------------------------- //

// 208 ns/op
func BenchmarkWriteReadDeleteMM(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for pb.Next() {
			k := r.Intn(IN)
			WriteMM(k)
			ReadMM(k)
			DeleteMM(k)
		}
	})
}

// 208 ns/op
func BenchmarkWriteReadDeleteRWM(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for pb.Next() {
			k := r.Intn(IN)
			WriteRWM(k)
			ReadRWM(k)
			DeleteRWM(k)
		}
	})
}

// 625 ns/op
func BenchmarkWriteReadDeleteSyncM(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for pb.Next() {
			k := r.Intn(IN)
			WriteSyncM(k)
			ReadSyncM(k)
			DeleteSyncM(k)
		}
	})
}

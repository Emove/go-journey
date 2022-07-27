package singleton

import (
	"sync"
	"testing"
)

const count = 100

func TestGetInstanceByDoubleCheck(t *testing.T) {
	ins1 := GetInstanceByDoubleCheck()
	ins2 := GetInstanceByDoubleCheck()
	if ins1 != ins2 {
		t.Fatal("ins1 is not equals to ins2")
	}

}

func TestParallelGetInstanceByDoubleCheck(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(count)
	instances := [count]*Singleton{}
	for i := 0; i < count; i++ {
		go func(i int) {
			instances[i] = GetInstanceByDoubleCheck()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < count; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance not equals")
		}
	}
}

func TestGetInstance(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("ins1 is not equals to ins2")
	}
}

func TestParallelGetInstance(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(count)
	instances := [count]*Singleton{}
	for i := 0; i < count; i++ {
		go func(i int) {
			instances[i] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < count; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance not equals")
		}
	}
}

// 两种方法性能相差并不大，而且耗时相当小
// 相差大概0.15ns
func BenchmarkGetInstanceByDoubleCheck(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = GetInstanceByDoubleCheck()
		}
	})
}

func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = GetInstance()
		}
	})
}

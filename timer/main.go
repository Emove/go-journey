package main

import (
	"fmt"
	"github.com/antlabs/timer"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

type Case struct {
	StartAt  int64
	EndAt    int64
	Duration time.Duration
}

var Queue chan *Case
var total int64
var totalDiff int64
var wg *sync.WaitGroup

func main() {

	nums, d := int64(1e6*10), 10*time.Second
	total = nums
	prepare(nums)
	wg = &sync.WaitGroup{}
	start := time.Now()
	do("std", nums, d, func(c *Case) {
		wg.Add(1)
		time.AfterFunc(c.Duration, func() {
			c.EndAt = time.Now().UnixMilli()
			Queue <- c
		})
	})

	wg.Wait()
	fmt.Printf("total nums: %d, base duration: %s, average diff: %d, elapsed: %dms\n", total, "10ms", totalDiff/total, time.Since(start).Milliseconds())
	totalDiff = 0
	antTimer := timer.NewTimer()
	go antTimer.Run()
	time.Sleep(1 * time.Second)
	start := time.Now()
	do("antlabs", nums, d, func(c *Case) {
		wg.Add(1)
		antTimer.AfterFunc(c.Duration, func() {
			c.EndAt = time.Now().UnixMilli()
			Queue <- c
		})
	})

	wg.Wait()
	fmt.Printf("total nums: %d, base duration: %s, average diff: %d, elapsed: %dms\n", total, "10ms", totalDiff/total, time.Since(start).Milliseconds())
}

func do(name string, nums int64, d time.Duration, fn func(c *Case)) {

	//fds := startPprof(name, nums)
	//defer endPprof(fds)

	times := d / (10 * time.Millisecond)
	p, i := nums/int64(times), nums%int64(times)
	fmt.Printf("times: %d, add %d func per time, i: %d\n", times, p, i)
	ticker := time.NewTicker(10 * time.Millisecond)
	cnt := 0
	var second *time.Timer
	second = time.AfterFunc(time.Second, func() {
		cnt = 0
		if cnt > 0 {
			fmt.Printf("per second count: %d\n", cnt)
			second.Stop()
		}
	})

	for range ticker.C {
		now := time.Now().UnixMilli()
		for k := int64(0); k < p; k++ {
			if i > 0 {
				d = d - time.Duration(i)
			}
			fn(&Case{StartAt: now, Duration: d})
			cnt++
		}
		if i >= 0 {
			fn(&Case{StartAt: now, Duration: d})
			cnt++
			i--
		}
		times--
		if times == 0 {
			break
		}
	}
}

func startPprof(name string, nums int64) []*os.File {
	cpu, mem := getFd(name, nums)
	err := pprof.StartCPUProfile(cpu)
	if err != nil {
		panic(err)
	}
	err = pprof.WriteHeapProfile(mem)
	if err != nil {
		panic(err)
	}
	return []*os.File{cpu, mem}
}

func endPprof(fds []*os.File) {
	pprof.StopCPUProfile()
	for _, fd := range fds {
		_ = fd.Close()
	}
}

func getFd(name string, nums int64) (*os.File, *os.File) {
	cpuPprofFileName := fmt.Sprintf("./%s/%s_%dM_cpu.pprof", name, name, nums/1e6)
	fd1, err := os.Open(cpuPprofFileName)
	if err != nil {
		fd1, err = os.Create(cpuPprofFileName) // 在当前路径下创建一个cpu.pprof文件
		if err != nil {
			panic(fmt.Sprintf("create cpu pprof failed, err: %v", err))
			return nil, nil
		}
	}
	memPprofFileName := fmt.Sprintf("./%s/%s_%dM_mem.pprof", name, name, nums/1e6)
	fd2, err := os.Open(memPprofFileName)
	if err != nil {
		fd2, err = os.Create(memPprofFileName) // 在当前路径下创建一个mem.pprof文件
		if err != nil {
			panic(fmt.Sprintf("create mem pprof failed, err: %v", err))
			return nil, nil
		}
	}
	return fd1, fd2
}

func prepare(nums int64) {
	Queue = make(chan *Case, nums)
	go func() {
		for {
			select {
			case c := <-Queue:
				totalDiff += abs(c.StartAt+int64(c.Duration/time.Millisecond), c.EndAt)
				wg.Done()
			}
		}
	}()
}

func abs(a, b int64) int64 {
	if a > b {
		return a - b
	}
	return b - a
}

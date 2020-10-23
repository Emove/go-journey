package main

import (
	"concurrent/code"
)

func main() {
	// testGoroutine()
	// testRuntime()
	testChannel()
}

func testGoroutine() {
	code.TestMultiGoroutine()
}

func testRuntime() {
	// code.TestReSchedule()
	// code.TestGoExit()
	code.TestSetMaxProcs(2)
}

func testChannel() {
	// code.NoBuffChanTest()
	// code.GetDataFromChannel()
	code.TestSingleTrackChannel()
}

// 主协程退出之后，其他任务就不执行了
// func main() {
// 	go func() {
// 		i := 0
// 		for {
// 			i++
// 			fmt.Printf("new goroutine：i = %d\n", i)
// 			time.Sleep(time.Second)
// 		}
// 	}()
// 	i := 0
// 	for {
// 		i++
// 		fmt.Printf("main goroutine: i = %d \n", i)
// 		time.Sleep(time.Second)
// 		if i == 2 {
// 			break
// 		}
// 	}
// }

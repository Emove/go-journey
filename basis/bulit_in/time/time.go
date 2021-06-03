/**
 * @author Emove
 * @date 2021/2/4
 */
package time

import (
	"fmt"
	"time"
)

func GetYesterdayTime() {
	fmt.Println(time.Now().Format("2006-01-02 03:04:05.000"))
	time := time.Now().Add(-24 * time.Hour)
	fmt.Println(time.Format("2006-01-02 03:04:05.000"))
}

func Now() {
	fmt.Println(time.Now().UnixNano())
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now().UnixNano())
}

func DatetimeSerialization() string {
	return time.Now().Format("20060102030405")
}

func TimeSerialization() string {
	return time.Now().Format("030405")
}

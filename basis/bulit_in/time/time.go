/**
 * @author Emove
 * @date 2021/2/4
 */
package time

import (
	"fmt"
	tt "time"
)

func GetYesterdayTime() {
	fmt.Println(tt.Now().Format("2006-01-02 03:04:05.000"))
	t := tt.Now().Add(-24 * tt.Hour)
	fmt.Println(t.Format("2006-01-02 03:04:05.000"))
}

func Now() {
	//fmt.Println(time.Now().UnixNano())
	//time.Sleep(1 * time.Second)
	//fmt.Println(time.Now().UnixNano())
	//fmt.Println(time.Now().Unix())
	//time.Sleep(1 * time.Second)
	//fmt.Println(time.Now().Unix())
	priority := 1
	fmt.Println(int64(priority*10000000000) + tt.Now().Unix())
}

func DatetimeSerialization() string {
	return tt.Now().Format("20060102030405")
}

func TimeSerialization() string {
	return tt.Now().Format("030405")
}

func FromMillis(millis int64) tt.Time {
	return tt.UnixMilli(millis)
}

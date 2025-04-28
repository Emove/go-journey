package time

import (
	"fmt"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	Now()
}

func TestDatetimeSerialization(t *testing.T) {
	fmt.Println(DatetimeSerialization())
}

func TestTimeSerialization(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(TimeSerialization())
	}
}

func TestFromMillis(t *testing.T) {
	n := time.Now().Add(time.Hour)
	um1 := n.UnixMilli()
	ti := FromMillis(n.UnixMilli())
	fmt.Println(um1)
	fmt.Println(ti.UnixMilli())
}

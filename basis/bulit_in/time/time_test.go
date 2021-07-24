package time

import (
	"fmt"
	"testing"
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

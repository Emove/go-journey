package time

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"
)

func TestNow(t *testing.T) {
	//Now()
	for i := 0; i < 10; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(1000))
		fmt.Println(n.String())
	}
}

func TestDatetimeSerialization(t *testing.T) {
	fmt.Println(DatetimeSerialization())
}

func TestTimeSerialization(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(TimeSerialization())
	}
}

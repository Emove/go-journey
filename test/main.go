package main

import (
	"encoding/binary"
	"fmt"
)

func main() {

	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, 76698383)
	fmt.Println(buf)
}

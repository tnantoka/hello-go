package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	buf := make([]byte, 2)
	value := 54321

	binary.BigEndian.PutUint16(buf[0:], uint16(value))
	fmt.Println(buf)          // [212 49]
	fmt.Println(212*256 + 49) // 54321
}

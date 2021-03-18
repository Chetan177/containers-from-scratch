package main

import (
	"fmt"
	"time"
)

const size = 10000000

func main() {
	fmt.Printf("Allocating huge buffer\n")

	m := [size]uint64{55}
	for i := 0; i < size; i++ {
		m[i] = uint64(i)
	}

	time.Sleep(5 * time.Second)
}

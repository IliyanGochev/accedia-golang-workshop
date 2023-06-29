package main

import (
	"fmt"
)

const (
	first  = iota + 1 // 1
	second            // 2
	_                 // skip 3
	fourth            // 4
)

const (
	_  = 1 << (iota * 10) // skip the first value
	KB                    // decimal:       1024 -> binary 00000000000000000000010000000000
	MB                    // decimal:    1048576 -> binary 00000000000100000000000000000000
	GB                    // decimal: 1073741824 -> binary 01000000000000000000000000000000
)

func main() {
	fmt.Printf("first: %d\n", first)
	fmt.Printf("second: %d\n", second)
	fmt.Printf("fourth: %d\n", fourth)

	fmt.Printf("KB: %d\n", KB)
	fmt.Printf("MB: %d\n", MB)
	fmt.Printf("GB: %d\n", GB)
}

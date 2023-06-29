package main

import (
	"fmt"
)

func main() {
	i := 1
	f := float64(i)
	u := uint(i)

	fmt.Printf("i:	Type: %T	Value: %v\n", i, i)
	fmt.Printf("f:	Type: %T	Value: %v\n", f, f)
	fmt.Printf("u:	Type: %T	Value: %v\n", u, u)
}

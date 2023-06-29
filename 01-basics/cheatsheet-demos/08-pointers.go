package main

import (
	"fmt"
)

func main() {
	var add *int // an int pointer
	num := 42

	add = &num  // stores the memory address of num
	val := *add // dereferencing

	fmt.Println(num) // 42
	fmt.Println(add) // e.g., 0xc0000ae008
	fmt.Println(val) // 42
}

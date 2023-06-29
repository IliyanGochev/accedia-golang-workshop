package main

import (
	"fmt"
)

func main() {
	var add *int
	num := 42

	add = &num  // stores the memory address of num
	val := *add // dereferencing

	fmt.Println("num:", num) // num: 42
	fmt.Println("add:", add) // add: e.g., 0xc0000ae008
	fmt.Println("val:", val) // val: 42
}

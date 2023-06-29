package main

import "fmt"

func main() {
	var a string
	a = "initial"
	fmt.Printf("a:	Type: %T	Value: %v\n", a, a)

	var b, c int = 1, 2
	fmt.Printf("b:	Type: %T	Value: %v\n", b, b)
	fmt.Printf("c:	Type: %T	Value: %v\n", c, c)

	var d int
	fmt.Printf("d:	Type: %T	Value: %v\n", d, d)

	e := true
	fmt.Printf("e:	Type: %T	Value: %v\n", e, e)
}

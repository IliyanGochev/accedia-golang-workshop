package main

import (
	"fmt"
)

const Pi = 3.14

func main() {
	var radius float32
	radius = 2

	circumference := 2 * Pi * radius

	fmt.Printf("The circumference of the circle is %f", circumference)
}

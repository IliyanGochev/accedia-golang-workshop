package main

import (
	"fmt"
)

var pl = fmt.Println

// function that takes two integers and returns their sum
func add(a int, b int) int {
	return a + b
}

// function that takes a slice of integers and returns the sum of all elements
func sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

// Variadic functions is a function that can take a variable number of arguments.
// In Go, you can create a variadic function by using the `...` syntax before the type of the final parameter in the function signature
func vSum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

// function that returns multiple values
func divide(div int, divisor int) (int, int) {
	factor := div / divisor
	remainder := div % divisor

	return factor, remainder
}

// function that takes a function as an argument
func apply(f func(int) int, a int) int {
	return f(a)
}

// One thing that is specific to Go when it comes to functions
// is that it allows you to return multiple values from a single function
// which can be useful  in certain scenarios. You can also use named return values,
// which can make the code more readable
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func square(n int) int {
	return n * n
}

func main() {
	// Go also allows you to pass functions as arguments to other functions
	// This can be useful when you need to pass a behavior to another part of your code
	// This enables functional programming in Go
	result := apply(square, 10)
	pl(result)

	//Variadic function example:
	vResult := vSum(1, 2, 3, 4, 5, 6)
	pl(vResult)

	// We can also pass a slice of values to a variadic function by using the `...` syntax.
	numbers := []int{1, 2, 3, 4, 5, 6}
	vResultSlice := vSum(numbers...)
	pl(vResultSlice)
}

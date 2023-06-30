package main

import (
	"fmt"
	"strconv"
)

func factorial(f int) int {
	var factorial = 1
	for i := 1; i <= f; i++ {
		factorial = factorial * i
	}
	return factorial
}

func numberOfDigits(number int) int {
	return len(strconv.Itoa(number))
}

func main() {
	//input := [5]int{1, 2, 3, 4, 5}
	input := [5]int{6, 6, 1, 4, 7}
	for _, number := range input {
		fmt.Print(numberOfDigits(factorial(number)), " ")
	}
}

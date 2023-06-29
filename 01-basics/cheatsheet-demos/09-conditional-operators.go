package main

import (
	"errors"
	"fmt"
	"time"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Divide by zero")
	}
	return a / b, nil
}

func main() {

	//Example 1: simple if-else structure

	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	fmt.Println("--------------------------------------------------------------------------------")

	//Example 2: initialization of a variables (result and err) in the if statement

	if result, err := Divide(10, 0); err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Response", result)
	}

	//uncommenting the following line will result in an error
	//fmt.Println(number, " can not be reached out of the if-else scope")

	fmt.Println("--------------------------------------------------------------------------------")

	//Example 3: Switch operator: multiple cases, separated by comma, use of the fallthrough keyword

	letter := 'y'

	switch letter {
	case 'a', 'e', 'o', 'u', 'i':
		fmt.Println("Vowel")
	case 'y':
		fmt.Print("Vowel and ")
		fallthrough
	default:
		fmt.Println("Consonant")
	}

	fmt.Println("--------------------------------------------------------------------------------")

	//Example 4: switch operator without a condition and using complex values for cases

	switch hour := time.Now().Hour(); { // missing expression means "true"
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

}

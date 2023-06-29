package main

import (
	"fmt"
)

func main() {

	//Example 1: simple loop, equivalent to while; the initialization and increment are out of the for statement
	//the scope of the declared variable (i) is in the whole main function

	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println(i, " is accessible outside the for block")
	fmt.Println("--------------------------------------------------------------------------------")

	//Example 2: A typical for loop, consisting of initialization, condition and iteration
	//the scope of the declared variable (j) is in the for block, but not outside of it

	for j := 5; j < 10; j++ {
		fmt.Println(j)
	}

	//uncommenting the following line will result in an error
	//fmt.Println(j, " is not accessible outside the for block")
	fmt.Println("--------------------------------------------------------------------------------")

	//Example 3: Loop without initialization, condition or iteration => infinite loop
	// demonstration of beak and continue keywords
	n := 0
	for {
		n++
		//comment out the following if to achieve infinite loop
		if n == 10 {
			break
		}
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	fmt.Println("--------------------------------------------------------------------------------")

	//Example 4: Iterating over slices with the range operator

	slice := []int{1, 2, 3}

	// iteration can be achieved with for loop
	for index := 0; index < len(slice); index++ {
		fmt.Println(index, slice[index])
	}

	// or with range : both are equivalent
	for index, sliceValue := range slice {
		fmt.Println(index, sliceValue)
	}

	fmt.Println("--------------------------------------------------------------------------------")

	//Example 5: Iterating over maps with the range operator

	favColors := map[string]string{"John": "blue", "Markus": "red", "Anna": "green"}

	for key, value := range favColors {
		fmt.Println("Name:", key, "Color:", value)
	}

	fmt.Println("--------------------------------------------------------------------------------")

	//Example 6: Iterating over strings: the unicode value of any character is accessible via the index

	title := "GolangWorkshop"
	for index, unicodeValue := range title {
		fmt.Println(index, unicodeValue)
	}

}

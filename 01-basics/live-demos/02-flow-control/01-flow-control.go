package main

import "fmt"

func main() {
	// slices and arrays

	ids := [3]int{13, 27, 36}

	for i := 0; i < 3; i++ {
		fmt.Println(i, ids[i])
	}

	for index := range ids {
		fmt.Println(index, ids[index])
	}

	for index, value := range ids {
		fmt.Println(index, value)
	}

	for _, value := range ids {
		fmt.Println(value)
	}

	// maps

	products := map[int]string{13: "Toothbrush", 27: "Tomato", 36: "Coffee"}

	for key, value := range products {
		fmt.Println(key, ":", value)
	}

	for key := range products {
		fmt.Println(key)
	}

	for _, value := range products {
		fmt.Println(value)
	}

	// strings
	// https://www.ssec.wisc.edu/~tomw/java/unicode.html
	password := "abc"

	for key, value := range password {
		fmt.Println(key, value)
	}
}

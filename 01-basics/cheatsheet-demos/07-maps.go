package main

import (
	"fmt"
)

func main() {
	fileExtensions := map[string]string{
		"Golang":     ".go",
		"Java":       ".java",
		"JavaScript": ".js",
		"Python":     "",
	}

	fileExtensions["Python"] = ".py" // update a value

	fmt.Printf("Python:		%v, 	Length: %d\n", fileExtensions["Python"], len(fileExtensions))

	fileExtensions["PHP"] = ".php" // add a new pair

	fmt.Printf("PHP:		%v, 	Length: %d\n", fileExtensions["PHP"], len(fileExtensions))

	delete(fileExtensions, "PHP") // rempove a pair

	fmt.Printf("PHP:		%v, 	Length: %d\n", fileExtensions["PHP"], len(fileExtensions))

	val, ok := fileExtensions["Golang"] // check for a value

	fmt.Printf("Golang:		%v, 	Exists: %v\n", val, ok)
}

package demo

import (
	"fmt"
)

// the purpose of this package is to be used as a demonstration for importing packages within the project
//this file does not contain the main function, so it's not executable

// this function is capitalized => it will be accessible everywhere, where the package is imported
func SayHi() {
	fmt.Println("Hello!")
}

// the lower case of this function encloses the scope of visibility on a package level
func sayBye() {
	fmt.Println("Goodbye!")
}

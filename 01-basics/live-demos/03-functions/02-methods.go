package main

import (
	"fmt"
)

var pl = fmt.Println
var sf = fmt.Sprintf

// In Go, methods are functions that are associated with a specific type
type Rectangle struct {
	width  int
	height int
}

// Method that calculates the area
func (r Rectangle) Area() int {
	return r.width * r.height
}

// Method that calculates the perimeter of a rectangle
func (r Rectangle) Perimeter() int {
	return 2 * (r.width + r.height)
}

// One thing that is specific to Go when it comes to methods is that it allows you to define methods on any type.
//This includes both custom types and built-in types, such as `int`, `string` and `float64`

type User struct {
	FullName string
	Email    string
	Age      int
}

func (p User) String() string {
	return sf("Name: %s, Email: %s, Age: %d", p.FullName, p.Email, p.Age)
}

// In this case, we've defined a custom type `User` and a method `String` associated with it.
// The `String` method returns a string representation of a `User` instance, which can be useful for printing and loggin..
// Go automatically calls this method when a `User` instance is passed to a function that expects a string

func main() {

	user := User{
		FullName: "Nikolay Pleshkov",
		Email:    "nikolay.pleshkov@emai.com",
		Age:      21,
	}

	pl(user.String())
}

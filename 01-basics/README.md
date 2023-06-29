# Go Workshop

# Table of contents

- [**Basics**](#basics)
  - [Overview](#overview)
    - [Characteristics](#characteristics)
    - [Usage](#usage)
  - [Installing Go](#installing-go)
  - [The `go` command](#the-go-command)
  - [Go projects](#go-projects)
    - [Creating a proejct](#creating-a-project)
    - [Naming requirements](#naming-requirements)
    - [Packages](#packages)
    - [Modules](#modules)
    - [Access Modifiers](#access-modifiers)
    - [go.mod](#go.mod)
    - [go.sum](#go.sum)
    - [Dependencies](#dependencies)
    - [Replacing dependencies](#replacing-dependencies)
    - [Synchronizing dependencies](#synchronizing-dependencies)
  - [Variables](#variables)
    - [Long syntax](#long-syntax)
    - [Short syntax](#short-syntax)
  - [Constants](#constants)
  - [Iota and constant expressions](#iota-and-constant-expressions)
  - [Data types](#data-types)
    - [Bool](#bool)
    - [String](#string)
    - [Integers](#int-int8-int16-int32-int64)
    - [Unsigned integers](#uint-uint8-uint16-uint32-uint64-uintptr)
    - [Byte](#byte)
    - [Rune](#rune)
    - [Floating point](#float32-float64)
    - [Complex numbers](#complex64-complex128)
    - [Aliases](#aliases)
    - [Pointers](#pointers)
    - [Collections](#collections)
      - [Structs](#structs)
      - [Arrays](#arrays)
      - [Slices](#slices)
      - [Maps](#maps)
  - [Type conversions](#type-conversions)
  - [Flow control](#flow-control)
    - [If](#if)
    - [Switch](#switch)
    - [For](#for)
      - [Single condition](#single-condition)
      - [Initial/condition/after for loop](#initialconditionafter-for-loop)
      - [Infinite loop](#infinite-loop)
    - [Range](#range)
      - [On arrays and slices](#on-arrays-and-slices)
      - [On maps](#on-maps)
      - [On strings](#on-strings)
  - [Functions](#functions)
    - [Closures](#closures)
    - [Variadic parameters](#variadic-parameters)
    - [Returning multiple values](#returning-multiple-values)
    - [Passing functions as arguments](#passing-functions-as-arguments)
    - [Deferred functions](#deferred-functions)
- [**Extra readings**](#extra-readings)
  - [Online Guides and Articles](#online-guides-and-articles)
  - [Books](#books)
  - [Videos](#videos)

# Basics

## Overview

Go/Golang was designed by Robert Griesemer, Rob Pike, and Ken Thompson at Google in 2007. They strived to create something that could handle high-performance networking and multithreading with the runtime speed of C and the readability and ease of use of Python.

### Characteristics

- statically typed
- compiled
- concurrent, although multithreaded when possible
- has a garbage collector
- has an extensive standard library
- "post-OOP":
  - structs instead of classes => no static or abstract
  - no type hierarchy
  - implicit interface implementation

### Usage

- Web services
- Web apps
- Task automation
- GUIs
- Machine learning

## Installing go

Follow the instructions [here](https://go.dev/doc/install), when done, you should be able to run `go version` in your terminal.

## The go command

The go command-line interface (CLI) is a cross-platform toolchain for developing, building, running, and publishing go applications. The motivation behind it, its history, and the conventions it abbides to can be found [here](https://go.dev/doc/articles/go_command). Here are some common commands and what they are used for.

```sh
go help <topic>               # helper info for the cli, eg. `go help mod init`

go mod init                   # creates a new go.mod file which defines a new module

go run <path to package>      # compiles and runs the main package

go build <path to package>    # compiles and creates an executable of the main package

go doc <topic>                # prints docs for a topic, eg. `go doc fmt.Println`
                              # docs are also available on https://pkg.go.dev/std

go get <name>                 # updates the go.mod file with the new dependency
                              # and downloads its source code

go mod tidy                   # cleans up unused or adds missing dependencies
```

## Go projects

### Creating a project

1. Create a new directory, for example `my-module`
2. Run `go mod init golang-workshop/my-module`. This will create a go.mod file (similar to a package.json in JS or the project file in .NET).
3. Create a main.go file and paste the following code inside it:

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

4. Run `go run ./main.go` and the program will run.

### Naming requirements

Module names follow the format `<prefix>/<name>` where prefix is the place where the project may be found, be it a company internal namespace or a repository address. For instance, it's common to see module names such as `github.com/accedia/golang-project` or `dev.azure.com/accedia/golang-project`.

### Packages

A package is a directory containing .go files, and it is the basic building block of a Go program. They are designed to be small and numerous. They are given lower case, single-word names. Packages can import other packages.

There are 2 types of packages:

- executable - containing the `main` function.
- non executable/non-main/utility packages - all other

#### Importing

Importing a package can be done using the following syntax:

```go
import "fmt"
import "math/rand"
```

Brackets can be used to import multiple packages

```go
import(
    "fmt"
    "math/rand"
)
```

There is no requirement for package names to be unique, hence, duplicates can happen. To resolve this, an alias can be used:

```go
import myAlias "fmt"

func main(){
    myAlias.Println("An alias was given to the package")
}
```

Similar to unused variables, imported packages that are unused are considered a linting error.

### Modules

In Golang, .go files are bundled into packages and those packages can form a module. For instance, when you created a project in the [creating a proejct](#creating-a-project) section, a `go.mod` file was created. The existance of this file differentiates between a package and a module.

Here are a few key characteristics to keep in mind:

- A `module` is a group of packages (or a single package) that are versioned
- Each module is identified by a unique string which is called the `module path`
- `go.mod` is at the root of dependency management in Go. It is used to specify module information, language version & dependencies.
- `go.sum` contains cryptographic hashes of the specific versions installed for each direct and indirect dependency.

### Access modifiers

There are only two forms of access in Go:

- local (private, unexported) - written in **lowercase**
- global (public, exported) - written in **Capitalcase**

```go
// global
func ExportedFunction() {
	fmt.Println("Im being exported, I can be accessed outside of the package!")
}

// local
func localFunction() {
	fmt.Println("Im a unexported, I am available only in the declared package!")
}
```

- All identifiers defined within a package are visible throughout that package

- Only **capitalized** names are accessible outside of the package where they are declared

```go
package demo

import (
	"fmt"
)

func SayHi() {
	fmt.Println("Hello!")
}

func sayBye() {
	fmt.Println("Goodbye!")
}

package main

import (
	"github.com/GolangWorkshop/golang-workshop/demo"
)

func main() {
    demo.SayHi()
    //demo.sayBye()   -> error: not accessible
}
```

Keep in mind that unexported identifiers are not a security measure and do not hide or protect any information

### Go.mod

- `go.mod` has the following structure:

```sh
module github.com/GolangWorkshop/golang-workshop

go 1.19

require (
	cloud.google.com/go/pubsub v1.17.1
	cloud.google.com/go/storage v1.10.0
  )

require (
	cloud.google.com/go v0.97.0 // indirect
)
```

The first line gives the `module path`. Depending on your use case this can be:

- a local path
- URI to a repository hosted on a version control system

The second line gives the specific version of Go

There are two `require` sections:

- the first contains dependencies that are used **directly** by the module
- the second one contains **indirect** dependencies marked by `// indirect`. These can be:
  - any unused imported dependencies
  - any indirect dependency which is not listed in the go.mod file of your direct dependency
  - any direct dependency that doesn't have a `go.mod` file

### Go.sum

- `go.sum` is an automatically generated file

- `go.sum` has the following structure:

```sh
github.com/google/go-cmp v0.5.9 h1:O2Tfq5qg4qc4AmwVlvv0oLiVAGB7enBSJ2x2DqQFi38=
github.com/google/go-cmp v0.5.9/go.mod h1:17dUlkBOakJ0+DkrSSNjCkIjxS6bF9zb3elmeNGIjoY=
```

- One dependency translates into two lines with the following structure

```
module path version h1: checksum
module path/go.mod version h1:checksum
```

- Line one records the checksum of the dependency
- Line two records the checksum of the go.mod of the dependency
- `checksum` guarantees the integrity of the module

### Dependencies

There are 2 ways for adding a new dependency:

- Using the `go get` command

```go
go get github.com/google/go-cmp/cmp
```

- Go can automatically fetch packages from remote repositories by adding them to the import statement and running `go get .` after it

```go
import (
  "fmt"
  "github.com/google/go-cmp/cmp"
)

func main() {
  fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
```

```go
PS C:\repos\golang-workshop> go get .
go: added github.com/google/go-cmp v0.5.9
```

### Replacing dependencies

It is possible to replace the code of a module with the code of another module by using this syntax in a go.mod file:

```sh
replace (
   gitlab.com/gitusername/corge => ./corgeforked
)
```

- The replacement module can be stored either locally or remotely
- The replacement module should have **the same first line of the go.mod file**
- Version specification for replacement for **remote modules is required** but for local is not
- A specific version of a module can be replaced OR all versions can be replaced

### Synchronizing dependencies

To synchronize dependencies remove unused ones and install missing ones, use `go mod tidy`

## Variables

### Long syntax

```go
var a int
var b, c bool
var d string = "hello world"
```

In go, all variables have zero-values. For instance, if a variable is only declared as a `bool` its initial value will be `false`

### Short syntax

A shorthand syntax is available inside any code block. The type of the variable is inferred (type infrerance):

```go
// b := "hello world" // short syntax outside block => won't compile

func main() {
  a := "hello world"
  b := 3
  c := 3.1
}
```

## Constants

- Must be defined at compile time
- The type can be inferred

```go
const Version string = "0.0.1"
const pi = 3.14

// or
const (
    Version string = "0.0.1"
    pi = 3.14
)
```

## Iota and constant expressions

- A successive integer constant `0, 1, 2 ...`
- Resets to `0` whenever `const` appears in the source code

```go
const (
	a = iota      // 0
	b = iota      // 1
)
const (
	c = iota      // 0
	d = iota + 10 // 11
)
const (
	e = iota      // 0
	f             // 1
	g             // 2
)
```

## Data types

Golang supports the following primitive data types:

### bool

```go
t := true
f := false
```

### string

```go
str := "world"
str2 := `Multiline
string`

// Golang supports complex string formatting and interpolation

fmt.Printf("Hello, %s %d!\n", str, 5)
```

For iterations over strings - check [Range](#range)

### int int8 int16 int32 int64

### uint uint8 uint16 uint32 uint64 uintptr

```go
i := 3
```

### byte

Alias for uint8

```go
b := byte('a')
```

### rune

Alias for int32; represents a Unicode code point

To see what `rune` does we need to handle unicode characters that are represented by more than 1 bytes. Notice the `Ö` in the next example. When casted to a byte array, we can see the two bytes needed to represent it `195 & 150`. In the case with the rune array, we see that it's represented by a single decimal value:

```go
s := "GÖ"
fmt.Println([]byte(s))  // [71 195 150]
fmt.Println([]rune(s))  // [71 214]; 214 is \u00d6 in decimal
```

### float32 float64

```go
f := 3.1
```

### complex64 complex128

```go
num := complex(3, 4)
r, im := real(num), imag(num)
```

### Aliases

Golang allows the creation of custom aliases that point to existing types. The usefullness of this will become more clear once we tackle generics.

```go
type UserID int // makes UserID and int interchangeable

// ...

var a UserID
a = 5
```

### Pointers

Golang supports working explicitly with pointers. A pointer has the following characteristics:

- Holds the memory address of a value
- The zero value of a pointer is `nil`
- Operators:
  - `&`
    - Provides the memory address of a variable
  - `*`
    - Dereferences a pointer to the underlying value
    - Denotes a type as a pointer

```go
var address *int     // denote an int pointer
num := 42

address = &num       // get the memory address
value := *add        // dereference

fmt.Println(num)     // 42
fmt.Println(address) // 0xc000...
fmt.Println(value)   // 42
```

### Collections

#### Structs

Instead of classes, Go has structs. Structs are a typed collection of fields used to define more complex data types. Fields are accessed by the `.` (dot) operator.

```go
type DatabaseConfig struct {
  ConnectionString string
}

type WebServerConfig struct {
  Host string
  Port int
}

type Config struct {
  Database DatabaseConfig
  WebServer WebServerConfig
}

dc := DatabaseConfig{ConnectionString: "..."}
wsc := WebServerConfig{
  Host:   "localhost",
  Port:   8080,        // ending comma is mandatory
}

p := &wsc

fmt.Println(p.Host)    // implicit dereference is allowed on structs
```

Methods are defined separately

```go
type Human struct {
  Name string
  age int
}

func (h Human) Speak(words string) {
  fmt.Println(words)
}
```

Constructor functions

```go
func NewHuman(Name string, age int) (Human, error) {
  h := Human{Name, age}
  return h, nil
}
```

Capitalization is important in Golang (check [Access Modifiers](#access-modifiers) for more):

- The `Human` struct and its `Name` field are accessible from other packages
- The `age` field is accessible only from the current package

#### Arrays

Arrays are a fix-sized collection of elements whose type is the same. Arrays are a value type. For iterations over elements - check [Range](#range).

Usage:

```go
var arr [3]int

// the zero value of an array is a ready-to-use array whose elements are themselves zeroed:
fmt.Println(arr[0]) // 0

arr[0] = 1
arr[1] = 2
arr[2] = 3

arr2 := [3]int{1, 2, 3}

fmt.Println(len(arr))       // use len() to get the length of the array
fmt.Println(arr == arr2)    // true!

```

#### Slices

Slices can be viewed as a reference to a section of an underlying array. The 2 main differences between an array and a slice are that slices can vary in length at runtime and are passed as a reference.

There are a few common ways to create a slice:

- declared only

  ```go
  var s []int
  fmt.Println(s) // [], since there is an empty underlying array
  fmt.Println(s == nil) // also behaves as nil for easier checking
  ```

- defined

  ```go
  s := []int{1, 2, 3} // note the missing array length
  ```

- defined as a part of an existing array:

  ```go
  a := [3]int{1, 2, 3}
  s := a[0:2] // slice = [1, 2]
  ```

- using make:

  ```go
  s := make([]int, 3) // slice = [0, 0, 0]
  ```

Simmilar to arrays, slices can be passed to the `len()` function to get the length of the slice and to the `cap()` function to get the number of elements in the underlying array, counting from the first element in the slice.

```go
a := [5]int{1, 2, 3, 4, 5}
s := a[1:3] // slice = [2, 3]

fmt.Println(len(s)) // 2
fmt.Println(cap(s)) // 4
```

For iterations over elements - check [Range](#range)

```go
s := []int{1, 2, 3}
s = append(s, 4, 5)     // s = [1 2 3 4 5]

s1 := s[1:]             // s1 = [2 3 4 5]
s2 := s[:2]             // s2 = [1 2]
s3 := s[1:3]            // s3 = [2 3]
```

#### Maps

Maps are a collection of key-value pairs. The key and value types must be known at compile time. For iterations over elements - check [Range](#range)

```go
fileExtensions := map[string]string{
    "Golang":       ".go",
    "Java":         ".java",
    "JavaScript":   ".js",
    "Python"        "",
}

fileExtensions["Python"] = ".py"    // update a value

fileExtensions["PHP"] = ".php"      // add a new pair,  len(fileExtensions) = 5

delete(fileExtensions, "PHP")       // rempove a pair,  len(fileExtensions) = 4

val, ok := fileExtensions["Golang"] // check for a value, val = ".go"  ok = true
```

## Type conversions

When converting a variable of one type to another, a function with the following signature can be used: `T(v)`, where `T` is the new type and `v` is the value to convert.

```go
i := 1
f := float64(i)
u := uint(i)
```

## Flow control

### If

The parentheses `()` are optional, but the braces `{}` are required

```go
num := 15
if num%2 == 0 {
    fmt.Println(num, "is even")
} else {
    fmt.Println(num, "is odd")
}
```

Go does not support ternary operators, but it has a different feature. An `if` statement can start with a short statement to execute before the condition. The scope of the created variables is the if/if-else block

```go
if num := 9; num < 0 {
    fmt.Println(num, "is negative")
} else if num == 0 {
    fmt.Println(num, "is zero")
} else {
    fmt.Println(num, "is positive")
}
```

Usually this construct is used for error handling

```go
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Divide by zero")
	}
	return a / b, nil
}

func main() {

	if result, err := Divide(10, 0); err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Response", result)
	}

}

```

### Switch

Switch cases are checked from top to bottom and will execute the code in the first matching case only. If the `fallthrough` keyword is used, control will be passed to the next case.

```go
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
```

The `switch` operator can run without a condition: It's equivalent to if-then-else chains

```go
switch hour := time.Now().Hour(); { // missing expression means "true"
  case hour < 12:
    fmt.Println("Good morning!")
  case hour < 17:
    fmt.Println("Good afternoon!")
  default:
    fmt.Println("Good evening!")
  }
```

### For

There is only one keyword for all types of loops: `for`. The parentheses `()` are optional, but the braces `{}` are required

#### Single condition

```go
i := 1
for i < 5 {
    fmt.Println(i)
    i++
}
```

#### Initial/condition/after for loop

The variables declared in the init statement of `for` are visible only in the scope of the for statement

```go
for j := 5; j < 10; j++ {
    fmt.Println(j)
}
```

#### Infinite loop

- `for` without a condition will loop repeatedly until `break` or `return` are used. The `continue` keyword skips to the next iteration.

```go
n := 0
for {
	n++
	if n == 10 {
		break
	}
	if n%2 == 0 {
		continue
	}
	fmt.Println(n)
}
```

### Range

The `range` operator iterates over elements in a variety of data structures

#### On arrays and slices

On [arrays](#array) and [slices](#slice) `range` provides both the index and value for each entry

```go
slice := []int{1,2,3}

for i:=0; i < len(slice); i++{
    fmt.Println(i, slice[i])
}

for index, sliceValue := range slice{
    fmt.Println(index, sliceValue)
}

```

#### On maps

On a [map](#map) `range` iterates over the key/value pairs

```go
favColors := map[string]string {"John": "blue", "Markus": "red", "Anna": "green"}

for key, value := range favColors {
    fmt.Println("Name:", key, "Color:", value)
}
```

#### On strings

On strings `range` iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself.

```go
title := "GolangWorkshop"
for index, unicodeValue := range title {
    fmt.Println(index, unicodeValue)
    }
```

## Functions

Declared using the `func` keyword, followed by the function name, its parameters (if any), and its return values (if any)

```go
// takes two integers and returns their sum
func getSum(x int, y int) int {
	return x + y
}

// takes a slice of integers and returns the sum of all elements
func sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
```

### First-class citizens

Functions can be assigned to variables, passed as arguments to other functions, and returned as values from functions:

```go
func add(x, y int) int {
    return x + y
}

func main() {
    // assigning a function to a variable
    var sum func(int, int) int = add

    result := sum(3, 4)
    fmt.Println(result) // 7
}
```

### Closures

Functions can read and manipulate the state of their enclosing scope. This provides a clean and efficient way to encapsulate state and behavior within a function, making it easier to write modular and maintainable code.

```go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    // Creating a closure
    a := adder()

    fmt.Println(a(3)) // Output: 3
}
```

### Variadic parameters

Functions can have variadic parameters, which are denoted by using the `...` syntax before the parameter type:

```go
func vSum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

vResult := vSum(1, 2, 3, 4, 5, 6)

numbers := []int{1, 2, 3, 4, 5, 6}
vResult2 := vSum(numbers...) // used to "spread" the collection
```

### Returning multiple values

Golang allows multiple values to be returned from a single function. It's also possible to use named return values:

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}
```

### Passing functions as arguments

Go allows to pass functions as arguments to other functions

```go
func apply(f func(int) int, a int) int {
	return f(a)
}
func square(n int) int {
	return n * n
}
result := apply(square, 10)
fmt.Println(result)
```

### Deferred functions

The `defer` keyword schedules a function call to be executed when the surrounding function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Usage:

- Simplify cleanup tasks
- Ensure that resources are released after a function has completed
  - E.g.: open a file in a function, then defer the `Close` method of the file to be called when the functions returns

```go
func readFile() error {
	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scan1 := bufio.NewScanner(f)

	for scan1.Scan() {
		pl("Prime :", scan1.Text())
	}

	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
```

## Extra readings

### Online Guides and Articles

- [Official documentation](https://go.dev/) : Official documentation, definitely worth checking out the blog section
- [A tour of Go](https://go.dev/tour/welcome/1) : Online sandtool and tutorial by the Go developers explaining the main concepts
- [Go by example](https://gobyexample.com/) : List of code snippets illustrating a wide variety of topics
- [Take your first steps with Go](https://learn.microsoft.com/en-us/training/paths/go-first-steps/) : Introduction to Go by Microsoft
- [Go 101](https://go101.org/) : An online collection of articles, examples and quizzes in Go, can be downloaded as an ebook as well
- [Learn Go in 100 Lines](https://fireship.io/lessons/learn-go-in-100-lines/) : Learn Go in 100 lines
- [Practical Go Lessons](https://www.practical-go-lessons.com/) : Practical Go lessons with examples. Exists also as a book

### Books

- [How to code in Go](https://www.digitalocean.com/community/books/how-to-code-in-go-ebook) : extensive tutorial with several examples. Exists also as an online resource or free ebook (in epub and pdf)
- [For the Love of Go](https://bitfieldconsulting.com/books/love) : A book by John Arundel, which covers the fundamentals of Go and has an online video course
- [Go Style Guide](https://google.github.io/styleguide/go/index) : Article about best practices when writing in Go

### Videos

- [Go Path](https://app.pluralsight.com/paths/skill/go-core-language) : Series of extensive courses in Go in Pluralsight
- [Golang Dojo](https://www.youtube.com/@GolangDojo) : YouTube channel that offers many Go specific tutorials
- [NerdCademy](https://www.youtube.com/@NerdCademyDev) : Another YouTube channel dedicated to Go
- [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs) : Google I/O 2012 lecture on Go Concurrency Patterns
- [Learn Go in one video](https://www.youtube.com/watch?v=YzLrWHZa-Kc) : A course in Go

package main

import (
	utils "github.com/GolangWorkshop/golang-workshop/utilities"
)

func main() {

	//the SayHi() function is accessible from another package as its exported via the capital letter
	utils.SayHi()
	//uncommenting the following line will result in an error
	//utils.sayBye()    //the sayBye() function is not exported
}

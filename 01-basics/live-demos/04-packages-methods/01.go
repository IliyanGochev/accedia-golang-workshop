package main

import (
	"fmt"

	"github.com/GolangWorkshop/demo_project/models"
)

func main() {

	//Packages
	fmt.Println(models.GetAuthorId())
	id, rating := models.GetAuthorIdAndRating()
	fmt.Println(id, rating)

	// https://pkg.go.dev/
	// search for uuid
	fmt.Println(models.GetBookId())

}

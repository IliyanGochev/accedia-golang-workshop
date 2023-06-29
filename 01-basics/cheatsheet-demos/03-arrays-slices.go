package main

import (
	"fmt"
)

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	arr2 := [3]int{1, 2, 3}

	fmt.Println("arr == arr2:	", arr == arr2) // true!

	s := []int{1, 2, 3}

	for index, value := range s {
		fmt.Printf("s:	index: %v,	value: %v\n", index, value)
	}

	s = append(s, 4, 5)

	fmt.Println("s after appends:", s) // [1 2 3 4 5]

	s1 := s[1:]
	s2 := s[:2]
	s3 := s[1:3]

	fmt.Println("s1:", s1) // [2 3 4 5]
	fmt.Println("s2:", s2) // [1 2]
	fmt.Println("s3:", s3) // [2 3]
}

func main() {
	package main

import "fmt"

func main() {
	var arr [3]int
	fmt.Println(arr, len(arr), cap(arr))

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr, len(arr), cap(arr))

	var arr2 = [3]int{1, 2, 3}
	fmt.Println(arr, len(arr2), cap(arr2))


	var s []int
	fmt.Println(s, len(s), cap(s), s == nil)

	s2 := make([]int, 0, 5)
	fmt.Println(s2, len(s2), cap(s2), s2 == nil)

	s2 = append(s2, 1, 2, 3, 4, 5, 6)
	fmt.Println(s2, len(s2), cap(s2), s2 == nil)

	s3 := s2[:]
	fmt.Println(s3)

	s4 := s2[1:]
	fmt.Println(s4)

	s5 := s2[1:3]
	fmt.Println(s5)
}

}

package main

import (
	"fmt"
	"strings"
)

func decode(code string) string {
	code = strings.ToLower(code)
	for i := len(code) - 1; i >= 0; i-- {
		if code[i] >= 97 && code[i] <= 122 {
			fmt.Print(code[i] - 96)
		}

	}
	return code
}

func main() {
	fmt.Println(decode("AB"))
	fmt.Println(decode("Z!y X	"))
	fmt.Println(decode("Go"))
}

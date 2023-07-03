package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decode_seq(seq string) {

	var pos = 0
	var neg = 0
	splits := strings.Split(seq, " ")
	for _, value := range splits {
		intval, _ := strconv.Atoi(value)
		if intval >= 0 {
			pos = pos + intval
		} else {
			neg = neg + intval
		}
	}

	if pos+neg == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println(pos - (-neg))
	}
}

func main() {
	decode_seq("1 -1 12 1 1 1 -5 3 -7 -10 4 0")
	decode_seq("5 -4 12 -7 -3 -2 0")
	decode_seq("10 18 30 0")
}

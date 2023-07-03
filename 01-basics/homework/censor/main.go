package main

import (
	"fmt"
	"strings"
)

type pair struct {
	coded   string
	missing string
}

func (p pair) Decode() {
	for _, val := range p.missing {
		char := fmt.Sprintf("%c", val)
		p.coded = strings.Replace(p.coded, "*", char, 1)
	}
	fmt.Println(p.coded)
}

func main() {
	inputs := [3]pair{
		{coded: "*h*s *s v*ry *tr*ng*", missing: "Tiiesae"},
		{coded: "A**Z*N*", missing: "MAIG"},
		{coded: "xyz", missing: ""},
	}

	for _, input := range inputs {
		input.Decode()
	}
}

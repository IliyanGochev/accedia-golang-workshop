package main

import (
	"fmt"
	"strings"
)

type Lamp struct {
	color string
	state string
}

func checkColor(color string) string {
	result := "black"
	switch color {
	case "red":
		result = "red"
	case "green":
		result = "green"
	case "orange":
		result = "orange"
	case "yellow":
		result = "yellow"
	case "blue":
		result = "blue"
	case "purple":
		result = "purple"
	}
	return result
}

func checkState(state string) string {
	result := "off"
	switch state {
	case "on":
		result = "on"
	}
	return result
}

func NewLamp(color string, state string) Lamp {
	var lamp = Lamp{color: checkColor(color), state: checkState(state)}
	return lamp
}

func (l Lamp) GetColor() string {
	return l.color
}

func (l *Lamp) SetColor(color string) {
	l.color = checkColor(color)
}

func (l Lamp) GetState() string {
	return l.state
}

func (l *Lamp) SetState(state string) {
	l.state = checkState(state)
}

func (l Lamp) SummonGenie() bool {
	result := false
	if l.color == "red" && l.state == "on" {
		result = true
	}
	return result
}

func (l Lamp) String() string {
	return fmt.Sprintf("%s %s", l.color, l.state)
}

func main() {
	var inputs = []string{
		"blue on purple off red on",
		"orange off red off turquoise on",
		"yellow purple on green off on",
	}

	for _, input := range inputs {
		commands := strings.Split(input, " ")
		lamp := NewLamp("", "")
		for idx, command := range commands {
			switch command {
			case "on", "off":
				lamp.SetState(command)
			default:
				lamp.SetColor(command)
			}

			fmt.Println(idx, lamp.String())
		}

		if lamp.SummonGenie() {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

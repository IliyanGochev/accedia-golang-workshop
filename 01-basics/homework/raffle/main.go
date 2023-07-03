package main

import (
	"fmt"
	"strconv"
	"strings"
)

func process_ticket(ticket_id int) {
	var found_digit = 0
	digits := strconv.Itoa(ticket_id)
	for _, val := range digits {
		v := fmt.Sprintf("%c", val)
		if strings.Count(digits, v) > 1 {
			found_digit, _ = strconv.Atoi(v)
		}
	}
	fmt.Println(found_digit)
}

func main() {
	process_ticket(234256)
	process_ticket(676678)
	process_ticket(2345)
}

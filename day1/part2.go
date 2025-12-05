package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	text := string(input)

	position := 50

	zero_stops := 0

	var direction string
	var clicks int

	for line := range strings.Lines(text) {
		line = strings.TrimSuffix(line, "\n")

		direction = line[:1]
		clicks, _ = strconv.Atoi(line[1:])

		for i := 0; i < clicks; i++ {
			if position == 0 {
				zero_stops += 1
			}

			if direction == "L" {
				if position == 0 {
					position = 99
				} else {
					position -= 1
				}
			} else {
				if position == 99 {
					position = 0
				} else {
					position += 1
				}
			}
		}

	}

	fmt.Println(zero_stops)
}

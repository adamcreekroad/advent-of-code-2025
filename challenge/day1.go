package challenge

import (
	"os"
	"strconv"
	"strings"
)

type DayOne struct {
	input string
}

func NewDayOne() *DayOne {
	bytes, _ := os.ReadFile("input/day1.txt")
	input := string(bytes)

	return &DayOne{input}
}

func (d *DayOne) PartOne() int {
	position := 50

	zero_stops := 0

	var direction string
	var clicks int

	for line := range strings.Lines(d.input) {
		line = strings.TrimSuffix(line, "\n")

		direction = line[:1]
		clicks, _ = strconv.Atoi(line[1:])

		for i := 0; i < clicks; i++ {

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

		if position == 0 {
			zero_stops += 1
		}
	}

	return zero_stops
}

func (d *DayOne) PartTwo() int {
	position := 50

	zero_stops := 0

	var direction string
	var clicks int

	for line := range strings.Lines(d.input) {
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

	return zero_stops
}

package challenge

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type DayThree struct {
	input string
}

func NewDayThree() *DayThree {
	bytes, _ := os.ReadFile("input/day3.txt")
	input := string(bytes)

	return &DayThree{input}
}

type battery struct {
	index  int
	charge int
	str    string
}

func (d *DayThree) PartOne() int {
	var result int
	var batteries []battery

	for bank := range strings.Lines(d.input) {
		batteries = nil

		for index, digit_str := range bank {
			digit, _ := strconv.Atoi(string(digit_str))

			batteries = append(batteries, battery{index, digit, string(digit_str)})
		}

		compare := func(a, b battery) int {
			return cmp.Compare(a.charge, b.charge)
		}

		l_max := slices.MaxFunc(batteries[:len(batteries)-2], compare)
		r_max := slices.MaxFunc(batteries[l_max.index+1:], compare)

		joltage, _ := strconv.Atoi(fmt.Sprintf("%d%d", l_max.charge, r_max.charge))

		result += joltage
	}

	return result
}

func (d *DayThree) PartTwo() int {
	var result int
	var batteries []battery
	var j_batteries []battery

	for bank := range strings.Lines(d.input) {
		batteries = nil
		j_batteries = nil

		for index, digit_str := range bank {
			digit, _ := strconv.Atoi(string(digit_str))

			batteries = append(batteries, battery{index, digit, string(digit_str)})
		}

		compare := func(a, b battery) int {
			return cmp.Compare(a.charge, b.charge)
		}

		min_index := 0
		max_index := len(batteries) - 1

		var joltage_str string
		for i := 11; i >= 0; i-- {
			max := slices.MaxFunc(batteries[min_index:max_index-i], compare)

			j_batteries = append(j_batteries, max)
			joltage_str += max.str
			min_index = max.index + 1
		}

		joltage, _ := strconv.Atoi(joltage_str)

		result += joltage
	}

	return result
}

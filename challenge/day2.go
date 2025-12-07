package challenge

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type DayTwo struct {
	input string
}

func NewDayTwo() *DayTwo {
	bytes, _ := os.ReadFile("input/day2.txt")
	input := string(bytes)

	return &DayTwo{input}
}

func (d *DayTwo) PartOne() int {
	id_regex, _ := regexp.Compile("([0-9]+)-([0-9]+)")

	id_ranges := strings.Split(d.input, ",")

	invalid_sum := 0

	for _, id_range := range id_ranges {
		scan := id_regex.FindAllStringSubmatch(id_range, 2)

		id_start, _ := strconv.Atoi(scan[0][1])
		id_end, _ := strconv.Atoi(scan[0][2])

		for i := id_start; i <= id_end; i++ {
			string_id := strconv.Itoa(i)
			digits := len(string_id)

			if digits%2 == 0 {
				middle := digits / 2

				pre := string_id[middle:]
				post := string_id[:middle]

				if pre == post {
					invalid_sum += i
				}
			}
		}
	}

	return invalid_sum
}

func (d *DayTwo) PartTwo() int {
	regexes := map[int]*regexp.Regexp{}

	id_regex, _ := regexp.Compile("([0-9]+)-([0-9]+)")

	id_ranges := strings.Split(d.input, ",")

	invalid_sum := 0

	for _, id_range := range id_ranges {
		scan := id_regex.FindAllStringSubmatch(id_range, 2)

		id_start, _ := strconv.Atoi(scan[0][1])
		id_end, _ := strconv.Atoi(scan[0][2])

	id:
		for id := id_start; id <= id_end; id++ {
			string_id := strconv.Itoa(id)
			digits := len(string_id)

		id_split:
			for i := 2; i <= digits; i++ {
				if digits%i == 0 {
					regex := findOrCompileRegex(regexes, digits/i)

					result := regex.FindAllStringSubmatch(string_id, i)

					for _, sub_id := range result {
						if sub_id[0] == result[0][0] {
							continue
						} else {
							continue id_split
						}
					}

					invalid_sum += id
					continue id
				}
			}
		}
	}

	return invalid_sum
}

func findOrCompileRegex(regexes map[int]*regexp.Regexp, digits int) *regexp.Regexp {
	existing, ok := regexes[digits]
	if ok {
		return existing
	}

	pattern := fmt.Sprintf("[0-9]{%d}", digits)
	compiled, _ := regexp.Compile(pattern)
	regexes[digits] = compiled

	return compiled
}

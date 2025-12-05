package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regexes map[int]*regexp.Regexp

func main() {
	regexes = map[int]*regexp.Regexp{}

	input, _ := os.ReadFile("input.txt")

	text := string(input)

	id_regex, _ := regexp.Compile("([0-9]+)-([0-9]+)")

	id_ranges := strings.Split(text, ",")

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
					regex := findOrCompileRegex(digits / i)

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

	fmt.Println(invalid_sum)
}

func findOrCompileRegex(digits int) *regexp.Regexp {
	existing, ok := regexes[digits]
	if ok {
		return existing
	}

	pattern := fmt.Sprintf("[0-9]{%d}", digits)
	compiled, _ := regexp.Compile(pattern)
	regexes[digits] = compiled

	return compiled
}

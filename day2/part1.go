package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	text := string(input)

	id_regex, _ := regexp.Compile("([0-9]+)-([0-9]+)")

	id_ranges := strings.Split(text, ",")

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

	fmt.Println(invalid_sum)
}

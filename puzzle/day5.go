package puzzle

import (
	"os"
	"strconv"
	"strings"
)

type DayFive struct {
	input string
}

func NewDayFive() *DayFive {
	bytes, _ := os.ReadFile("input/day5.txt")
	input := string(bytes)

	return &DayFive{input}
}

func (d *DayFive) PartOne() int {
	var result int

	split := strings.Split(d.input, "\n\n")

	fresh := [][2]int{}
	for fresh_range := range strings.SplitSeq(split[0], "\n") {
		split := strings.Split(fresh_range, "-")

		min, _ := strconv.Atoi(split[0])
		max, _ := strconv.Atoi(split[1])

		fresh = append(fresh, [2]int{min, max})
	}

id_iter:
	for id_str := range strings.SplitSeq(split[1], "\n") {
		id, _ := strconv.Atoi(id_str)

		for _, fresh_range := range fresh {
			if id > fresh_range[0] && id < fresh_range[1] {
				result += 1
				continue id_iter
			}
		}
	}

	return result
}

func (d *DayFive) PartTwo() int {
	var result int

	split := strings.Split(d.input, "\n\n")

	fresh := [][2]int{}
	for fresh_range_str := range strings.SplitSeq(split[0], "\n") {
		split := strings.Split(fresh_range_str, "-")

		min_id, _ := strconv.Atoi(split[0])
		max_id, _ := strconv.Atoi(split[1])

		fresh = append(fresh, [2]int{min_id, max_id})
	}

	consolidated := make([][2]int, len(fresh))
	copy(consolidated, fresh)
	var removed int

	for {
		consolidated, removed = consolidateRanges(consolidated)

		if removed == 0 {
			break
		}
	}

	for _, fresh_range := range consolidated {
		count := fresh_range[1] - fresh_range[0] + 1

		result += count
	}

	return result
}

func consolidateRanges(fresh [][2]int) ([][2]int, int) {
	consolidated := [][2]int{}
	removed := 0

	var min_id int
	var max_id int

outer:
	for _, fresh_range := range fresh {
		min_id = fresh_range[0]
		max_id = fresh_range[1]

		for index, con_range := range consolidated {
			if min_id > con_range[1] || max_id < con_range[0] {
				continue
			}

			consolidated[index][0] = min(min_id, con_range[0])
			consolidated[index][1] = max(max_id, con_range[1])
			removed += 1
			continue outer
		}

		consolidated = append(consolidated, [2]int{min_id, max_id})
	}

	return consolidated, removed
}

package puzzle

import (
	"os"
	"strings"
)

type DayFour struct {
	input string
}

const (
	at_rune = 64
)

func NewDayFour() *DayFour {
	bytes, _ := os.ReadFile("input/day4.txt")
	input := string(bytes)

	return &DayFour{input}
}

func (d *DayFour) PartOne() int {
	var result int

	input := strings.ReplaceAll(d.input, "\n", "")

	check_pos := [8][2]int{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}

grid_pos:
	for index, char := range input {
		if char != at_rune {
			continue
		}

		x := index % 140
		y := index / 140

		adjacent_paper := 0

		for _, pos := range check_pos {
			check_x := x + pos[0]
			check_y := y + pos[1]

			if check_x < 0 || check_y < 0 || check_x > 139 || check_y > 139 {
				continue
			}

			if input[index+pos[0]+(pos[1]*140)] == at_rune {
				adjacent_paper += 1
			}

			if adjacent_paper == 4 {
				continue grid_pos
			}
		}

		result += 1
	}

	return result
}

func (d *DayFour) PartTwo() int {
	var result int

	grid := strings.ReplaceAll(d.input, "\n", "")

	for {
		new_grid, paper_removed := removePaper(grid)

		result += paper_removed

		if new_grid == grid {
			break
		} else {
			grid = new_grid
		}
	}

	return result
}

func removePaper(grid string) (string, int) {
	new_grid := ""
	paper_removed := 0

	check_pos := [8][2]int{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}

grid_pos:
	for index, char := range grid {
		if char != at_rune {
			new_grid += string(char)
			continue
		}

		x := index % 140
		y := index / 140

		adjacent_paper := 0

		for _, pos := range check_pos {
			check_x := x + pos[0]
			check_y := y + pos[1]

			if check_x < 0 || check_y < 0 || check_x > 139 || check_y > 139 {
				continue
			}

			if grid[index+pos[0]+(pos[1]*140)] == at_rune {
				adjacent_paper += 1
			}

			if adjacent_paper == 4 {
				new_grid += string(char)
				continue grid_pos
			}
		}

		new_grid += "."
		paper_removed += 1
	}

	return new_grid, paper_removed
}

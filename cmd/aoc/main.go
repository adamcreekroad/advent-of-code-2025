package main

import (
	"flag"
	"fmt"

	"github.com/adamcreekroad/advent-of-code-2025/challenge"
)

var day int
var part int

func init() {
	flag.IntVar(&day, "d", 1, "Which day to run")
	flag.IntVar(&part, "p", 1, "Which part to run")

	flag.Parse()
}

func main() {
	fmt.Printf("Running day %d part %d...\n\n", day, part)

	handler := challenge.NewDay(day)

	var result int

	if part == 1 {
		result = handler.PartOne()
	} else {
		result = handler.PartTwo()
	}

	fmt.Println("The result is:", result)
}

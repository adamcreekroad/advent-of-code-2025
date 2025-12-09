package puzzle

type Puzzle interface {
	PartOne() int
	PartTwo() int
}

func NewDay(day int) Puzzle {
	var handler Puzzle

	switch day {
	case 1:
		handler = NewDayOne()
	case 2:
		handler = NewDayTwo()
	case 3:
		handler = NewDayThree()
	case 4:
		handler = NewDayFour()
	case 5:
		handler = NewDayFive()
	}

	return handler
}

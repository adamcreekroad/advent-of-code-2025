package challenge

type Challenge interface {
	PartOne() int
	PartTwo() int
}

func NewDay(day int) Challenge {
	var handler Challenge

	switch day {
	case 1:
		handler = NewDayOne()
	case 2:
		handler = NewDayTwo()
	}

	return handler
}

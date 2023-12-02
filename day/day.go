package day

import "time"

// Day is an interface for each day's solution to implement
type Day interface {
	ReadFile(path string) error
	Part1()
	Part2()
	GetDayNumber() int
}

// GetToday calculates the current day of the month
// and returns the solution for that day
func GetToday() Day {
	_, _, d := time.Now().Date()
	return GetDay(d)
}

// GetDay takes a day of the month and returns the
// solution for that day
func GetDay(d int) Day {
	days := GetDays()
	return days[d-1]
}

// GetDays returns all the solutions in order
func GetDays() []Day {
	return []Day{
		&Day1{},
	}
}

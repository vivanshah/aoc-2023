package day

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Day4 struct {
	Entries []string
}

func (d *Day4) GetDayNumber() int {
	return 4
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day4) ReadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	d.Entries = []string{}
	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		line = strings.TrimSuffix(line, "\n")
		d.Entries = append(d.Entries, line)
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return err
	}
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day4) Part1() {
	fmt.Println("Day 4 Part 1")
	var points int
	points = 0
	for row, l := range d.Entries {
		cardParts := strings.Split(l, ": ")
		numberParts := strings.Split(cardParts[1], " | ")
		wn := strings.Fields(numberParts[0])
		winningNumbers := map[int]bool{}
		for _, n := range wn {
			winningNumber, _ := strconv.Atoi(n)
			winningNumbers[winningNumber] = true

		}

		mn := strings.Fields(numberParts[1])
		myNumbers := map[int]bool{}
		for _, n := range mn {
			myNumber, _ := strconv.Atoi(n)
			myNumbers[myNumber] = true
		}
		fmt.Println(winningNumbers, " | ", myNumbers)
		matches := 0
		for k, _ := range myNumbers {
			if winningNumbers[k] {
				matches++
			}
		}
		fmt.Println("Card ", row+1, " has ", matches, "matches")
		cardPoints := int(math.Pow(float64(2), float64(matches-1)))
		fmt.Println("Card ", row+1, " worth ", cardPoints, "points")
		points += cardPoints
	}
	fmt.Println(points)
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day4) Part2() {
	fmt.Println("Day 4 Part 2")
	cardCount := 0
	cardCopies := make([]int, len(d.Entries))
	for row, l := range d.Entries {
		cardParts := strings.Split(l, ": ")
		numberParts := strings.Split(cardParts[1], " | ")
		wn := strings.Fields(numberParts[0])
		winningNumbers := map[int]bool{}
		for _, n := range wn {
			winningNumber, _ := strconv.Atoi(n)
			winningNumbers[winningNumber] = true
		}
		mn := strings.Fields(numberParts[1])
		myNumbers := map[int]bool{}
		for _, n := range mn {
			myNumber, _ := strconv.Atoi(n)
			myNumbers[myNumber] = true
		}
		matches := 0
		for k, _ := range myNumbers {
			if winningNumbers[k] {
				matches++
			}
		}
		cardCopies[row] += 1

		for i := 0; i < cardCopies[row]; i++ {
			for m := 1; m <= matches; m++ {
				cardCopies[row+m]++
			}
		}
	}

	for _, p := range cardCopies {
		cardCount += p
	}
	fmt.Println(cardCount)
}

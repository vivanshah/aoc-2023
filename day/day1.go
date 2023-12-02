package day

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Day1 struct {
	Entries []string
}

func (d *Day1) GetDayNumber() int {
	return 1
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day1) ReadFile(path string) error {
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
func (d *Day1) Part1() {
	fmt.Println("Day 1 Part 1")
	sum := 0
	for _, s := range d.Entries {
		lineSum := []byte{}
		for i := 0; i < len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' {
				lineSum = append(lineSum, s[i])
				break
			}
		}
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] >= '0' && s[i] <= '9' {
				lineSum = append(lineSum, s[i])
				break
			}
		}
		ls, _ := strconv.Atoi(string(lineSum))
		sum += ls
	}
	fmt.Println(sum)
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day1) Part2() {
	fmt.Println("Day 1 Part 2")
	sum := 0
	for _, s := range d.Entries {
		lineSum := []byte{}
		for i := 0; i < len(s); i = i {
			num := isNumWord(s, i)
			if num >= 0 {
				lineSum = append(lineSum, byte(num)+'0')
				break
			}
			if s[i] >= '0' && s[i] <= '9' {
				lineSum = append(lineSum, s[i])
				break
			}
			i++
		}
		for i := len(s) - 1; i >= 0; i-- {
			num := isNumWord(s, i)
			if num >= 0 {
				lineSum = append(lineSum, byte(num)+'0')
				break
			}
			if s[i] >= '0' && s[i] <= '9' {
				lineSum = append(lineSum, s[i])
				break
			}
		}

		ls, _ := strconv.Atoi(string(lineSum))
		sum += ls
	}
	fmt.Println(sum)
}

func isNumWord(s string, i int) int {
	numWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for n, w := range numWords {
		if i+len(w) < len(s) {
			if strings.Index(s[i:i+len(w)], w) == 0 {
				return n + 1
			}
		}
	}
	return -1

}

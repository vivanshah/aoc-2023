package day

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Day2 struct {
	Entries []string
}

func (d *Day2) GetDayNumber() int {
	return 2
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day2) ReadFile(path string) error {
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
func (d *Day2) Part1() {
	fmt.Println("Day 2 Part 1")
	sum := 0
	for _, line := range d.Entries {
		if len(line) == 0 {
			continue
		}
		game := parseGame(line)
		if checkGame(game, 12, 13, 14) {
			sum += game.gameNumber
		}
	}
	fmt.Println(sum)
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day2) Part2() {
	fmt.Println("Day 2 Part 2")
	sum := 0
	for _, line := range d.Entries {
		if len(line) == 0 {
			continue
		}
		game := parseGame(line)
		mR, mG, mB := minCounts(game)
		sum += mR * mG * mB
	}
	fmt.Println(sum)
}

type game struct {
	gameNumber int
	rounds     []round
}
type round struct {
	red, green, blue int
}

func checkGame(g game, red int, green int, blue int) bool {

	for _, round := range g.rounds {
		if round.red > red || round.blue > blue || round.green > green {
			//fmt.Println("invalid game", g)
			return false
		}
	}
	return true

}

func parseGame(line string) game {
	game := game{}
	parts := strings.Split(line, ": ")
	game.gameNumber, _ = strconv.Atoi(strings.ReplaceAll(parts[0], "Game ", ""))
	rounds := strings.Split(parts[1], "; ")
	for _, p := range rounds {
		round := round{}
		p = strings.TrimSpace(p)
		counts := strings.Split(p, ", ")
		for _, c := range counts {
			c = strings.TrimSpace(c)
			bits := strings.Split(c, " ")
			count, _ := strconv.Atoi(bits[0])
			switch {
			case bits[1] == "blue":
				round.blue = count
				break
			case bits[1] == "red":
				round.red = count
				break
			case bits[1] == "green":
				round.green = count
				break
			}
		}
		game.rounds = append(game.rounds, round)
	}
	return game
}

func minCounts(g game) (int, int, int) {
	var mR, mG, mB int
	for _, r := range g.rounds {
		if r.red > mR {
			mR = r.red
		}
		if r.green > mG {
			mG = r.green
		}
		if r.blue > mB {
			mB = r.blue
		}
	}
	return mR, mG, mB
}

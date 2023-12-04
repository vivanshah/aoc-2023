package day

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Day3 struct {
	Entries     []string
	Grid        [][]bool
	GearGrid    [][]Gear
	AllGears    []Gear
	PartNumbers []int
}
type Gear struct {
	AdjacentParts map[int]bool
}

func (d *Day3) Ratio(g Gear) int {
	if len(g.AdjacentParts) != 2 {
		return 0
	}
	product := 1
	for k := range g.AdjacentParts {
		product = product * d.PartNumbers[k]
	}
	return product
}

func (d *Day3) GetDayNumber() int {
	return 3
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day3) ReadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	d.Entries = []string{}
	d.Grid = [][]bool{}
	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	for i := 0; ; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		line = strings.TrimSuffix(line, "\n")
		d.Entries = append(d.Entries, line)
		d.Grid = append(d.Grid, make([]bool, len(line)))
		d.GearGrid = append(d.GearGrid, make([]Gear, len(line)))
		for col, c := range line {
			if string(c) == "." || (line[col] >= '0' && line[col] <= '9') {
				d.Grid[i][col] = false
			} else {
				d.Grid[i][col] = true
				if string(c) == "*" {
					ng := Gear{AdjacentParts: map[int]bool{}}
					d.GearGrid[i][col] = ng
					d.AllGears = append(d.AllGears, ng)
				}
			}
		}
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
func (d *Day3) Part1() {
	fmt.Println("Day 3 Part 1")
	inNum := false
	currNum := []byte{}
	nearSymbol := false
	sum := 0
	for row, l := range d.Entries {
		for col := 0; col < len(l); col++ {
			if l[col] >= '0' && l[col] <= '9' {
				inNum = true
				currNum = append(currNum, l[col])
				if nearSymbol == false {
					nearSymbol = d.isNearSymbol(row, col)
				}
				continue
			} else {
				if inNum && nearSymbol {
					num, _ := strconv.Atoi(string(currNum))
					sum += num
				}
				inNum = false
				nearSymbol = false
				currNum = []byte{}

			}
		}
	}
	fmt.Println(sum)

}

// Part2 executes part 2 of of this day's puzzle
func (d *Day3) Part2() {
	fmt.Println("Day 3 Part 2")
	inNum := false
	currNum := []byte{}
	nearSymbol := false
	sum := 0
	for row, l := range d.Entries {
		for col := 0; col < len(l); col++ {
			if l[col] >= '0' && l[col] <= '9' {
				inNum = true
				currNum = append(currNum, l[col])
				if nearSymbol == false {
					nearSymbol = d.addToGears(row, col, len(d.PartNumbers))
				}
				continue
			} else {
				var num int
				if inNum && nearSymbol {
					num, _ = strconv.Atoi(string(currNum))
					sum += num
					d.PartNumbers = append(d.PartNumbers, num)
				}
				inNum = false
				nearSymbol = false
				currNum = []byte{}

			}
		}
	}
	ratioProduct := 0
	for _, v := range d.AllGears {
		ratioProduct += d.Ratio(v)
	}
	fmt.Println(ratioProduct)

}

func (d *Day3) isNearSymbol(row int, col int) bool {
	numRows := len(d.Grid)
	numCols := len(d.Grid[0])
	for dRow := -1; dRow <= 1; dRow++ {
		for dCol := -1; dCol <= 1; dCol++ {
			newRow := row + dRow
			newCol := col + dCol

			if newRow >= 0 && newRow < numRows && newCol >= 0 && newCol < numCols {
				if newRow != row || newCol != col {
					if d.Grid[newRow][newCol] {
						return true
					}
				}
			}
		}
	}
	return false
}

func (d *Day3) addToGears(row int, col int, partNum int) bool {
	numRows := len(d.Grid)
	numCols := len(d.Grid[0])
	for dRow := -1; dRow <= 1; dRow++ {
		for dCol := -1; dCol <= 1; dCol++ {
			newRow := row + dRow
			newCol := col + dCol

			if newRow >= 0 && newRow < numRows && newCol >= 0 && newCol < numCols {
				if newRow != row || newCol != col {
					if d.Grid[newRow][newCol] {
						if string(d.Entries[newRow][newCol]) == "*" {
							//i'm a gear!
							d.GearGrid[newRow][newCol].AdjacentParts[partNum] = true
						}
						return true
					}

				}
			}
		}
	}
	return false
}

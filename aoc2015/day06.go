package aoc2015

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Day06 solves the sixth day problem
// "Probably a Fire Hazard"
func Day06(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	var countLit, brightness uint64 // countLit is Part 1; brightness is Part 2
	// generate powerGrid (0 is off; 1 is on)
	// so powerGrid[x][y] represents the x'th lightbulb in y'th row
	powerGrid := make([][]bool, 1000)
	nordicGrid := make([][]uint64, 1000)
	for ind := range powerGrid {
		powerGrid[ind] = make([]bool, 1000)
		nordicGrid[ind] = make([]uint64, 1000)
	}

	// manipulation functions
	turnOn := func(grid [][]bool, nordGrid [][]uint64, row, col int) {
		grid[row][col] = true
		nordGrid[row][col]++
	}
	turnOff := func(grid [][]bool, nordGrid [][]uint64, row, col int) {
		grid[row][col] = false
		if nordGrid[row][col] > 0 {
			nordGrid[row][col]--
		}
	}
	toggle := func(grid [][]bool, nordGrid [][]uint64, row, col int) {
		grid[row][col] = !grid[row][col]
		nordGrid[row][col] += 2
	}
	parseRowCol := func(dimStr string, row, col *int) error {
		// dimStr should be of form "ROW,COL"
		dims := strings.Split(dimStr, ",")
		if len(dims) != 2 {
			return fmt.Errorf("%q not comma separated", dims)
		}
		var err error
		if *row, err = strconv.Atoi(dims[0]); err != nil {
			return fmt.Errorf("unknown row %q in query %q", dims[0], dimStr)
		}
		if *col, err = strconv.Atoi(dims[1]); err != nil {
			return fmt.Errorf("unknown col %q in query %q", dims[1], dimStr)
		}
		return nil
	}

	// now read each line in scanner
	for scanner.Scan() {
		rawQuery := scanner.Text()
		query := strings.Fields(rawQuery)
		// note that query can either be one of the following
		// toggle ROW1,COL1 through ROW2,COL2
		// turn off ROW1,COL1 through ROW2,COL2
		// turn on ROW1,COL1 through ROW2,COL2
		// so query is of length 4 or 5
		if len(query) != 4 && len(query) != 5 {
			err = fmt.Errorf("unknown spacing: %q", rawQuery)
			return
		}
		var row1, col1, row2, col2 int
		var manipFn func(grid [][]bool, nordGrid [][]uint64, row, col int)
		if query[0] == "toggle" {
			if err = parseRowCol(query[1], &row1, &col1); err != nil {
				return
			}
			if err = parseRowCol(query[3], &row2, &col2); err != nil {
				return
			}
			manipFn = toggle
		} else if query[0] == "turn" {
			if err = parseRowCol(query[2], &row1, &col1); err != nil {
				return
			}
			if err = parseRowCol(query[4], &row2, &col2); err != nil {
				return
			}
			if query[1] == "off" {
				manipFn = turnOff
			} else if query[1] == "on" {
				manipFn = turnOn
			} else {
				err = fmt.Errorf("unknown function: %q in %q neither off nor on", query[1], rawQuery)
				return
			}
		} else {
			return "", "", fmt.Errorf("unknown function: %q in %q", query[0], rawQuery)
		}
		// now this is a really bad for loop
		for row := row1; row <= row2; row++ {
			for col := col1; col <= col2; col++ {
				manipFn(powerGrid, nordicGrid, row, col)
			}
		}
	}

	// now count
	for row := range powerGrid {
		for col := range powerGrid[row] {
			if powerGrid[row][col] {
				countLit++
			}
			brightness += nordicGrid[row][col]
		}
	}

	answer1 = strconv.FormatUint(countLit, 10)
	answer2 = strconv.FormatUint(brightness, 10)
	return
}

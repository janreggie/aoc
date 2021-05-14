package aoc2020

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// seat represents a seat on the waiting area
type seat byte

const (
	floor    seat = '.'
	empty    seat = 'L'
	occupied seat = '#'
	oob      seat = ' ' // out of bounds
)

// makeSeatRow generates a row of seats from a string
func makeSeatRow(row string) ([]seat, error) {
	result := make([]seat, len(row))
	for ii := range row {
		switch cur := seat(row[ii]); cur {
		case floor:
			result[ii] = floor
		case empty:
			result[ii] = empty
		case occupied: // this shouldn't happen
			result[ii] = occupied
		default:
			err := errors.Errorf("could not create row: invalid byte `%c`", cur)
			return nil, err
		}
	}
	return result, nil
}

// waitingArea represents the state of a waiting area
type waitingArea struct {
	representation [][]seat

	width  int
	height int
	stable bool // has the waiting area achieved stability?
}

func (area waitingArea) String() string {
	var sb strings.Builder
	for yy := range area.representation {
		for xx := range area.representation[yy] {
			switch seat := area.representation[yy][xx]; seat {
			case floor:
				sb.WriteByte('.')
			case empty:
				sb.WriteByte('L')
			case occupied:
				sb.WriteByte('#')
			default:
				sb.WriteByte('?')
			}
		}
		sb.WriteByte('\n')
	}
	fmt.Fprintf(&sb, "width: %v, height: %v, stable: %v", area.width, area.height, area.stable)
	return sb.String()
}

// generateWaitingArea generates a waitingArea from a scanner
func generateWaitingArea(scanner *bufio.Scanner) (waitingArea, error) {
	var result waitingArea
	result.representation = make([][]seat, 0)
	for scanner.Scan() {
		row, err := makeSeatRow(scanner.Text())
		if err != nil {
			return result, errors.Wrapf(err, "could not parse invalid row %s", scanner.Text())
		}
		if result.width == 0 {
			result.width = len(row)
		}
		if len(row) != result.width {
			return result, errors.Wrapf(err, "row %s of unequal width, should be of width %d", scanner.Text(), result.width)
		}
		result.representation = append(result.representation, row)
	}
	result.height = len(result.representation)
	return result, nil
}

// copy returns a copy of itself
func (area waitingArea) copy() waitingArea {
	rep := make([][]seat, area.height)
	for ii := range rep {
		rep[ii] = make([]seat, area.width)
		copy(rep[ii], area.representation[ii])
	}
	return waitingArea{
		representation: rep,
		width:          area.width,
		height:         area.height,
		stable:         area.stable,
	}
}

// isStable returns if area is stable
func (area waitingArea) isStable() bool {
	return area.stable
}

// at returns the seat at column x and row y. If out of bounds, return oob.
func (area waitingArea) at(x, y int) seat {
	if x < 0 || x >= area.width || y < 0 || y >= area.height {
		return oob
	}
	return area.representation[y][x]
}

// setAt sets value at column x and row y. If out of bounds, do nothing.
func (area waitingArea) setAt(x, y int, val seat) {
	if x < 0 || x >= area.width || y < 0 || y >= area.height {
		return
	}
	area.representation[y][x] = val
}

// statistics returns how many seats are empty, occupied, and floor respectively
func (area waitingArea) statistics() (int, int, int) {
	e, o, f := 0, 0, 0
	for yy := 0; yy < area.height; yy++ {
		for xx := 0; xx < area.width; xx++ {
			switch area.at(xx, yy) {
			case empty:
				e++
			case occupied:
				o++
			case floor:
				f++
			}
		}
	}
	return e, o, f
}

// countSurroundingOccupied returns how many occupied seats are there surrounding (x,y)
func (area waitingArea) countSurroundingOccupied(x, y int) int {
	result := 0
	possible := []struct{ m, n int }{
		{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1},
		{x - 1, y}, {x + 1, y},
		{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
	}
	for _, coord := range possible {
		if area.at(coord.m, coord.n) == occupied {
			result++
		}
	}
	return result
}

// countVisibleOccupied returns how many occupied seats are there that is directly visible
// from one of the eight directions from (x,y)
func (area waitingArea) countVisibleOccupied(x, y int) int {
	result := 0
	directions := []struct{ i, j int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	for _, dir := range directions {
		xx, yy := x+dir.i, y+dir.j
		for area.at(xx, yy) == floor {
			xx += dir.i
			yy += dir.j
		}
		if area.at(xx, yy) == occupied {
			result++
		}
	}
	return result
}

// equal checks if two waitingAreas have the same width, height, and seats
func (area waitingArea) equal(aa waitingArea) bool {
	if area.width != aa.width || area.height != aa.height {
		return false
	}
	for yy := 0; yy < area.height; yy++ {
		for xx := 0; xx < area.width; xx++ {
			if aa.at(xx, yy) != area.at(xx, yy) {
				return false
			}
		}
	}
	return true
}

// iterateSimple returns the waitingArea after one round of the rules provided in part 1.
// Returns a copy of itself if already stable
func (area waitingArea) iterateSimple() waitingArea {
	if area.stable {
		return area.copy()
	}

	nextArea := waitingArea{
		width:  area.width,
		height: area.height,
		stable: area.stable,
	}
	nextArea.representation = make([][]seat, area.height)
	for ii := range nextArea.representation {
		nextArea.representation[ii] = make([]seat, area.width)
	}

	for yy := 0; yy < area.height; yy++ {
		for xx := 0; xx < area.width; xx++ {
			currentSeat := area.at(xx, yy)
			switch currentSeat {
			case floor:
				nextArea.setAt(xx, yy, floor)
			case empty:
				if area.countSurroundingOccupied(xx, yy) == 0 {
					nextArea.setAt(xx, yy, occupied)
				} else {
					nextArea.setAt(xx, yy, empty)
				}
			case occupied:
				if area.countSurroundingOccupied(xx, yy) >= 4 {
					nextArea.setAt(xx, yy, empty)
				} else {
					nextArea.setAt(xx, yy, occupied)
				}
			}
		}
	}

	if area.equal(nextArea) {
		nextArea.stable = true
	}
	return nextArea
}

// iterateTolerant returns the waitingArea after one round of the rules provided in part 2.
// Returns a copy of itself if already stable.
func (area waitingArea) iterateTolerant() waitingArea {
	if area.stable {
		return area.copy()
	}

	nextArea := waitingArea{
		width:  area.width,
		height: area.height,
		stable: area.stable,
	}
	nextArea.representation = make([][]seat, area.height)
	for ii := range nextArea.representation {
		nextArea.representation[ii] = make([]seat, area.width)
	}

	for yy := 0; yy < area.height; yy++ {
		for xx := 0; xx < area.width; xx++ {
			currentSeat := area.at(xx, yy)
			switch currentSeat {
			case floor:
				nextArea.setAt(xx, yy, floor)
			case empty:
				if area.countVisibleOccupied(xx, yy) == 0 {
					nextArea.setAt(xx, yy, occupied)
				} else {
					nextArea.setAt(xx, yy, empty)
				}
			case occupied:
				if area.countVisibleOccupied(xx, yy) >= 5 {
					nextArea.setAt(xx, yy, empty)
				} else {
					nextArea.setAt(xx, yy, occupied)
				}
			}
		}
	}

	if area.equal(nextArea) {
		nextArea.stable = true
	}
	return nextArea
}

// Day11 solves the eleventh day puzzle "Seating System"
//
// Input
//
// A grid representing seat layout on a waiting area. For example:
//
//  some sample input indented to become a code block
//
// It is guaranteed that the rows of the grid are equal in width, and that they initially contain only `L` or `.`.
func Day11(input string) (answer1, answer2 string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	area, err := generateWaitingArea(scanner)
	if err != nil {
		err = errors.Wrapf(err, "could not parse scanner from puzzle input")
		return
	}
	original := area.copy()
	for !area.isStable() {
		area = area.iterateSimple()
	}
	_, countOccupied, _ := area.statistics()
	answer1 = strconv.Itoa(countOccupied)

	// and now for part 2
	area = original
	for !area.isStable() {
		area = area.iterateTolerant()
	}
	_, countOccupied, _ = area.statistics()
	answer2 = strconv.Itoa(countOccupied)

	return
}

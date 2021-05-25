package aoc2016

import (
	"fmt"
	"image"
	"strconv"
	"strings"
)

type direction int

const (
	// B0 indicates NS or EW; B1 indicates positive or negative
	north direction = iota // 0b00
	east                   // 0b01
	south                  // 0b10
	west                   // 0b11
)

// checkPointInSlice: return true if in slice
func checkPointInSlice(slice []image.Point, pt image.Point) bool {
	for _, v := range slice {
		if v == pt {
			return true
		}
	}
	return false
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// movePosition: return image.Point which is old offset by amt in direction
func movePosition(old image.Point, dir direction, amt int) image.Point {
	ns := dir%2 == 0            // true for NS; false for EW
	positive := (dir>>1)%2 == 0 // true for POSITIVE effect; false for NEGATIVE
	new := image.Pt(0, 0)
	if ns {
		if positive {
			new = old.Add(image.Pt(0, amt))
		} else {
			new = old.Add(image.Pt(0, -amt))
		}
	} else {
		if positive {
			new = old.Add(image.Pt(amt, 0))
		} else {
			new = old.Add(image.Pt(-amt, 0))
		}
	}
	return new
}

// Day01 solves the first day puzzle "No Time for a Taxicab".
//
// Input
//
// A single line containing 161 instructions separated by ", ",
// where each instruction describes a turn direction and the distance to travel.
// For example:
//
//	R5, L5, R5, R3
//
// It is guaranteed that the distance to travel per instruction
// is no greater than 200. It is also guaranteed that each
// instruction is prefixed by either "R" or "L".
func Day01(input string) (answer1, answer2 string, err error) {

	currentDirection := north
	current := image.Pt(0, 0) // current position
	allVisited := make([]image.Point, 0, 32)
	found := false // have we found the intersection?
	intersection := image.Pt(0, 0)

	turns := strings.Split(input, ", ")
	for _, turn := range turns {
		if len(turn) < 2 {
			err = fmt.Errorf("invalid string: `%v`", turn)
			return
		}
		if turn[0] == 'R' {
			currentDirection++
		} else if turn[0] == 'L' {
			currentDirection--
		} else {
			err = fmt.Errorf("invalid string: `%v`", turn)
			return
		}
		currentDirection = currentDirection % 4

		amt, e := strconv.Atoi(turn[1:])
		if e != nil {
			err = e
			return
		}

		new := movePosition(current, currentDirection, amt)
		if !found {
			trans := current
			for i := 0; i < amt; i++ {
				trans = movePosition(trans, currentDirection, 1)
				if !checkPointInSlice(allVisited, trans) {
					allVisited = append(allVisited, trans)
				} else {
					found = true
					intersection = trans
				}
			}
		}
		current = new
	}

	answer1 = strconv.Itoa(absInt(current.X) + absInt(current.Y))
	answer2 = strconv.Itoa(absInt(intersection.X) + absInt(intersection.Y))
	return
}

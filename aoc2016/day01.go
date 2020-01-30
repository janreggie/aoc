package aoc2016

import (
	"bufio"
	"fmt"
	"image"
	"strconv"
	"unicode"
	"unicode/utf8"
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

func absInt64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func readUntilCommaSpace(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}
	// Scan until comma or when space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == ',' || unicode.IsSpace(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

// movePosition: return image.Point which is old offset by amt in direction
func movePosition(old image.Point, dir direction, amt int64) image.Point {
	ns := dir%2 == 0            // true for NS; false for EW
	positive := (dir>>1)%2 == 0 // true for POSITIVE effect; false for NEGATIVE
	new := image.Pt(0, 0)
	if ns {
		if positive {
			new = old.Add(image.Pt(0, int(amt)))
		} else {
			new = old.Add(image.Pt(0, int(-amt)))
		}
	} else {
		if positive {
			new = old.Add(image.Pt(int(amt), 0))
		} else {
			new = old.Add(image.Pt(int(-amt), 0))
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
func Day01(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	currentDirection := north
	current := image.Pt(0, 0) // current position
	allVisited := make([]image.Point, 0, 32)
	found := false // have we found the intersection?
	intersection := image.Pt(0, 0)
	scanner.Split(readUntilCommaSpace)
	for scanner.Scan() {
		raw := scanner.Text() // raw == [R|L](count)
		// check direction
		if len(raw) < 2 {
			err = fmt.Errorf("invalid string: `%v`", raw)
			return
		}
		if raw[0] == 'R' {
			currentDirection++
		} else if raw[0] == 'L' {
			currentDirection--
		} else {
			err = fmt.Errorf("invalid string: `%v`", raw)
			return
		}
		currentDirection = currentDirection % 4

		amt, e := strconv.ParseInt(raw[1:], 10, 64)
		if e != nil {
			err = e
			return
		}

		new := movePosition(current, currentDirection, amt)
		if !found {
			trans := current
			for i := int64(0); i < amt; i++ {
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

	answer1 = strconv.FormatInt(absInt64(int64(current.X))+absInt64(int64(current.Y)), 10)
	answer2 = strconv.FormatInt(absInt64(int64(intersection.X))+absInt64(int64(intersection.Y)), 10)
	return
}

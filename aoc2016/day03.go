package aoc2016

import (
	"strconv"
	"strings"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/pkg/errors"
)

// checkSides returns true if three sides of a triangle are valid
func checkSides(input [3]uint64) bool {
	a, b, c := input[0], input[1], input[2]
	sum := func(a, b uint64) uint64 {
		return a + b
	}
	// This may be a bit too verbose...
	return a < sum(b, c) &&
		b < sum(c, a) &&
		c < sum(a, b)
}

// Day03 solves the third day puzzle "Squares With Three Sides".
//
// Input
//
// A file containing several lines each containing three numbers that describe
// a triangle's sides. For example:
//
//	  539  165  487
//	  815  742   73
//	  353  773  428
//	  526  152  680
//	  433  711  557
//	  168  632  306
//
// It is guaranteed that there are no more than 1600 lines
// and that no side is more than three digits long.
func Day03(input string) (answer1, answer2 string, err error) {
	var a1, a2 uint64 // answer1, answer2
	var rowCount uint // number of rows so far (either 0, 1, or 2)

	var columnSides [3][3]uint64 // part two; columnSides[0] is first column's sides

	for _, line := range aoc.SplitLines(input) {
		rawSides := strings.Fields(line)
		if len(rawSides) != 3 {
			err = errors.Errorf("%v doesn't represent a triangle", line)
			return
		}

		var sides [3]uint64
		for ii := range rawSides {
			sides[ii], err = strconv.ParseUint(rawSides[ii], 10, 64)
			if err != nil {
				err = errors.Wrapf(err, "could not parse %v", rawSides[ii])
				return
			}
		}

		// Part 1 logic
		if checkSides(sides) {
			a1++
		}

		// Part 2 logic
		for ii := range sides {
			columnSides[ii][rowCount] = sides[ii]
		}
		rowCount++
		if rowCount == 3 {
			for _, ss := range columnSides {
				if checkSides(ss) {
					a2++
				}
			}
			rowCount = 0
		}
	}

	answer1 = strconv.FormatUint(a1, 10)
	answer2 = strconv.FormatUint(a2, 10)
	return
}

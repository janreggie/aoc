package aoc2019

import (
	"strconv"

	aoc "github.com/janreggie/aoc/internal"
)

// Day01 solves the first day puzzle "The Tyranny of the Rocket Equation".
//
// Input
//
// A file containing 100 lines each representing a number. For example:
//
//	108808
//	87954
//	148957
//	110182
//	126668
//	148024
//	96915
//	117727
//	147378
//
// It is guaranteed that all numbers are no more than six digits long.
func Day01(input string) (answer1, answer2 string, err error) {

	var totalFuel int64 // total fuel required
	var cumuFuel int64  // answer2 (where fuel requires fuel!)
	massToFuel := func(mass int64) int64 {
		return (mass / 3) - 2 // will automatically round down!!
	}
	for _, line := range aoc.SplitLines(input) {
		mass, e := strconv.ParseInt(line, 10, 64)
		if e != nil {
			err = e
			return
		}
		totalFuel += massToFuel(mass)
		for fuel := massToFuel(mass); fuel > 0; fuel = massToFuel(fuel) {
			cumuFuel += fuel
		}
	}

	answer1 = strconv.FormatInt(totalFuel, 10)
	answer2 = strconv.FormatInt(cumuFuel, 10)
	return
}

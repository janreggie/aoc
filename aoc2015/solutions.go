// Package aoc2015 implements the solutions
// for Advent of Code 2015.
package aoc2015

import (
	"errors"
)

// AllSolutions is the slice of all solutions
// so that they may be indexed easily.
// Note that slice index starts at 0 and not 1
// so AllSolutions[0] is the solution for Day 1.
var AllSolutions = []func(string) (string, string, error){
	Day01,         // day 1
	Day02,         // day 2
	Day03,         // day 3
	Day04,         // day 4
	Day05,         // day 5
	Day06,         // day 6
	Day07,         // day 7
	Day08,         // day 8
	Day09,         // day 9
	Day10,         // day 10
	Day11,         // day 11
	Day12,         // day 12
	Day13,         // day 13
	Day14,         // day 14
	Unimplemented, // day 15
	Day16,         // day 16
	Day17,         // day 17
	Unimplemented, // day 18
	Unimplemented, // day 19
	Day20,         // day 20
	Unimplemented, // day 21
	Unimplemented, // day 22
	Unimplemented, // day 23
	Unimplemented, // day 24
	Unimplemented, // day 25
}

// Unimplemented function returns an error
// simply stating unimplemented.
func Unimplemented(string) (string, string, error) {
	return "", "", errors.New("unimplemented")
}

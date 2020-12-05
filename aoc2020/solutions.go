// Package aoc2020 implements the solutions
// for Advent of Code 2020.
package aoc2020

import (
	"bufio"
	"errors"
)

// AllSolutions is the slice of all solutions
// so that they may be indexed easily.
// Note that slice index starts at 0 and not 1
// so AllSolutions[0] is the solution for Day 1.
var AllSolutions = []func(*bufio.Scanner) (string, string, error){
	Day01,         // day 1
	Day02,         // day 2
	Day03,         // day 3
	Day04,         // day 4
	Day05,         // day 5
	Unimplemented, // day 6
	Unimplemented, // day 7
	Unimplemented, // day 8
	Unimplemented, // day 9
	Unimplemented, // day 10
	Unimplemented, // day 11
	Unimplemented, // day 12
	Unimplemented, // day 13
	Unimplemented, // day 14
	Unimplemented, // day 15
	Unimplemented, // day 16
	Unimplemented, // day 17
	Unimplemented, // day 18
	Unimplemented, // day 19
	Unimplemented, // day 20
	Unimplemented, // day 21
	Unimplemented, // day 22
	Unimplemented, // day 23
	Unimplemented, // day 24
	Unimplemented, // day 25
}

// Unimplemented function returns an error
// simply stating unimplemented.
func Unimplemented(scanner *bufio.Scanner) (string, string, error) {
	return "", "", errors.New("unimplemented")
}

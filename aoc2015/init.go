package aoc2015

import "bufio"

// aoc2015 implements the solutions
// for Advent of Code 2015

// AllSolutions is the slice of all solutions
// so that they may be indexed easily.
// Note that slice index starts at 0 and not 1
// so day 1
var AllSolutions = []func(*bufio.Scanner) (string, string, error){
	Day01,   // day 1
	Day02,   // day 2
	Day03,   // day 3
	Day04,   // day 4
	Day05,   // day 5
	Day06,   // day 6
	nothing, // day 7
	nothing, // day 8
	nothing, // day 9
	nothing, // day 10
	nothing, // day 11
	nothing, // day 12
	nothing, // day 13
	nothing, // day 14
	nothing, // day 15
	nothing, // day 16
	nothing, // day 17
	nothing, // day 18
	nothing, // day 19
	nothing, // day 20
	nothing, // day 21
	nothing, // day 22
	nothing, // day 23
	nothing, // day 24
	nothing, // day 25
}

// nothing function is just here
func nothing(scanner *bufio.Scanner) (string, string, error) {
	// return nothing
	return "unimplemented", "unimplemented", nil
}

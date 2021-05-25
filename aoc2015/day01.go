package aoc2015

import (
	"fmt"
	"strconv"
)

// Day01 solves the first day puzzle "Not Quite Lisp".
//
// Input
//
// A single line of length 7000 containing nothing but the characters ')' and '('.
// For example:
//
//	()((()((()(()()())(())()()()((())())))((()()(()()((()(())()()())(((()(()()))))(())))(()(()())()))()()))))))()))))((((
//
// If there exists any character that is neither a '(' or a ')', return an error.
//
func Day01(input string) (answer1, answer2 string, err error) {
	currentFloor := 0
	position := 1
	passedBasement := false

	for _, readFloor := range input {
		if readFloor == '(' {
			currentFloor++
		} else if readFloor == ')' {
			currentFloor--
		} else {
			err = fmt.Errorf("unexpected: %v", readFloor)
			return
		}
		if !passedBasement && currentFloor == -1 {
			answer2 = strconv.Itoa(position)
			passedBasement = true
		}
		position++
	}
	answer1 = strconv.Itoa(currentFloor)
	return
}

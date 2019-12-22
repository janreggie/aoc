package aoc2015

import (
	"bufio"
	"fmt"
	"strconv"
)

// Day01 solves the first day puzzle
// "Not Quite Lisp"
func Day01(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	currentFloor := 0
	position := 1
	passedBasement := false
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		readFloor := scanner.Text() // each character (`(` or `)`)
		if readFloor == "(" {
			currentFloor++
		} else if readFloor == ")" {
			currentFloor--
		} else {
			err = fmt.Errorf("unexpected: %v", readFloor)
			return
		}
		if currentFloor == -1 && !passedBasement {
			answer2 = strconv.Itoa(position)
			passedBasement = true
		}
		position++
	}
	answer1 = strconv.Itoa(currentFloor)
	return
}

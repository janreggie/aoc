package aoc2019

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

// checkAdjacent returns true if there are at least two adjacent numbers.
func checkAdjacent(input int) bool {
	// We can use a recursive solution here,
	// but iterative works too.
	for ; input >= 10; input /= 10 {
		if input%10 == (input/10)%10 {
			return true
		}
	}
	return false
}

func checkDuality(input int) bool {
	// Returns true if there exists a group of
	// adjacent digits of exactly length 2
	// For instance, 111122 fits since 22 exists
	// but 123444 does not fit since 444 is length 3
	// even if it contains 44.

	// let us store the pointer for the stuff
	current, count, next := input%10, 1, 0
	// note input!=input/10 means that
	// if input == 0 then still continue
	for input /= 10; input != input/10; input /= 10 {
		next = input % 10
		if next == current {
			count++
		} else if count == 2 { // hey, a duplicate!
			return true
		} else {
			current = next
			count = 1
		}
	}
	if count == 2 {
		return true
	}
	return false

}

func checkNonDecreasing(input int) bool {
	// Returns true if all digits of input are *non-decreasing*
	for ; input >= 10; input /= 10 {
		if input%10 < (input/10)%10 {
			return false
		}
	}
	return true
}

// Day04 solves the fourth day puzzle
// "Secure Container"
func Day04(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	var current, lowerBound, upperBound int
	var count1, count2 int
	// This looks more like a number theory problem
	// than a programming problem
	// i.e., iterating for every number is too wasteful.
	// Although maybe that's the best solution?
	// This requires further improvement.
	// Note the criteria
	//
	// * It is a six-digit number.
	// * The value is within the range given in your puzzle input.
	// * Two adjacent digits are the same (like 22 in 122345).
	// * Going from left to right, the digits never decrease;
	//   they only ever increase or stay the same (like 111123 or 135679).
	// Let us first get our puzzle input
	scanner.Scan()
	raw := strings.Split(scanner.Text(), "-")
	if len(raw) != 2 {
		err = fmt.Errorf("invalid input: %v", raw)
		return
	}
	lowerBound, err = strconv.Atoi(raw[0])
	if err != nil {
		err = fmt.Errorf("cannot convert %v: %v", raw[0], err)
		return
	}
	upperBound, err = strconv.Atoi(raw[1])
	if err != nil {
		err = fmt.Errorf("cannot convert %v: %v", raw[10], err)
		return
	}

	// How about this?
	for current = lowerBound; current <= upperBound; current++ {
		if checkNonDecreasing(current) && checkAdjacent(current) {
			count1++
			if checkDuality(current) {
				glog.Infof("current: %v", current)
				count2++
			}
		}
	}

	answer1 = strconv.Itoa(count1)
	answer2 = strconv.Itoa(count2)

	return
}

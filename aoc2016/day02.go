package aoc2016

import (
	"strings"
)

// keypadNumber represents a keypad number from 1 to 9 (or 1 to D). (Day 2)
// If keypadNumber is 10, 11, 12, or 13 those represent the letters "A", "B", "C", and "D".
type keypadNumber uint

func (number keypadNumber) Rune() rune {
	if number > 9 {
		return rune(number + 55) // 10 -> 'A', 11 -> 'B', ...
	}
	return rune(number + 48)
}

// iterate moves number using move. move can either be `U`,
// `D`, `L`, or `R`. If not any of them, it will do nothing.
func (number *keypadNumber) iterate(move rune) {
	switch move {
	case 'U':
		if *number >= 4 {
			*number -= 3
		}
	case 'D':
		if *number <= 6 {
			*number += 3
		}
	case 'L':
		if *number%3 != 1 {
			*number--
		}
	case 'R':
		if *number%3 != 0 {
			*number++
		}
	}
}

// iterateDeux iterates number using the Part 2 rules.
func (number *keypadNumber) iterateDeux(move rune) {
	switch move {
	case 'U':
		if *number == 3 || *number == 13 {
			*number -= 2
		} else if (*number >= 6 && *number <= 8) || (*number >= 10 && *number <= 12) {
			*number -= 4
		}
	case 'D':
		if *number == 1 || *number == 11 {
			*number += 2
		} else if (*number >= 2 && *number <= 4) || (*number >= 6 && *number <= 8) {
			*number += 4
		}
	case 'L':
		if *number != 1 && *number != 2 && *number != 5 && *number != 10 && *number != 13 {
			*number--
		}
	case 'R':
		if *number != 1 && *number != 4 && *number != 9 && *number != 12 && *number != 13 {
			*number++
		}
	}
}

// iterateString iterates number using a string of `U`, `D`, `L`, and `R`'s via iterate.
func (number *keypadNumber) iterateString(movement string) {
	for _, move := range movement {
		number.iterate(move)
	}
}

// iterateStringDeux iterates number using a string via iterateDeux.
func (number *keypadNumber) iterateStringDeux(movement string) {
	for _, move := range movement {
		number.iterateDeux(move)
	}
}

// Day02 solves the second day puzzle "Bathroom Security".
//
// Input
//
// A file of some number of lines, each containing variable length and
// only the characters `U`, `D`, `L`, and `R`. For example:
//
//	ULL
//	RRDDD
//	LURDL
//	UUUUD
//
// It is guaranteed that no line exceeds 600 characters,
// and there are no more than five lines.
func Day02(input string) (answer1, answer2 string, err error) {

	var initOne, initTwo keypadNumber = 5, 5

	var a1, a2 strings.Builder
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		initOne.iterateString(line)
		initTwo.iterateStringDeux(line)
		a1.WriteRune(initOne.Rune())
		a2.WriteRune(initTwo.Rune())
	}

	answer1 = a1.String()
	answer2 = a2.String()

	return
}

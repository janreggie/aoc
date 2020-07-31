package aoc2016

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// decompress implements the decompression algorithm in Y2016 D09
func decompress(input string) (string, error) {
	// initialize the Scanner's state
	var sb strings.Builder // writing output string
	type state int
	const (
		Normal      state = iota // just output the same text
		CountChars               // parsing the first number (before `x`)
		CountRepeat              // parsing the second number
		GetToRepeat              // parsing the characters to be repeated
	)
	currentState := Normal
	var charCount, repeatCount int // number of characters to parse and number of times to repeat
	var charCountRaw, repeatCountRaw strings.Builder
	var repeatingBuffer strings.Builder
	input = strings.Join(strings.Fields(input), "") // remove all whitespace

	for _, character := range input {
		switch currentState {
		case Normal:
			if character != '(' {
				sb.WriteRune(character)
				continue
			}
			// otherwise...
			currentState = CountChars
			charCount, repeatCount = 0, 0

		case CountChars:
			if character != 'x' {
				charCountRaw.WriteRune(character)
				continue
			}
			// otherwise...
			currentState = CountRepeat
			var err error
			charCount, err = strconv.Atoi(charCountRaw.String())
			if err != nil {
				err := errors.Wrapf(err, "could not parse %v in CountChars", charCountRaw.String())
				return sb.String(), err
			}
			charCountRaw.Reset()

		case CountRepeat:
			if character != ')' {
				repeatCountRaw.WriteRune(character)
				continue
			}
			// otherwise...
			currentState = GetToRepeat
			var err error
			repeatCount, err = strconv.Atoi(repeatCountRaw.String())
			if err != nil {
				err := errors.Wrapf(err, "could not parse %v in CountRepeat", repeatCountRaw.String())
				return sb.String(), err
			}
			repeatCountRaw.Reset()

		case GetToRepeat:
			// get the next charCount characters (unless if there are no longer any)
			if charCount > 0 {
				repeatingBuffer.WriteRune(character)
				charCount--
				continue
			}
			// but once it hits zero
			currentState = Normal
			for ii := 0; ii < repeatCount; ii++ {
				sb.WriteString(repeatingBuffer.String())
			}
			repeatingBuffer.Reset()
			// don't forget to write the remaining character
			sb.WriteRune(character)

		default:
			err := fmt.Errorf("unknown state: %v", currentState)
			return sb.String(), err
		}
	}

	// in case if repeatingBuffer hasn't been exhausted
	if repeatingBuffer.Len() != 0 {
		for ii := 0; ii < repeatCount; ii++ {
			sb.WriteString(repeatingBuffer.String())
		}
	}

	return sb.String(), nil
}

// Day09 solves the ninth day puzzle "Explosives in Cyberspace".
//
// Input
//
// A single line, containing a sequence of bytes representing the experimental format.
// It is assumed that the data is valid, and it does not contain more than 15,000 characters.
func Day09(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	var sb strings.Builder
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}
	input := sb.String()

	fmt.Println("Compressed: ", input)
	decompressed, err := decompress(input)
	if err != nil {
		return
	}

	fmt.Println("Decompressed: ", decompressed)
	answer1 = strconv.Itoa(len(decompressed))
	return
}

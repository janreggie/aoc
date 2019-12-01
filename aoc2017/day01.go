package aoc2017

import (
	"bufio"
	"strconv"
)

// Day01 solves the first day puzzle
// "Inverse Captcha"
func Day01(scanner *bufio.Scanner) (answer1, answer2 string, e error) {
	// scanner.Split(bufio.ScanBytes)
	var input string
	var adjacentCaptcha, halfwayCaptcha int64 // answer1, answer2
	for scanner.Scan() {
		input = scanner.Text() // string should not be THAT long
	}

	// now two for loops
	// first one for adjacents
	for ii, length := 0, len(input); ii < length; ii++ {
		// check if input[ii] equals input[ii+1] (except when ii == length-1)
		if input[ii] == input[(ii+1)%length] {
			toAdd, err := strconv.ParseInt(string(input[ii]), 10, 64)
			if err != nil {
				e = err
				return
			}
			adjacentCaptcha += toAdd
		}
	}

	// second one for halfways
	for ii, halflength := 0, len(input)/2; ii < halflength; ii++ {
		if input[ii] == input[ii+halflength] {
			toAdd, err := strconv.ParseInt(string(input[ii]), 10, 64)
			if err != nil {
				e = err
				return
			}
			halfwayCaptcha += toAdd * 2
		}
	}

	answer1 = strconv.FormatInt(adjacentCaptcha, 10)
	answer2 = strconv.FormatInt(halfwayCaptcha, 10)
	return
}

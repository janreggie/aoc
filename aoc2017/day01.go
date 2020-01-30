package aoc2017

import (
	"bufio"
	"strconv"
	"sync"

	"github.com/pkg/errors"
)

// Day01 solves the first day puzzle "Inverse Captcha".
//
// Input
//
// A file containing a single line of 2000 characters consisting only of numerals.
// For example:
//
//	91212129
//
// It is guaranteed that the first line does not contain any non-numeric characters.
func Day01(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	// scanner.Split(bufio.ScanBytes)
	var adjacentCaptcha, halfwayCaptcha int64 // answer1, answer2
	if !scanner.Scan() {
		err = errors.Wrap(scanner.Err(), "could not read from scanner")
		return
	}
	input := scanner.Text() // string should not be THAT long

	// now two for loops
	var wg sync.WaitGroup
	wg.Add(2)
	// first one for adjacents
	go func() {
		for ii, length := 0, len(input); ii < length; ii++ {
			// check if input[ii] equals input[ii+1] (except when ii == length-1)
			if input[ii] == input[(ii+1)%length] {
				toAdd, e := strconv.ParseInt(string(input[ii]), 10, 64)
				if e != nil {
					err = e
					return
				}
				adjacentCaptcha += toAdd
			}
		}
		answer1 = strconv.FormatInt(adjacentCaptcha, 10)
		wg.Done()
	}()

	// second one for halfways
	go func() {
		for ii, halflength := 0, len(input)/2; ii < halflength; ii++ {
			if input[ii] == input[ii+halflength] {
				toAdd, e := strconv.ParseInt(string(input[ii]), 10, 64)
				if e != nil {
					err = e
					return
				}
				halfwayCaptcha += toAdd * 2
			}
		}
		answer2 = strconv.FormatInt(halfwayCaptcha, 10)
		wg.Done()
	}()

	wg.Wait()
	return
}

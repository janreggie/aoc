package aoc2018

import (
	"strconv"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/pkg/errors"
)

// Day01 solves the first day puzzle "Chronal Calibration".
//
// Input
//
// A file containing 958 lines, each of which is either the character `+`
// or `-` followed by a number. For example:
//
//	-19
//	-18
//	+13
//	+14
//	+7
//	-3
//	+14
//	+5
//	-11
//	-13
//	-22
//	+4
//	+12
//
// It is guaranteeed that these numbers will be no more than 100000
// and will be no less than -100000.
func Day01(input string) (answer1, answer2 string, err error) {

	allInstr := []int64{} // all instructions for safekeeping
	allFreqs := make(map[int64]struct{})
	allFreqs[0] = struct{}{}
	currentFreq := int64(0)

	for _, line := range aoc.SplitLines(input) {
		toAdd, e := strconv.ParseInt(line, 10, 64)
		if e != nil {
			err = errors.Wrapf(e, "could not parse %s", line)
			return
		}
		allInstr = append(allInstr, toAdd)
		currentFreq += toAdd
		if answer2 == "" {
			if _, found := allFreqs[currentFreq]; found {
				answer2 = strconv.FormatInt(currentFreq, 10)
			}
		}
		allFreqs[currentFreq] = struct{}{}
	}
	answer1 = strconv.FormatInt(currentFreq, 10)

	findRepeatingInd := 0
	for answer2 == "" {
		currentFreq += allInstr[findRepeatingInd]
		if _, found := allFreqs[currentFreq]; found {
			answer2 = strconv.FormatInt(currentFreq, 10)
			break
		}
		allFreqs[currentFreq] = struct{}{}

		findRepeatingInd++
		if findRepeatingInd >= len(allInstr) {
			findRepeatingInd = 0
		}
	}

	return
}

package aoc2018

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/janreggie/AdventOfCode/tools"
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
// It is guaranteeed that the absolute values of these numbers
// will be no more than 100000.
func Day01(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	scanner.Split(bufio.ScanWords)
	var frequency int64   // both actually
	var onceFreq int64    // answer1
	var repdFreq int64    // answer2
	hasRepeated := false  // answer2 (has repeated?)
	allInstr := []int64{} // answer2 (in case hasn't repeated)
	allFreqs := tools.NewBinaryTree(0)
	repeatCycle := func() {
		// will repeat until when hasRepeated has become false or something's found
		if hasRepeated == false && allFreqs.Update(frequency) == true {
			// do nothing really
		} else {
			repdFreq = frequency
			hasRepeated = true // stop the cycle!
		}
	}
	for scanner.Scan() {
		raw := scanner.Text()
		toAdd, e := strconv.ParseInt(raw, 10, 64)
		if e != nil {
			err = fmt.Errorf("%v is invalid: %v", raw, e)
			return
		}
		frequency += toAdd
		// check if frequency is in the thingy
		allInstr = append(allInstr, toAdd) // append for later
		repeatCycle()
	}
	onceFreq = frequency
	for instr, length := 0, len(allInstr); !hasRepeated; instr = (instr + 1) % length { // while it has not repeated
		// allInstr[instr] shall be run
		frequency += allInstr[instr]
		repeatCycle()
	}

	answer1 = strconv.FormatInt(onceFreq, 10)
	answer2 = strconv.FormatInt(repdFreq, 10)
	return
}

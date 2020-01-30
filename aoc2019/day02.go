package aoc2019

import (
	"bufio"
	"strconv"

	"github.com/janreggie/AdventOfCode/tools/intcode"
)

// Day02 solves the second day puzzle
// "1202 Program Alarm"
func Day02(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	var ic *intcode.IntCode
	ic, err = intcode.NewFromScanner(scanner)
	if err != nil {
		return
	}
	ic.Install(intcode.SimpleAdder)
	ic.Install(intcode.SimpleMultiplier)
	instrs := ic.Snapshot() // for later
	ic.SetLocation(1, 12)
	ic.SetLocation(2, 2)
	if errOpn := ic.Operate(); !intcode.IsHalt(errOpn) {
		err = errOpn
		return
	}
	first, err := ic.GetLocation(0)
	if err != nil {
		return
	}
	answer1 = strconv.FormatInt(first, 10)

	// now for part 2
	// There are a total of 10000 solutions
	// where memory locations 1 and 2 can be between 0 to 99.
	// I am thinking of using a more... algebraic solution
	// but that doesn't seem to be possible.

bruteforce:
	for mem1 := int64(0); mem1 < 100; mem1++ {
		for mem2 := int64(0); mem2 < 100; mem2++ {
			ic.Format(instrs)
			if err = ic.SetLocation(1, mem1); err != nil {
				return
			}
			if err = ic.SetLocation(2, mem2); err != nil {
				return
			}
			if errOpn := ic.Operate(); !intcode.IsHalt(errOpn) {
				err = errOpn
				return
			}
			if val, e := ic.GetLocation(0); val == 19690720 && e == nil {
				answer2 = strconv.FormatInt(mem1*100+mem2, 10)
				break bruteforce
			}
		}
	}
	return
}

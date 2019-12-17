package aoc2019

import (
	"bufio"
	"strconv"

	"github.com/janreggie/AdventOfCode/structs/intcode"
)

// Day05 solves the fifth day puzzle
// "Sunny with a Chance of Asteroids"
func Day05(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	var ic *intcode.IntCode
	if ic, err = intcode.NewFromScanner(scanner); err != nil {
		return
	}
	ic.Install(intcode.Inputter)
	ic.Install(intcode.OutputAndHalt)
	ic.Install(intcode.Adder)
	ic.Install(intcode.Multiplier)
	mem := ic.Snapshot()

	// Part 1
	ic.SetInput(1)
	if err = ic.Operate(); !intcode.IsHalt(err) {
		return
	}
	err = nil
	// the last element of Output is our answer1
	answer1 = strconv.Itoa(ic.Output()[len(ic.Output())-1])

	// Part 2
	ic.Format(mem)
	ic.SetInput(5)
	ic.Install(intcode.JumpIfFalse)
	ic.Install(intcode.JumpIfTrue)
	ic.Install(intcode.LessThan)
	ic.Install(intcode.Equals)
	if err = ic.Operate(); !intcode.IsHalt(err) {
		return
	}
	err = nil
	answer2 = strconv.Itoa(ic.Output()[len(ic.Output())-1])

	return
}

package aoc2019

import (
	"bufio"
	"strconv"

	"github.com/janreggie/AdventOfCode/aoc2019/intcode"
)

// Day05 solves the fifth day puzzle "Sunny with a Chance of Asteroids".
//
// Input
//
// A single line containing several integers (that could be negative) separated by commas,
// representing an Intcode program. For example:
//
//	3,3,1105,-1,9,1101,0,0,12,4,12,99,1
//
// It is assumed that the numbers do not exceed five digits long,
// and that the tape length is no more than 700.
//
// The Intcode computer should now be able to support the following Opcodes:
// INPUT (3), OUTPUT (4), JUMP-IF-TRUE (5), JUMP-IF-FALSE (6), LESSTHAN (7),
// and EQUALS (8). In addition, it should be able to support parameter modes.
//
// The first Opcode of the Intcode program should be an INPUT (3).
func Day05(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	var ic *intcode.Intcode
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
	answer1 = strconv.FormatInt(ic.Output()[len(ic.Output())-1], 10)

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
	answer2 = strconv.FormatInt(ic.Output()[len(ic.Output())-1], 10)

	return
}

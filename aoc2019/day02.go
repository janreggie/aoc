package aoc2019

import (
	"bufio"
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

// adopted in aoc2016
func readUntilCommaSpace(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}
	// Scan until comma or when space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == ',' || unicode.IsSpace(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

// day02Run runs a tape (slice) of Intcode
func day02Run(intcode []int64) error {
	pc := 0 // program counter in intcode
	for intcode[pc] != 99 {
		// do things to the computer
		// make sure pc+3 < len(intcode)
		if pc+3 >= len(intcode) {
			return fmt.Errorf("invalid pc: %v", pc)
		}
		op1, op2, op3 := intcode[pc+1], intcode[pc+2], intcode[pc+3]
		var opR int64
		// assert first that op1, op2 < len(intcode)
		if int(op1) >= len(intcode) || int(op2) >= len(intcode) || int(op3) >= len(intcode) {
			return fmt.Errorf("invalid operands %v and %v (len(intcode): %v)", op1, op2, len(intcode))
		}
		if intcode[pc] == 1 {
			// access intcode[pc+1] and intcode[pc+2]
			opR = intcode[op1] + intcode[op2]
		} else if intcode[pc] == 2 {
			opR = intcode[op1] * intcode[op2]
		} else {
			return fmt.Errorf("invalid opcode %v at position %v", intcode[pc], pc)
		}
		intcode[op3] = opR

		pc += 4
	}
	return nil
}

// Day02 solves the second day puzzle
// "1202 Program Alarm"
func Day02(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	scanner.Split(readUntilCommaSpace)
	allInstr := make([]int64, 0, 12) // i dunno how long it can be
	for scanner.Scan() {
		raw := scanner.Text()
		instr, e := strconv.ParseInt(raw, 10, 64)
		if e != nil {
			err = fmt.Errorf("cannot parse %v: %v", raw, e)
			return
		}
		allInstr = append(allInstr, instr)
	}
	backInstr := make([]int64, len(allInstr))
	for ii := range allInstr {
		backInstr[ii] = allInstr[ii]
	}
	// okay all instructions are in allInstr
	// but let's change some values first!
	if len(allInstr) < 3 {
		err = fmt.Errorf("allInstr is length %v: at least 3 required", len(allInstr))
		return
	}
	allInstr[1] = 12
	allInstr[2] = 2

	day02Run(allInstr)
	answer1 = strconv.FormatInt(allInstr[0], 10)

	// now for part 2
	// There are a total of 10000 solutions
	// where memory locations 1 and 2 can be between 0 to 99.
	// I am thinking of using a more... algebraic solution
	// but that doesn't seem to be possible.
bruteforce:
	for mem1 := int64(0); mem1 < 100; mem1++ {
		for mem2 := int64(0); mem2 < 100; mem2++ {
			for ii := range allInstr {
				allInstr[ii] = backInstr[ii]
			}
			allInstr[1], allInstr[2] = mem1, mem2
			day02Run(allInstr)
			if allInstr[0] == 19690720 {
				answer2 = strconv.FormatInt(allInstr[1]*100+allInstr[2], 10)
				break bruteforce
			}
		}
	}

	return
}

package aoc2020

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// instruction contains an operation (string) and an argument (integer)
type instruction struct {
	op  string // `acc`, `jmp`, or `nop`
	arg int    // any number following op
}

func (instr instruction) String() string {
	return fmt.Sprintf("%s %+v", instr.op, instr.arg)
}

// readInstruction reads a string representing an instruction (e.g., `acc +1`)
func readInstruction(input string) (instruction, error) {
	result := instruction{}
	_, err := fmt.Sscanf(input, "%s %d", &result.op, &result.arg)
	if err != nil {
		return result, errors.Wrapf(err, "could not read instruction input `%s`", input)
	}
	return result, nil
}

// generateInstructionList generates a list of instructions (and an applicable error) from a bufio.Scanner
func generateInstructionList(scanner *bufio.Scanner) ([]instruction, error) {
	result := make([]instruction, 0)
	for scanner.Scan() {
		instr, err := readInstruction(scanner.Text())
		if err != nil {
			return nil, errors.Wrapf(err, "could not generate list of instructions")
		}
		result = append(result, instr)
	}
	return result, nil
}

// computer represents a basic "computer" that can take in a list of instructions
type computer struct {
	instructions []instruction
	accumulator  int
	instrPointer int   // at which instruction are we
	history      []int // history[i] indicates how many times instructions[i] has been executed
}

// newComputer generates a computer from a list of instructions
func newComputer(instrs []instruction) computer {
	result := computer{}
	result.instructions = make([]instruction, len(instrs))
	copy(result.instructions, instrs)
	result.accumulator = 0
	result.instrPointer = 0
	result.history = make([]int, len(instrs))
	return result
}

func (cc computer) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Instruction count: %d\n", len(cc.instructions))
	fmt.Fprintf(&sb, "Current instruction: %d (%v)\n", cc.instrPointer, cc.instructions[cc.instrPointer])
	fmt.Fprintf(&sb, "Accumulator value: %d\n", cc.accumulator)
	return sb.String()
}

// currentHistory returns how many times the instruction at instrPointer has been executed
func (cc computer) currentHistory() int {
	if cc.instrPointer < 0 || cc.instrPointer >= len(cc.instructions) {
		return 0
	}
	return cc.history[cc.instrPointer]
}

// acc returns the value in the accumulator
func (cc computer) acc() int { return cc.accumulator }

// execute executes the current instruction.
// If cc.instrPointer is out of bounds, return some error.
// If some instruction is invalid, return an error. Otherwise, return nil.
func (cc *computer) execute() error {
	if cc.instrPointer < 0 || cc.instrPointer >= len(cc.instructions) {
		return oobInstrPointerError{ptr: cc.instrPointer, len: len(cc.instructions)}
	}
	instr := cc.instructions[cc.instrPointer]
	cc.history[cc.instrPointer]++

	switch instr.op {
	case "nop":
		cc.instrPointer++
		return nil
	case "acc":
		cc.accumulator += instr.arg
		cc.instrPointer++
		return nil
	case "jmp":
		cc.instrPointer += instr.arg
		return nil
	default:
		return errors.Errorf("could not parse instruction %v", instr)
	}
}

// reset resets the computer's variables, including history
func (cc *computer) reset() {
	cc.accumulator = 0
	cc.instrPointer = 0
	cc.history = make([]int, len(cc.instructions))
}

// replaceInstruction replaces the instruction on a certain index with something else.
// If index is out of bounds, do nothing.
func (cc *computer) replaceInstruction(ind int, instr instruction) {
	if ind < 0 || ind >= len(cc.instructions) {
		return
	}
	cc.instructions[ind] = instr
}

// oobInstrPointerError is returned when the instructionPointer of a computer is out-of-bounds
type oobInstrPointerError struct {
	ptr int // instructionPointer value
	len int // length of instructions
}

// Error implements the error interface
func (err oobInstrPointerError) Error() string {
	return fmt.Sprintf("pointer at %d but length of instructions is %d", err.ptr, err.len)
}

// Day08 solves the eighth day puzzle "Handheld Halting"
//
// Input
//
// A list of instructions, each containing an operation and an argument. For example:
//
//   nop +0
//   acc +1
//   jmp +4
//   acc +3
//   jmp -3
//   acc -99
//   acc +1
//   jmp -4
//   acc +6
//
// It is guaranteed that the operations are either one of `acc`, `jmp`, or `nop`,
// and that all arguments are integers.
func Day08(input string) (answer1, answer2 string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	instructions, err := generateInstructionList(scanner)
	if err != nil {
		err = errors.Wrapf(err, "could not read from scanner")
		return
	}

	computer := newComputer(instructions)
	for computer.currentHistory() == 0 {
		computer.execute()
	}
	answer1 = strconv.Itoa(computer.acc())

	// Now, for the second part, this will be a bit more interesting...
	computer.reset()
	// each instruction (that isn't acc) may be erroneous, so change it
PARENT:
	for ii, instr := range instructions {
		if instr.op == "acc" {
			continue
		}
		computer.reset()
		newInstr := instruction{arg: instr.arg} // the new instruction to be replaced
		if instr.op == "nop" {
			newInstr.op = "jmp"
		} else {
			newInstr.op = "nop"
		}
		computer.replaceInstruction(ii, newInstr)
		for computer.currentHistory() == 0 {
			err := computer.execute()
			if errors.As(err, &oobInstrPointerError{}) {
				break PARENT
			}
		}
		computer.replaceInstruction(ii, instr)
	}
	answer2 = strconv.Itoa(computer.acc())
	return
}

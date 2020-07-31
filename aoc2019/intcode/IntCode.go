// Package intcode implements programming in "Intcode"
// for Advent of Code 2019.
package intcode

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

// Halt is a module that is built in to the Intcode
var Halt *Module = NewModule(ModuleConfig{
	Opcode:        99,
	Mnemonic:      "HALT",
	Parameterized: false,
	Function: func(ic *Intcode) error {
		return NewHaltError("HALT (99)") // that's its literally function
	},
})

// Intcode implements an "Intcode" computer
// consisting of a program counter and a tape of memory
// as well as a list of "modules"
// which are functions that take in an Intcode
// and may return an error.
type Intcode struct {
	pc           int64     // program counter
	mem          []int64   // memory
	modules      []*Module // all modules it has
	input        []int64   // an input "stack" (FILO)
	output       []int64   // an output "stack"
	relativeBase int64     // the "relative base" of the computer
}

// New generates an Intcode using a memory reel
func New(mem []int64) (ic *Intcode) {
	ic = &Intcode{pc: 0, mem: make([]int64, len(mem))}
	copy(ic.mem, mem)
	ic.Install(Halt)
	return
}

// NewFromScanner takes an Intcode from a *bufio.Scanner.
// Originally adapted from aoc2019.Day02
func NewFromScanner(scanner *bufio.Scanner) (ic *Intcode, err error) {
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
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
	})
	// now let's start scanning
	ic = &Intcode{pc: 0, mem: make([]int64, 0)}
	for scanner.Scan() {
		raw := scanner.Text()
		instr, e := strconv.ParseInt(raw, 10, 64)
		if e != nil {
			err = fmt.Errorf("cannot parse %v: %v", raw, err)
			return
		}
		ic.mem = append(ic.mem, instr)
	}
	ic.Install(Halt)
	return
}

// Install installs a module
func (ic *Intcode) Install(module *Module) {
	ic.modules = append(ic.modules, module)
	return
}

// UninstallAll removes all modules
func (ic *Intcode) UninstallAll() {
	ic.modules = make([]*Module, 0)
	return
}

// Operate performs instructions on the Intcode computer
// depending on the modules it has
func (ic *Intcode) Operate() (err error) {
	// run through the modules
	for ic.pc < int64(len(ic.mem)) { // while we haven't reached the end yet
		// let's go through all the modules
		for _, module := range ic.modules {
			opcode := ic.Current()
			if module.parameterized {
				opcode = opcode % 100 // we only care about the last two
			}
			if module.opcode != opcode {
				err = NewInvalidOpcodeError(opcode, ic.pc, module)
				continue
			}
			// but hey it's equal now!!
			err = module.function(ic)
			break // out of the loop once we found a module
		}
		if err != nil {
			if IsHalt(err) {
				return err // I mean... it's a halt. that's okay.
			}
			if _, isInvalid := err.(*InvalidOpcodeError); isInvalid {
				err = NewOperationError(errors.New("could not find a suitable opcode"), ic)
				return
			}
			err = NewOperationError(err, ic)
			return // either something happened in module.function or it cannot find a module
		}
	}
	return
}

// Snapshot returns a copy of its memory
func (ic *Intcode) Snapshot() (mem []int64) {
	mem = make([]int64, ic.Len())
	for i := range ic.mem {
		mem[i] = ic.mem[i]
	}
	return
}

// allocateMore allocates number amount of memory locations for ic.mem
func (ic *Intcode) allocateMore(amount int64) {
	// make sure amount > 0
	if amount < 0 {
		return
	}
	ic.mem = append(ic.mem, make([]int64, amount)...)
}

// GetLocation returns the value of the memory at a particular location.
// If location is more than the memory length,
// ic.mem is reallocated.
// If location is negative it will simply return an error.
func (ic *Intcode) GetLocation(location int64) (value int64, err error) {
	if location < 0 {
		err = NewOutOfBoundsError(location, ic.Len())
	}
	if location >= ic.Len() {
		ic.allocateMore(location - ic.Len() + 1)
	}
	value = ic.mem[location]
	return
}

// GetNext returns a fragment of memory after Current()
// containing the next count locations
func (ic *Intcode) GetNext(count int64) (mem []int64, err error) {
	mem = make([]int64, count)
	for ii := int64(0); ii < count; ii++ {
		if mem[ii], err = ic.GetLocation(ic.pc + ii + 1); err != nil {
			return
		}
	}
	return
}

// SetLocation sets the value of the memory at some location.
// If location is more than the memory length,
// ic.mem is reallocated.
// If location is negative it will simply return an error.
func (ic *Intcode) SetLocation(location, value int64) (err error) {
	if location < 0 {
		err = NewOutOfBoundsError(location, ic.Len())
		return
	}
	if location >= ic.Len() {
		ic.allocateMore(location - ic.Len() + 1)
	}
	ic.mem[location] = value
	return
}

// PC returns the current value for the program counter
func (ic *Intcode) PC() (pc int64) {
	pc = ic.pc
	return
}

// Current returns the current memory location at the program counter
func (ic *Intcode) Current() (value int64) {
	return ic.mem[ic.pc]
}

// Len returns the length of the memory
func (ic *Intcode) Len() (length int64) {
	return int64(len(ic.mem))
}

// Increment increments the program counter by a set amount
func (ic *Intcode) Increment(value int64) (err error) {
	if (value+ic.pc) > ic.Len() || (value+ic.pc) < 0 {
		// greater than since we can assume pc equals Len()
		return NewOutOfBoundsError(value+ic.pc, ic.Len())
	}
	ic.pc += value
	return
}

// Jump jumps the program counter to some value
func (ic *Intcode) Jump(value int64) (err error) {
	if value > ic.Len() || value < 0 {
		// greater than since we can assume pc equals Len() (where it will just halt)
		return NewOutOfBoundsError(value, ic.Len())
	}
	ic.pc = value
	return
}

// RelativeBase returns the relative base of the ic computer
func (ic *Intcode) RelativeBase() (relativeBase int64) {
	relativeBase = ic.relativeBase
	return
}

// AdjustRelativeBase adjusts the relative base by some amount,
// increasing or decreasing it.
func (ic *Intcode) AdjustRelativeBase(amount int64) {
	ic.relativeBase += amount
	return
}

// SetRelativeBase sets the relative base by some amount.
func (ic *Intcode) SetRelativeBase(amount int64) {
	ic.relativeBase = amount
	return
}

// Rewind jumps PC to zero
func (ic *Intcode) Rewind() {
	ic.Jump(0) // this cannot return any error
	return
}

// Format formats the memory, input, and outputs and sets PC and relative base to zero
// but does not remove installed modules
func (ic *Intcode) Format(mem []int64) {
	ic.mem = make([]int64, len(mem))
	copy(ic.mem, mem)
	ic.SetInput()
	ic.ResetOutput()
	ic.SetRelativeBase(0)
	ic.Rewind()
	return
}

// Input returns a copy of its inputs
func (ic *Intcode) Input() (input []int64) {
	input = make([]int64, len(ic.input))
	for ii := range input {
		input[ii] = ic.input[ii]
	}
	return
}

// PushToInput pushes a value to the input queue
func (ic *Intcode) PushToInput(input int64) {
	ic.input = append(ic.input, input)
	return
}

// SetInput sets the input
func (ic *Intcode) SetInput(inputs ...int64) {
	ic.input = make([]int64, len(inputs))
	for ii := range inputs {
		ic.input[ii] = inputs[ii]
	}
}

// GetInput removes an input from the queue
func (ic *Intcode) GetInput() (input int64, err error) {
	if len(ic.input) == 0 {
		err = fmt.Errorf("input is length zero")
		return
	}
	input = ic.input[len(ic.input)-1]
	ic.input = ic.input[:len(ic.input)-1]
	return
}

// ResetOutput resets the outputs
func (ic *Intcode) ResetOutput() {
	ic.output = make([]int64, 0)
}

// PushToOutput pushes a value to its outputs
func (ic *Intcode) PushToOutput(value int64) {
	ic.output = append(ic.output, value)
}

// Output prints the output
func (ic *Intcode) Output() (output []int64) {
	output = make([]int64, len(ic.output))
	for ii := range output {
		output[ii] = ic.output[ii]
	}
	return
}

// GetOutput removes an output from the stack
func (ic *Intcode) GetOutput() (output int64, err error) {
	if len(ic.output) == 0 {
		err = fmt.Errorf("output is length zero")
		return
	}
	output = ic.output[len(ic.output)-1]
	ic.output = ic.output[:len(ic.output)-1]
	return
}

// Package intcode implements programming in "IntCode"
// for Advent of Code 2019.
package intcode

import (
	"bufio"
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

// Module is a module with an opcode and a ParamCount
// which does something to an ic computer *IntCode using the next ParamCount memory locations.
// Calling function can return an error if its params turns out to be invalid
// e.g., accessing an invalid memory address.
//
// It is assumed that calling function only happens if ic.Current() equals the opcode,
// unless if the Module supports "parameter modes",
// where in that case ic.Current()%100 is checked instead.
// Note that function will affect the IntCode computer
// e.g., changing its memory, inputs and outputs.
type Module struct {
	opcode        int                     // opcode (if 0 then check will occur in function)
	mnemonic      string                  // "name" of the opcode
	parameterized bool                    // should module support parameter modes?
	function      func(ic *IntCode) error // what does it do to the computer?
}

// NewModule generates a module object with several attributes using a config struct
func NewModule(config struct {
	Opcode        int                     // opcode (if 0 then check will occur in function)
	Mnemonic      string                  // "name" of the opcode
	Parameterized bool                    // should module support parameter modes?
	Function      func(ic *IntCode) error // what does it do to the computer?
}) *Module {
	return &Module{
		opcode:        config.Opcode,
		mnemonic:      config.Mnemonic,
		parameterized: config.Parameterized,
		function:      config.Function,
	}
}

// Halt is a module that is built in to the IntCode
var Halt *Module = &Module{
	// ToRun:    func(opcode int) bool { return opcode == 99 },
	opcode:        99,
	mnemonic:      "HALT",
	parameterized: false,
	function: func(ic *IntCode) error {
		return NewHaltError("HALT (99)") // that's its literally function
	},
}

// IntCode implements an "IntCode" computer
// consisting of a program counter and a tape of memory
// as well as a list of "modules"
// which are functions that take in an IntCode
// and may return an error.
type IntCode struct {
	pc      int       // program counter
	mem     []int     // memory
	modules []*Module // all modules it has
	input   []int     // an input "stack" (FILO)
	output  []int     // an output "stack"
}

// New generates an IntCode using a memory reel
func New(mem []int) (ic *IntCode) {
	ic = &IntCode{pc: 0, mem: make([]int, len(mem))}
	copy(ic.mem, mem)
	ic.Install(Halt)
	return
}

// NewFromScanner takes an IntCode from a *bufio.Scanner.
// Originally adapted from aoc2019.Day02
func NewFromScanner(scanner *bufio.Scanner) (ic *IntCode, err error) {
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
	ic = &IntCode{pc: 0, mem: make([]int, 0)}
	for scanner.Scan() {
		raw := scanner.Text()
		instr, e := strconv.Atoi(raw)
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
func (ic *IntCode) Install(module *Module) {
	ic.modules = append(ic.modules, module)
	return
}

// UninstallAll removes all modules
func (ic *IntCode) UninstallAll() {
	ic.modules = make([]*Module, 0)
	return
}

// Operate performs instructions on the IntCode computer
// depending on the modules it has
func (ic *IntCode) Operate() (err error) {
	// run through the modules
	for ic.pc < len(ic.mem) { // while we haven't reached the end yet
		// let's go through all the modules
		for _, module := range ic.modules {
			opcode := ic.Current()
			if module.parameterized {
				opcode = opcode % 100 // we only care about the last two
			}
			if module.opcode != opcode {
				err = NewInvalidOpcodeError(opcode, module)
				continue
			}
			// but hey it's equal now!!
			err = module.function(ic)
			break // out of the loop once we found a module
		}
		if err != nil {
			return // either something happened in module.function or it cannot find a module
		}
	}
	return
}

// Snapshot returns a copy of its memory
func (ic *IntCode) Snapshot() (mem []int) {
	mem = make([]int, ic.Len())
	for i := range ic.mem {
		mem[i] = ic.mem[i]
	}
	return
}

// GetLocation returns the value of the memory at a particular location
func (ic *IntCode) GetLocation(location int) (value int, err error) {
	if location >= ic.Len() || location < 0 {
		err = NewOutOfBoundsError(location, ic.Len())
	}
	value = ic.mem[location]
	return
}

// GetNext returns a fragment of memory after Current()
// containing the next count locations
func (ic *IntCode) GetNext(count int) (mem []int, err error) {
	mem = make([]int, count)
	for ii := 0; ii < count; ii++ {
		if mem[ii], err = ic.GetLocation(ic.pc + ii + 1); err != nil {
			return
		}
	}
	return
}

// SetLocation sets the value of the memory at some location
func (ic *IntCode) SetLocation(location, value int) (err error) {
	if location >= ic.Len() || location < 0 {
		return NewOutOfBoundsError(value, ic.Len())
	}
	ic.mem[location] = value
	return
}

// PC returns the current value for the program counter
func (ic *IntCode) PC() (pc int) {
	pc = ic.pc
	return
}

// Current returns the current memory location at the program counter
func (ic *IntCode) Current() (value int) {
	return ic.mem[ic.pc]
}

// Len returns the length of the memory
func (ic *IntCode) Len() (length int) {
	return len(ic.mem)
}

// Increment increments the program counter by a set amount
func (ic *IntCode) Increment(value int) (err error) {
	if (value+ic.pc) > ic.Len() || (value+ic.pc) < 0 {
		// greater than since we can assume pc equals Len()
		return NewOutOfBoundsError(value+ic.pc, ic.Len())
	}
	ic.pc += value
	return
}

// Jump jumps the program counter to some value
func (ic *IntCode) Jump(value int) (err error) {
	if value > ic.Len() || value < 0 {
		// greater than since we can assume pc equals Len() (where it will just halt)
		return NewOutOfBoundsError(value, ic.Len())
	}
	ic.pc = value
	return
}

// Rewind jumps PC to zero
func (ic *IntCode) Rewind() {
	ic.Jump(0) // this cannot return any error
}

// Format formats the memory, input, and outputs and sets PC to zero
// but does not remove installed modules
func (ic *IntCode) Format(mem []int) {
	ic.mem = make([]int, len(mem))
	copy(ic.mem, mem)
	ic.SetInput()
	ic.ResetOutput()
	ic.Rewind()
}

// Input returns a copy of its inputs
func (ic *IntCode) Input() (input []int) {
	input = make([]int, len(ic.input))
	for ii := range input {
		input[ii] = ic.input[ii]
	}
	return
}

// PushToInput pushes a value to the input queue
func (ic *IntCode) PushToInput(input int) {
	ic.input = append(ic.input, input)
	return
}

// SetInput sets the input
func (ic *IntCode) SetInput(inputs ...int) {
	ic.input = make([]int, len(inputs))
	for ii := range inputs {
		ic.input[ii] = inputs[ii]
	}
}

// GetInput removes an input from the queue
func (ic *IntCode) GetInput() (input int, err error) {
	if len(ic.input) == 0 {
		err = fmt.Errorf("input is length zero")
		return
	}
	input = ic.input[len(ic.input)-1]
	ic.input = ic.input[:len(ic.input)-1]
	return
}

// ResetOutput resets the outputs
func (ic *IntCode) ResetOutput() {
	ic.output = make([]int, 0)
}

// PushToOutput pushes a value to its outputs
func (ic *IntCode) PushToOutput(value int) {
	ic.output = append(ic.output, value)
}

// Output prints the output
func (ic *IntCode) Output() (output []int) {
	output = make([]int, len(ic.output))
	for ii := range output {
		output[ii] = ic.output[ii]
	}
	return
}

// GetOutput removes an output from the stack
func (ic *IntCode) GetOutput() (output int, err error) {
	if len(ic.output) == 0 {
		err = fmt.Errorf("output is length zero")
		return
	}
	output = ic.output[len(ic.output)-1]
	ic.output = ic.output[:len(ic.output)-1]
	return
}

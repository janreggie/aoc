package intcode

import "fmt"

// getFromMemory turns an int flags (opcode without 2 LSDs), a slice of parameters, and an IntCode reel
// into the appropriate parameters, depending on slice inputs.
// e.g., if flags is 10 (from LSD: position, immediate),
// parameters is []int{4,3} (excluding last one)
// then parameters becomes []int{4, GetLocation(3)}.
// May return an error if OutOfBounds.
func getFromMemory(flags int, parameters []int, ic *IntCode) (err error) {
	for ii := range parameters {
		switch flag := flags % 10; flag {
		case 0: // position mode
			if parameters[ii], err = ic.GetLocation(parameters[ii]); err != nil {
				return
			}
		case 1: // immediate mode
			// do nothing
		default:
			return fmt.Errorf("unimplemented flag (%v)", flag)
		}
		flags /= 10
	}
	return
}

// SimpleAdder is a simple program that adds values.
// Adapted from https://adventofcode.com/2019/day/2.
//
// Memory:
//  1 ARG1 ARG2 ARG3
// Procedure:
//  mem[ARG3] = mem[ARG1]+mem[ARG2]
//  pc += 4
var SimpleAdder *Module = &Module{
	// mem: 1 ARG1 ARG2 ARG3
	// add: mem[ARG3] = mem[ARG1]+mem[ARG2]
	opcode:   1,
	mnemonic: "ADD",
	function: func(ic *IntCode) (err error) {
		// assume that Current() is 1
		// Now check if the next ones are in memory
		var params []int
		if params, err = ic.GetNext(3); err != nil {
			return
		}
		if params[0], err = ic.GetLocation(params[0]); err != nil {
			return
		}
		if params[1], err = ic.GetLocation(params[1]); err != nil {
			return
		}
		if err = ic.SetLocation(params[2], params[0]+params[1]); err != nil {
			return
		}
		return ic.Increment(4)
	},
}

// SimpleMultiplier is a simple program that adds values.
// Adapted from https://adventofcode.com/2019/day/2.
//
// Memory:
//  2 ARG1 ARG2 ARG3
// Procedure:
//  mem[ARG3] = mem[ARG1]*mem[ARG2]
//  pc += 4
//
var SimpleMultiplier *Module = &Module{
	// mem: 2 ARG1 ARG2 ARG3
	// mul: mem[ARG3] = mem[ARG1]*mem[ARG2]
	opcode:   2,
	mnemonic: "MUL",
	function: func(ic *IntCode) (err error) {
		// assume that Current() is 2
		// Now check if the next ones are in memory
		var params []int
		if params, err = ic.GetNext(3); err != nil {
			return
		}
		if params[0], err = ic.GetLocation(params[0]); err != nil {
			return
		}
		if params[1], err = ic.GetLocation(params[1]); err != nil {
			return
		}
		if err = ic.SetLocation(params[2], params[0]*params[1]); err != nil {
			return
		}
		return ic.Increment(4)
	},
}

// Adder is a module that adds values with support for parameterized mode.
// Adapted from https://adventofcode.com/2019/day/5.
//
// Memory:
//  1 ARG1 ARG2 ARG3
// Procedure:
//  mem[ARG3] = mem[ARG1]+mem[ARG2]
//  pc += 4
var Adder *Module = &Module{
	opcode:        1,
	mnemonic:      "ADD",
	parameterized: true,
	function: func(ic *IntCode) (err error) {
		var params []int
		if params, err = ic.GetNext(3); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params[0:2], ic); err != nil {
			return
		}
		if err = ic.SetLocation(params[2], params[0]+params[1]); err != nil {
			return
		}
		return ic.Increment(4)
	},
}

// Multiplier is a module that adds values with support for parameterized mode.
// Adapted from https://adventofcode.com/2019/day/5.
//
// Memory:
//  2 ARG1 ARG2 ARG3
// Procedure:
//  mem[ARG3] = mem[ARG1]*mem[ARG2]
//  pc += 4
var Multiplier *Module = &Module{
	opcode:        2,
	mnemonic:      "MUL",
	parameterized: true,
	function: func(ic *IntCode) (err error) {
		var params []int
		if params, err = ic.GetNext(3); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params[0:2], ic); err != nil {
			return
		}
		if err = ic.SetLocation(params[2], params[0]*params[1]); err != nil {
			return
		}
		return ic.Increment(4)
	},
}

// Inputter reads from the input and sets it to a specific address
//
// Memory:
//  3 ARG1
// Procedure:
//  mem[ARG1] = input
//  pc += 2
var Inputter *Module = &Module{
	opcode:        3,
	parameterized: false,
	mnemonic:      "INPUT",
	function: func(ic *IntCode) (err error) {
		var params []int
		if params, err = ic.GetNext(1); err != nil {
			return
		}
		if err = ic.SetLocation(params[0], ic.input); err != nil {
			return
		}
		return ic.Increment(2)
	},
}

// Outputter is a module that outputs the value at its only parameter.
//
// Memory:
//  4 ARG1
// Procedure:
//  output = append(output, mem[ARG1])
//  pc += 2
var Outputter *Module = &Module{
	opcode:        4,
	parameterized: true,
	mnemonic:      "OUTPUT",
	function: func(ic *IntCode) (err error) {
		var params []int
		if params, err = ic.GetNext(1); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic); err != nil {
			return
		}
		ic.PushToOutput(params[0])
		return ic.Increment(2)
	},
}

// OutputAndHalt is a module that outputs the value at its only parameter
// and, if non-zero, will halt immediately. Used for aoc2019.Day05.
//
// Memory:
//  4 ARG1
// Procedure:
//  output = append(output, mem[ARG1])
//  if mem[ARG1] != 0 then halt
//  pc += 2
var OutputAndHalt *Module = &Module{
	opcode:        4,
	parameterized: true,
	mnemonic:      "OUTPUT",
	function: func(ic *IntCode) (err error) {
		var params []int
		if params, err = ic.GetNext(1); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic); err != nil {
			return
		}
		ic.PushToOutput(params[0])
		if params[0] != 0 {
			return NewHaltError("OUTPUT (4)")
		}
		return ic.Increment(2)
	},
}

// JumpIfTrue is a module that sets the instruction pointer to the second parameter
// if the first parameter is non-zero
//
// Memory:
//  5 ARG1 ARG2
// Procedure:
//  if mem[ARG1] != 0 then jump(mem[ARG2])
//  pc += 3
var JumpIfTrue *Module = &Module{
	opcode:        5,
	parameterized: true,
	mnemonic:      "JUMPIFTRUE",
	function: func(ic *IntCode) (err error) {
		var params []int
		if params, err = ic.GetNext(2); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic); err != nil {
			return
		}
		// from here on, params[0] would be the value we would check
		// and params[1] would be where we would want to jump
		if params[0] != 0 {
			return ic.Jump(params[1])
		}
		return ic.Increment(3)
	},
}

// JumpIfFalse is a module that sets the instruction pointer to the second parameter
// if the first parameter is zero
//
// Memory:
//  5 ARG1 ARG2
// Procedure:
//  if mem[ARG1] == 0 then jump(mem[ARG2])
//  pc += 3
var JumpIfFalse *Module = &Module{
	opcode:        6,
	parameterized: true,
	mnemonic:      "JUMPIFTRUE",
	function: func(ic *IntCode) (err error) {
		var params []int
		if params, err = ic.GetNext(2); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic); err != nil {
			return
		}
		// from here on, params[0] would be the value we would check
		// and params[1] would be where we would want to jump
		if params[0] == 0 {
			return ic.Jump(params[1])
		}
		return ic.Increment(3)
	},
}

// LessThan is a module that stores 1 in the third parameter
// if the first is less than the second; otherwise it will store 0
//
// Memory:
//  7 ARG1 ARG2 ARG3
// Procedure:
//  if mem[ARG1] < mem[ARG2] then mem[ARG3]=1 else mem[ARG3]=0
//  pc += 4
var LessThan *Module = &Module{
	opcode:        7,
	parameterized: true,
	mnemonic:      "LESSTHAN",
	function: func(ic *IntCode) (err error) {
		var params []int
		if params, err = ic.GetNext(3); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params[0:2], ic); err != nil {
			return
		}
		// now write to mem[params[2]] depending on what's with params[0]&params[1]
		if params[0] < params[1] {
			if err = ic.SetLocation(params[2], 1); err != nil {
				return
			}
		} else {
			if err = ic.SetLocation(params[2], 0); err != nil {
				return
			}
		}
		return ic.Increment(4)
	},
}

// Equals is a module that stores 1 in the third parameter
// if the first equals the second; otherwise it will store 0
//
// Memory:
//  8 ARG1 ARG2 ARG3
// Procedure:
//  if mem[ARG1] == mem[ARG2] then mem[ARG3]=1 else mem[ARG3]=0
//  pc += 4
var Equals *Module = &Module{
	opcode:        8,
	parameterized: true,
	mnemonic:      "EQUALS",
	function: func(ic *IntCode) (err error) {
		var params []int
		if params, err = ic.GetNext(3); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params[0:2], ic); err != nil {
			return
		}
		// now write to mem[params[2]] depending on what's with params[0]&params[1]
		if params[0] == params[1] {
			if err = ic.SetLocation(params[2], 1); err != nil {
				return
			}
		} else {
			if err = ic.SetLocation(params[2], 0); err != nil {
				return
			}
		}
		return ic.Increment(4)
	},
}

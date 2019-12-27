package intcode

import "fmt"

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
	opcode        int64                   // opcode (if 0 then check will occur in function)
	mnemonic      string                  // "name" of the opcode
	parameterized bool                    // should module support parameter modes?
	function      func(ic *IntCode) error // what does it do to the computer?
}

// ModuleConfig is a structure representing the configuration of a module
type ModuleConfig struct {
	Opcode        int64                   // opcode (if 0 then check will occur in function)
	Mnemonic      string                  // "name" of the opcode
	Parameterized bool                    // should module support parameter modes?
	Function      func(ic *IntCode) error // what does it do to the computer?
}

// NewModule generates a module object with several attributes using a config struct
func NewModule(config ModuleConfig) *Module {
	return &Module{
		opcode:        config.Opcode,
		mnemonic:      config.Mnemonic,
		parameterized: config.Parameterized,
		function:      config.Function,
	}
}

// getFromMemory turns an int flags (opcode without 2 LSDs), a slice of parameters, and an IntCode reel
// into the appropriate parameters, depending on slice inputs.
// The parameters are taken from the memory, raw,
// that is, the flags are unknown.
// If writeEnable is true, it is assumed that parameters[len(parameters)-1]
// is the address to be written,
// making it as a value as is, or added by ic.RelativeBase().
// The parameter slice is then modified, depending on said flags.
//
// Example: if flags is 10 (from LSD: position, immediate, position),
// parameters is []int64{4,3, 2}, writeEnable is true (thus using the raw value 2),
// then parameters becomes []int64{4, GetLocation(3), 3}.
//
// Modes:
// * Position mode (flag 0): use memory location
// * Immediate mode (flag 1): use the value itself
// * Relative mode (flag 2): use memory location plus the relative base of the computer
//
//
// May return an error if OutOfBounds.
func getFromMemory(flags int64, parameters []int64, ic *IntCode, writeEnable bool) (err error) {
	parameterCount := len(parameters)
	// to ignore or not to ignore parameters[len(parameters)-1]?
	// check writeEnable
	if writeEnable {
		// then parameters[len(parameters)-1] would be direct
		// therefore reduce the number of "parameters" to get
		parameterCount--
	}
	for ii := 0; ii < parameterCount; ii++ {
		switch flag := flags % 10; flag {
		case 0: // position mode
			if parameters[ii], err = ic.GetLocation(parameters[ii]); err != nil {
				return
			}
		case 1: // immediate mode
			// do nothing
		case 2: // relative mode
			if parameters[ii], err = ic.GetLocation(parameters[ii] + ic.RelativeBase()); err != nil {
				return
			}
		default:
			return fmt.Errorf("unimplemented flag (%v)", flag)
		}
		flags /= 10
	}
	// now for the last flag...
	if writeEnable {
		switch flag := flags % 10; flag {
		case 0, 1: // position mode
			// do nothing
		case 2: // relative mode
			parameters[parameterCount] = parameters[parameterCount] + ic.RelativeBase()
		}
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
		var params []int64
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
		var params []int64
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
		var params []int64
		if params, err = ic.GetNext(3); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic, true); err != nil {
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
		var params []int64
		if params, err = ic.GetNext(3); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic, true); err != nil {
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
//  mem[ARG1], err = ic.GetInput()
//  pc += 2
var Inputter *Module = &Module{
	opcode:        3,
	parameterized: true, // because apparently 203 exists...
	mnemonic:      "INPUT",
	function: func(ic *IntCode) (err error) {
		var params []int64
		var input int64
		if params, err = ic.GetNext(1); err != nil {
			return
		}
		if input, err = ic.GetInput(); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic, true); err != nil {
			return
		}
		if err = ic.SetLocation(params[0], input); err != nil {
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
		var params []int64
		if params, err = ic.GetNext(1); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic, false); err != nil {
			return
		}
		ic.PushToOutput(params[0])
		return ic.Increment(2)
	},
}

// OutputToInput is a moule that, instead of pushing its parameter to Output,
// it pushes the value to Input
var OutputToInput *Module = &Module{
	opcode:        4,
	parameterized: true,
	mnemonic:      "OUTPUT",
	function: func(ic *IntCode) (err error) {
		var params []int64
		if params, err = ic.GetNext(1); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic, false); err != nil {
			return
		}
		ic.PushToInput(params[0])
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
		var params []int64
		if params, err = ic.GetNext(1); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic, false); err != nil {
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
		var params []int64
		if params, err = ic.GetNext(2); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic, false); err != nil {
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
		var params []int64
		if params, err = ic.GetNext(2); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic, false); err != nil {
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
		var params []int64
		if params, err = ic.GetNext(3); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic, true); err != nil {
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
		var params []int64
		if params, err = ic.GetNext(3); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic, true); err != nil {
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

// ChangeRelativeBase adjusts the relative base of the computer
// by its parameter.
//
// Memory:
//  9 ARG1
// Procedure:
//  relativeBase += mem[ARG1]
//  pc += 2
var ChangeRelativeBase *Module = &Module{
	opcode:        9,
	parameterized: true,
	mnemonic:      "RELBASE",
	function: func(ic *IntCode) (err error) {
		var params []int64
		if params, err = ic.GetNext(1); err != nil {
			return
		}
		if err = getFromMemory(ic.Current()/100, params, ic, false); err != nil {
			return
		}
		// now pull out an ic.AdjustRelativeBase
		ic.AdjustRelativeBase(params[0])
		return ic.Increment(2)
	},
}

// InstallAdderMultiplier installs the Adder and Multiplier modules
// to the IntCode computer
func InstallAdderMultiplier(ic *IntCode) {
	ic.Install(Adder)
	ic.Install(Multiplier)
	return
}

// InstallJumpers installs the
// JumpIfFalse, JumpIfTrue, LessThan, and Equals modules
func InstallJumpers(ic *IntCode) {
	ic.Install(JumpIfFalse)
	ic.Install(JumpIfTrue)
	ic.Install(LessThan)
	ic.Install(Equals)
}

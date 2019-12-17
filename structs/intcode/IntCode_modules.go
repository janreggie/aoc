package intcode

// SimpleAdder is a simple program that adds values.
// Adapted from https://adventofcode.com/2019/day/2.
//
// Memory:
//  1 LOC1 LOC2 LOC3
// Procedure:
//  mem[LOC3] = mem[LOC1]+mem[LOC2]
//  pc += 4
var SimpleAdder *Module = &Module{
	// mem: 1 LOC1 LOC2 LOC3
	// add: mem[LOC3] = mem[LOC1]+mem[LOC2]
	Opcode:   1,
	Mnemonic: "ADD",
	Function: func(ic *IntCode) (err error) {
		// assume that Current() is 1
		// Now check if the next ones are in memory
		var params []int
		if params, err = ic.GetNext(3); err != nil {
			return err
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
//  2 LOC1 LOC2 LOC3
// Procedure:
//  mem[LOC3] = mem[LOC1]*mem[LOC2]
//  pc += 4
//
var SimpleMultiplier *Module = &Module{
	// mem: 2 LOC1 LOC2 LOC3
	// mul: mem[LOC3] = mem[LOC1]*mem[LOC2]
	Opcode:   2,
	Mnemonic: "MUL",
	Function: func(ic *IntCode) (err error) {
		// assume that Current() is 2
		// Now check if the next ones are in memory
		var params []int
		if params, err = ic.GetNext(3); err != nil {
			return err
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

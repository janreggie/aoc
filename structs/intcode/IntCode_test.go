package intcode_test

import (
	"fmt"
	"testing"

	"github.com/janreggie/AdventOfCode/structs/intcode"
	"github.com/stretchr/testify/assert"
)

// This example tries to implement Addition and Multiplication modules
func ExampleModule_Add_and_Multiply() {
	add := &intcode.Module{
		// mem: 1 LOC1 LOC2 LOC3
		// add: mem[LOC3] = mem[LOC1]+mem[LOC2]
		Opcode:     1,
		Mnemonic:   "AND",
		ParamCount: 3,
		Function: func(ic *intcode.IntCode) (err error) {
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
	mul := &intcode.Module{
		// mem: 2 LOC1 LOC2 LOC3
		// mul: mem[LOC3] = mem[LOC1]*mem[LOC2]
		Opcode:     2,
		Mnemonic:   "MUL",
		ParamCount: 3,
		Function: func(ic *intcode.IntCode) (err error) {
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
	// let's use a simple computer
	ic := intcode.New([]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50})
	ic.Install(add)
	ic.Install(mul)
	err := ic.Operate()
	_, isHalt := err.(*intcode.HaltError)
	fmt.Printf("halted properly: %v\n", isHalt)
	fmt.Println(ic.Snapshot())
	// Output: halted properly: true
	// [3500 9 10 70 2 3 11 0 99 30 40 50]
}

func TestIntcode_Operate(t *testing.T) {
	assert := assert.New(t)
	ic := intcode.New([]int{})
	ic.Install(intcode.SimpleAdder)
	ic.Install(intcode.SimpleMultiplier)
	type testCase struct {
		input  []int
		output []int
	}
	cases := []testCase{
		{input: []int{1, 0, 0, 0, 99}, output: []int{2, 0, 0, 0, 99}},
		{input: []int{2, 3, 0, 3, 99}, output: []int{2, 3, 0, 6, 99}},
		{input: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, output: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{input: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, output: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
	}
	for _, eachCase := range cases {
		ic.Format(eachCase.input)
		err := ic.Operate()
		assert.True(intcode.IsHalt(err), err)
		assert.Equal(eachCase.output, ic.Snapshot())
	}
}

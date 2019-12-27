package intcode_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/structs/intcode"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name    string
	initial []int64
	input   []int64
	output  []int64
}

func processCases(cases []testCase, assert *assert.Assertions, ic *intcode.IntCode) {
	for _, eachCase := range cases {
		ic.Format(eachCase.initial)
		ic.SetInput(eachCase.input...)
		err := ic.Operate()
		assert.True(intcode.IsHalt(err), err, eachCase.name)
		// assert.Equal(eachCase.final, ic.Snapshot(), eachCase.name)
		assert.Equal(eachCase.output, ic.Output(), eachCase.name)
	}
}

func TestInputAndOutput(t *testing.T) {
	assert := assert.New(t)
	ic := intcode.New([]int64{})
	ic.Install(intcode.Inputter)
	ic.Install(intcode.Outputter)
	ic.Install(intcode.ChangeRelativeBase)
	testCases := []testCase{
		{name: "copy input to output: position mode",
			initial: []int64{3, 0, 4, 0, 99},
			input:   []int64{4},
			// final:   []int64{4, 0, 4, 0, 99},
			output: []int64{4}},
		{name: "copy input to output: absolute mode",
			initial: []int64{3, 0, 104, 8, 99},
			input:   []int64{6},
			// final:   []int64{6, 0, 104, 8, 99},
			output: []int64{8}},
		{name: "copy input to output: relative mode",
			initial: []int64{109, 1, 203, 6, 204, 6, 99, 6},
			input:   []int64{4},
			output:  []int64{4}},
	}
	processCases(testCases, assert, ic)
}

func TestCompare(t *testing.T) {
	assert := assert.New(t)
	ic := intcode.New([]int64{})
	ic.Install(intcode.Adder)
	ic.Install(intcode.Multiplier)
	ic.Install(intcode.Inputter)
	ic.Install(intcode.Outputter)
	ic.Install(intcode.LessThan)
	ic.Install(intcode.Equals)
	cases := []testCase{
		{name: "using Equals (position) if input == 8 output 1 else 0; input:8",
			initial: []int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input:   []int64{8},
			output:  []int64{1}},
		{name: "using Equals (position) if input == 8 output 1 else 0; input:5",
			initial: []int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input:   []int64{5},
			output:  []int64{0}},
		{name: "using LessThan (position) if input < 8 output 1 else 0; input:7;",
			initial: []int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input:   []int64{7},
			output:  []int64{1}},
		{name: "using LessThan (position) if input < 8 output 1 else 0; input:8;",
			initial: []int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input:   []int64{8},
			output:  []int64{0}},
		{name: "using Equals (immediate) if input == 8 output 1 else 0; input:5",
			initial: []int64{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input:   []int64{5},
			output:  []int64{0}},
		{name: "using Equals (immediate) if input == 8 output 1 else 0; input:8",
			initial: []int64{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input:   []int64{8},
			output:  []int64{1}},
		{name: "using Lessthan (immediate) if input < 8 output 1 else 0; input:5",
			initial: []int64{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input:   []int64{5},
			output:  []int64{1}},
		{name: "using Lessthan (immediate) if input < 8 output 1 else 0; input:8",
			initial: []int64{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input:   []int64{8},
			output:  []int64{0}},
	}
	processCases(cases, assert, ic)
}

func TestJump(t *testing.T) {
	assert := assert.New(t)
	ic := intcode.New([]int64{})
	ic.Install(intcode.Adder)
	ic.Install(intcode.Multiplier)
	ic.Install(intcode.Inputter)
	ic.Install(intcode.Outputter)
	ic.Install(intcode.JumpIfTrue)
	ic.Install(intcode.JumpIfFalse)
	cases := []testCase{
		{name: "using Jump (position) if input == 0 output 0 else 1; input:0",
			initial: []int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input:   []int64{0},
			output:  []int64{0}},
		{name: "using Jump (position) if input == 0 output 0 else 1; input:3",
			initial: []int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input:   []int64{3},
			output:  []int64{1}},
	}
	processCases(cases, assert, ic)
}

func TestLEQ(t *testing.T) {
	assert := assert.New(t)
	ic := intcode.New([]int64{})
	ic.Install(intcode.Adder)
	ic.Install(intcode.Multiplier)
	ic.Install(intcode.Inputter)
	ic.Install(intcode.Outputter)
	ic.Install(intcode.JumpIfTrue)
	ic.Install(intcode.JumpIfFalse)
	ic.Install(intcode.LessThan)
	ic.Install(intcode.Equals)
	cases := []testCase{
		{name: "output = input < 8 ? 999 : input == 8 ? 1000 : 1001; input: 7",
			initial: []int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:  []int64{7},
			output: []int64{999}},
		{name: "output = input < 8 ? 999 : input == 8 ? 1000 : 1001; input: 8",
			initial: []int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:  []int64{8},
			output: []int64{1000}},
		{name: "output = input < 8 ? 999 : input == 8 ? 1000 : 1001; input: 9",
			initial: []int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:  []int64{9},
			output: []int64{1001}},
	}
	processCases(cases, assert, ic)
}

func TestChangeRelativeBase(t *testing.T) {
	assert := assert.New(t)
	ic := intcode.New([]int64{})
	ic.Install(intcode.Adder)
	ic.Install(intcode.Multiplier)
	ic.Install(intcode.Inputter)
	ic.Install(intcode.Outputter)
	ic.Install(intcode.JumpIfTrue)
	ic.Install(intcode.JumpIfFalse)
	ic.Install(intcode.LessThan)
	ic.Install(intcode.Equals)
	ic.Install(intcode.ChangeRelativeBase)
	cases := []testCase{
		{name: "quine",
			initial: []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			input:   []int64{},
			output:  []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		}}
	processCases(cases, assert, ic)
}

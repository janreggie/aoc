package intcode_test

import (
	"fmt"
	"testing"

	"github.com/janreggie/AdventOfCode/tools/intcode"
	"github.com/stretchr/testify/assert"
)

// This example tries to implement Addition and Multiplication modules
func ExampleNewModule() {
	// mem: 1 LOC1 LOC2 LOC3
	// add: mem[LOC3] = mem[LOC1]+mem[LOC2]
	add := intcode.NewModule(struct {
		Opcode        int64
		Mnemonic      string
		Parameterized bool
		Function      func(ic *intcode.Intcode) error
	}{
		Opcode: 1, Mnemonic: "ADD",
		Function: func(ic *intcode.Intcode) (err error) {
			// assume that Current() is 1
			// Now check if the next ones are in memory
			var params []int64
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
		}})
	// mem: 2 LOC1 LOC2 LOC3
	// mul: mem[LOC3] = mem[LOC1]*mem[LOC2]
	mul := intcode.NewModule(struct {
		Opcode        int64
		Mnemonic      string
		Parameterized bool
		Function      func(ic *intcode.Intcode) error
	}{
		Opcode: 2, Mnemonic: "MUL",
		Function: func(ic *intcode.Intcode) (err error) {
			// assume that Current() is 2
			// Now check if the next ones are in memory
			var params []int64
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
	})
	// let's use a simple computer
	ic := intcode.New([]int64{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50})
	ic.Install(add)
	ic.Install(mul)
	err := ic.Operate()
	_, isHalt := err.(*intcode.HaltError)
	fmt.Printf("halted properly: %v\n", isHalt)
	fmt.Println(ic.Snapshot())
	// Output: halted properly: true
	// [3500 9 10 70 2 3 11 0 99 30 40 50]
}

func TestIntCode_Operate(t *testing.T) {
	assert := assert.New(t)
	ic := intcode.New([]int64{})
	ic.Install(intcode.Adder)
	ic.Install(intcode.Multiplier)
	type testCase struct {
		name    string
		initial []int64
		final   []int64
	}
	cases := []testCase{
		{name: "simple addition",
			initial: []int64{1, 0, 0, 0, 99},
			final:   []int64{2, 0, 0, 0, 99}},
		{name: "simple multiplication",
			initial: []int64{2, 3, 0, 3, 99},
			final:   []int64{2, 3, 0, 6, 99}},
		{name: "combining addition and multiplication",
			initial: []int64{1, 1, 1, 4, 99, 5, 6, 0, 99},
			final:   []int64{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{name: "combining addition and multiplication 2.0",
			initial: []int64{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			final:   []int64{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
		{name: "parametric mode: using immediate mode",
			initial: []int64{1002, 4, 3, 4, 33},
			final:   []int64{1002, 4, 3, 4, 99}},
	}
	for _, eachCase := range cases {
		ic.Format(eachCase.initial)
		err := ic.Operate()
		assert.True(intcode.IsHalt(err), err, eachCase.name)
		assert.Equal(eachCase.final, ic.Snapshot(), eachCase.name)
	}
}

func TestIntCode_Int64Support(t *testing.T) {
	assert := assert.New(t)
	ic := intcode.New([]int64{})
	intcode.InstallAdderMultiplier(ic)
	intcode.InstallJumpers(ic)
	ic.Install(intcode.Inputter)
	ic.Install(intcode.Outputter)
	type testCase struct {
		name   string
		memory []int64
		output []int64
	}
	cases := []testCase{
		{name: "34915192 times 34915192 equals 1219070632396864",
			memory: []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			output: []int64{1219070632396864}},
		{name: "1125899906842624 is a really large number",
			memory: []int64{104, 1125899906842624, 99},
			output: []int64{1125899906842624}},
	}
	for _, eachCase := range cases {
		ic.Format(eachCase.memory)
		err := ic.Operate()
		assert.True(intcode.IsHalt(err), err, eachCase.name)
		assert.Equal(eachCase.output, ic.Output(), eachCase.name)
	}
}

func TestIntCode_DynamicMemory(t *testing.T) {
	assert := assert.New(t)
	ic := intcode.New([]int64{})
	intcode.InstallAdderMultiplier(ic)
	intcode.InstallJumpers(ic)
	ic.Install(intcode.Inputter)
	ic.Install(intcode.Outputter)
	type testCase struct {
		name    string
		initial []int64
		final   []int64
		output  []int64
	}
	cases := []testCase{
		{name: "saves 3 to some memory location and returns it",
			initial: []int64{1001, 7, 3, 8, 4, 8, 99},
			final:   []int64{1001, 7, 3, 8, 4, 8, 99, 0, 3},
			output:  []int64{3}},
		{name: "saves 3 to some memory location and returns 0",
			initial: []int64{1001, 7, 3, 9, 4, 10, 99},
			final:   []int64{1001, 7, 3, 9, 4, 10, 99, 0, 0, 3, 0},
			output:  []int64{0}},
	}
	for _, eachCase := range cases {
		ic.Format(eachCase.initial)
		err := ic.Operate()
		assert.True(intcode.IsHalt(err), err, eachCase.name)
		assert.Equal(eachCase.output, ic.Output(), eachCase.name)
		assert.Equal(eachCase.final, ic.Snapshot(), eachCase.name)
	}
}

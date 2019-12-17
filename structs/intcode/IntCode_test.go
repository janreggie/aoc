package intcode_test

import (
	"fmt"
	"testing"

	"github.com/janreggie/AdventOfCode/structs/intcode"
	"github.com/stretchr/testify/assert"
)

// This example tries to implement Addition and Multiplication modules
func ExampleNewModule() {
	// mem: 1 LOC1 LOC2 LOC3
	// add: mem[LOC3] = mem[LOC1]+mem[LOC2]
	add := intcode.NewModule(struct {
		Opcode        int
		Mnemonic      string
		Parameterized bool
		Function      func(ic *intcode.IntCode) error
	}{
		Opcode: 1, Mnemonic: "ADD",
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
		}})
	// mem: 2 LOC1 LOC2 LOC3
	// mul: mem[LOC3] = mem[LOC1]*mem[LOC2]
	mul := intcode.NewModule(struct {
		Opcode        int
		Mnemonic      string
		Parameterized bool
		Function      func(ic *intcode.IntCode) error
	}{
		Opcode: 2, Mnemonic: "MUL",
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
	})
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
	ic.Install(intcode.Adder)
	ic.Install(intcode.Multiplier)
	type testCase struct {
		initial []int
		final   []int
	}
	cases := []testCase{
		{initial: []int{1, 0, 0, 0, 99}, final: []int{2, 0, 0, 0, 99}},
		{initial: []int{2, 3, 0, 3, 99}, final: []int{2, 3, 0, 6, 99}},
		{initial: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, final: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{initial: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, final: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
		{initial: []int{1002, 4, 3, 4, 33}, final: []int{1002, 4, 3, 4, 99}},
	}
	for _, eachCase := range cases {
		ic.Format(eachCase.initial)
		err := ic.Operate()
		assert.True(intcode.IsHalt(err), err)
		assert.Equal(eachCase.final, ic.Snapshot())
	}
}

package aoc2019_test

import (
	"fmt"
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2019"
	"github.com/janreggie/AdventOfCode/structs/intcode"
	"github.com/stretchr/testify/assert"
)

type queue struct {
	internal []int
}

func newqueue(values ...int) *queue {
	v := make([]int, len(values))
	copy(v, values)
	return &queue{internal: v}
}

func (q *queue) push(val int) {
	q.internal = append(q.internal, val)
}

func (q *queue) pop() (val int) {
	if len(q.internal) == 0 {
		panic("queue of length zero")
	}
	val = q.internal[0]
	q.internal = q.internal[1:]
	return
}

func inspectIntCode(ic *intcode.IntCode) {
	// let's take a look shall we?
	fmt.Printf("State: PC: %v (%v)\n", ic.PC(), ic.Current())
	fmt.Printf("Memory: %v\n", ic.Snapshot())
	fmt.Printf("Input: %v\n", ic.Input())
	fmt.Printf("Output: %v\n", ic.Output())
}

// Honestly this is just me playing around
// with the specifications for Day 07
func TestDay07_scratchpad_d1(t *testing.T) {
	var ic *intcode.IntCode
	var err error
	var q *queue
	ic = intcode.New([]int{})
	intcode.InstallAdderMultiplier(ic)
	ic.Install(intcode.Inputter)
	ic.Install(intcode.Outputter)
	intcode.InstallJumpers(ic)

	// Use program 3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0
	// and phase 4,3,2,1,0
	q = newqueue(4, 3, 2, 1, 0)
	ic.Format([]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0})
	ic.PushToOutput(0)
	// put in phase setting 4 and input signal 0
	for len(q.internal) > 0 {
		ic.Rewind()
		ic.PushToInput(q.pop())
		if out, err := ic.GetOutput(); err != nil {
			panic(err)
		} else {
			ic.PushToInput(out)
		}
		if err = ic.Operate(); !intcode.IsHalt(err) {
			panic(err)
		}
		inspectIntCode(ic)
	}
	panic(nil)
}

func TestDay07_scratchpad_d2(t *testing.T) {
	var ic *intcode.IntCode
	var err error

	// permutation := []int{9, 8, 7, 6, 5}
	// memory := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
	// 	27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
	permutation := []int{9, 7, 8, 5, 6}
	memory := []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
		-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
		53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}

	ic = intcode.New(memory)
	intcode.InstallAdderMultiplier(ic)
	intcode.InstallJumpers(ic)
	ic.Install(intcode.Inputter)
	ic.Install(intcode.OutputToInput)
	ic.PushToInput(0)
	for ii := range permutation {
		ic.Rewind()
		// input, output := ic.Input(), ic.Output()
		// ic.Format(memory)
		// ic.SetInput(input...)
		// for _, v := range output {
		// 	ic.PushToInput(v)
		// }
		ic.PushToInput(permutation[ii])
		if err = ic.Operate(); !intcode.IsHalt(err) {
			panic(err)
		}
		inspectIntCode(ic)
	}
	panic(nil)
}

func TestDay07(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{"3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", "43210", ""},
		{"3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0", "54321", ""},
		{"3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0", "65210", ""},
		// {"3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5", "", "139629729"},
		// {"3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10", "", "18216"},
	}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2019.Day07, assert)
	}
}

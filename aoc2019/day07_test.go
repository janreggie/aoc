package aoc2019

import (
	"testing"

	"github.com/janreggie/AdventOfCode/internal"
	"github.com/stretchr/testify/assert"
)

const day07myInput = `3,8,1001,8,10,8,105,1,0,0,21,42,67,84,109,122,203,284,365,446,99999,3,9,1002,9,3,9,1001,9,5,9,102,4,9,9,1001,9,3,9,4,9,99,3,9,1001,9,5,9,1002,9,3,9,1001,9,4,9,102,3,9,9,101,3,9,9,4,9,99,3,9,101,5,9,9,1002,9,3,9,101,5,9,9,4,9,99,3,9,102,5,9,9,101,5,9,9,102,3,9,9,101,3,9,9,102,2,9,9,4,9,99,3,9,101,2,9,9,1002,9,3,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,99`

type queue struct {
	internal []int64
}

func newqueue(values ...int64) *queue {
	v := make([]int64, len(values))
	copy(v, values)
	return &queue{internal: v}
}

func (q *queue) push(val int64) {
	q.internal = append(q.internal, val)
}

func (q *queue) pop() (val int64) {
	if len(q.internal) == 0 {
		panic("queue of length zero")
	}
	val = q.internal[0]
	q.internal = q.internal[1:]
	return
}

// // Honestly this is just me playing around
// // with the specifications for Day 07
// func TestDay07_scratchpad_d1(t *testing.T) {
// 	var ic *intcode.Intcode
// 	var err error
// 	var q *queue
// 	ic = intcode.New([]int64{})
// 	intcode.InstallAdderMultiplier(ic)
// 	ic.Install(intcode.Inputter)
// 	ic.Install(intcode.Outputter)
// 	intcode.InstallJumpers(ic)

// 	// Use program 3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0
// 	// and phase 4,3,2,1,0
// 	q = newqueue(4, 3, 2, 1, 0)
// 	ic.Format([]int64{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0})
// 	ic.PushToOutput(0)
// 	// put in phase setting 4 and input signal 0
// 	for len(q.internal) > 0 {
// 		ic.Rewind()
// 		ic.PushToInput(q.pop())
// 		if out, err := ic.GetOutput(); err != nil {
// 			panic(err)
// 		} else {
// 			ic.PushToInput(out)
// 		}
// 		if err = ic.Operate(); !intcode.IsHalt(err) {
// 			panic(err)
// 		}
// 		inspectIntCode(ic)
// 	}
// 	panic(nil)
// }

// func TestDay07_scratchpad_d2(t *testing.T) {
// 	var ic *intcode.Intcode
// 	var err error

// 	// permutation := []int64{9, 8, 7, 6, 5}
// 	// memory := []int64{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
// 	// 	27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
// 	permutation := []int64{9, 7, 8, 5, 6}
// 	memory := []int64{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
// 		-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
// 		53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}

// 	ic = intcode.New(memory)
// 	intcode.InstallAdderMultiplier(ic)
// 	intcode.InstallJumpers(ic)
// 	ic.Install(intcode.Inputter)
// 	ic.Install(intcode.OutputToInput)
// 	ic.PushToInput(0)
// 	for ii := range permutation {
// 		ic.Rewind()
// 		// input, output := ic.Input(), ic.Output()
// 		// ic.Format(memory)
// 		// ic.SetInput(input...)
// 		// for _, v := range output {
// 		// 	ic.PushToInput(v)
// 		// }
// 		ic.PushToInput(permutation[ii])
// 		if err = ic.Operate(); !intcode.IsHalt(err) {
// 			panic(err)
// 		}
// 		inspectIntCode(ic)
// 	}
// 	panic(nil)
// }

func TestDay07(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Details: "Y2019D07 my input",
			Input:   day07myInput,
			Result1: "262086",
			Result2: ""},
	}
	for _, tt := range testCases {
		tt.Test(Day07, assert)
	}
}

func BenchmarkDay07(b *testing.B) {
	internal.Benchmark(Day07, b, day07myInput)
}

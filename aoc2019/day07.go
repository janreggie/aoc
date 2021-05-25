package aoc2019

import (
	"strconv"

	"github.com/janreggie/aoc/aoc2019/intcode"
)

func permuteIntslice(xs []int64) (permuts [][]int64) {
	var rc func([]int64, int64)
	rc = func(a []int64, k int64) {
		if k == int64(len(a)) {
			permuts = append(permuts, append([]int64{}, a...))
		} else {
			for i := k; i < int64(len(xs)); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

// Day07 solves the seventh day puzzle "Amplification Circuit".
//
// Input
//
// A single line containing several integers separated by commas,
// representing an Intcode program. For example:
//
//	3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0
//
// It is assumed that the numbers do not exceed five digits long,
// and that the tape length is no more than 700.
//
// The first opcode of the Intcode program should be an INPUT (3).
func Day07(input string) (answer1, answer2 string, err error) {

	ic, err := intcode.NewFromString(input)
	if err != nil {
		return
	}

	memory := ic.Snapshot()
	intcode.InstallAdderMultiplier(ic)
	intcode.InstallJumpers(ic)
	ic.Install(intcode.Inputter)
	ic.Install(intcode.OutputToInput)
	ic.PushToOutput(0) // there is no previous amplifier
	var highestOutput int64

	// now suppose we do not know the phase setting sequence...
	// Due to the chaotic nature of the program,
	// a greedy solution would not help.
	// Hence let us use all permutations of the phase settings.

	// Part 1
	highestOutput = 0
	for _, permutation := range permuteIntslice([]int64{0, 1, 2, 3, 4}) {
		// now run the program...
		ic.Format(memory)
		ic.PushToInput(0) // previous amplifier
		var output int64
		for ii := range permutation {
			ic.Rewind()
			ic.PushToInput(permutation[ii])
			if err = ic.Operate(); !intcode.IsHalt(err) {
				return
			}
		}
		// check the output
		output, err = ic.GetInput()
		if err != nil {
			return
		}
		if output > highestOutput {
			highestOutput = output
		}
	}
	answer1 = strconv.FormatInt(highestOutput, 10)

	answer2 = "unimplemented"
	return
}

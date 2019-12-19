package aoc2019

import (
	"bufio"
	"strconv"

	"github.com/golang/glog"
	"github.com/janreggie/AdventOfCode/structs/intcode"
)

func inspectIntCode(ic *intcode.IntCode) {
	// let's take a look shall we?
	glog.Infof("State: PC: %v (%v)\n", ic.PC(), ic.Current())
	glog.Infof("Memory: %v\n", ic.Snapshot())
	glog.Infof("Input: %v\n", ic.Input())
	glog.Infof("Output: %v\n", ic.Output())
}

func permuteIntslice(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

// Day07 solves the seventh day puzzle
// "Amplification Circuit"
func Day07(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	var ic *intcode.IntCode
	var memory []int
	if ic, err = intcode.NewFromScanner(scanner); err != nil {
		return
	}
	memory = ic.Snapshot()
	intcode.InstallAdderMultiplier(ic)
	intcode.InstallJumpers(ic)
	ic.Install(intcode.Inputter)
	ic.Install(intcode.OutputToInput)
	ic.PushToOutput(0) // there is no previous amplifier
	var highestOutput int

	// now suppose we do not know the phase setting sequence...
	// Due to the chaotic nature of the program,
	// a greedy solution would not help.
	// Hence let us use all permutations of the phase settings.

	// Part 1
	highestOutput = 0
	for _, permutation := range permuteIntslice([]int{0, 1, 2, 3, 4}) {
		// now run the program...
		ic.Format(memory)
		ic.PushToInput(0) // previous amplifier
		var output int
		for ii := range permutation {
			// glog.Infof("at %v", ii)
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
	answer1 = strconv.Itoa(highestOutput)

	// // Part 2
	// highestOutput = 0
	// permutation := []int{9, 8, 7, 6, 5}
	// // for _, permutation := range permuteIntslice([]int{5, 6, 7, 8, 9}) {
	// // 	// now do things...
	// // 	ic.Format(memory) // go back to clean...
	// // 	_ = permutation
	// // }
	// ic.Format(memory)
	// ic.PushToInput(0)
	// for ii := range permutation {
	// 	ic.Rewind()
	// 	ic.PushToInput(permutation[ii])
	// }

	return
}

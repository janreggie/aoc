package aoc2020

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Day15 solves the fifteenth day puzzle "Rambunctious Recitation"
//
// Input
//
// A single line containing the "starting numbers" of the memory game. For example:
//
//  some sample input indented to become a code block
//
// It is guaranteed that the numbers are nonnegative.
func Day15(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	scanner.Scan()
	rawNums := strings.Split(scanner.Text(), ",")
	var memory [30000000]int
	whenLastSaid := make(map[int]int)
	for ii, vv := range rawNums {
		var e error
		memory[ii], e = strconv.Atoi(vv)
		if e != nil {
			err = errors.Wrapf(e, "couldn't parse %s from input", vv)
			return
		}
		whenLastSaid[memory[ii]] = ii
	}
	ind := len(rawNums) // at where shall we start?
	lastSpoken := memory[ind-1]

	// findWhenLastSaid finds the largest index less than ind that num appears in memory.
	// It will first check in allSaid if it is even worth in finding it.
	// If such "does not exist" i.e., first timer, return -1.
	findWhenLastSaid := func(ind, num int) int {
		if vv, ok := whenLastSaid[num]; ok && vv <= ind-2 {
			return vv
		}
		return -1
	}

	// bar := progressbar.Default(int64(len(memory)))
	for ind < len(memory) {
		if ff := findWhenLastSaid(ind, lastSpoken); ff == -1 {
			memory[ind] = 0
		} else {
			memory[ind] = ind - 1 - ff
		}
		whenLastSaid[memory[ind-1]] = ind - 1 // memory[ind-1] used to avoid writing the current no. for the first time
		lastSpoken = memory[ind]
		ind++

		if ind%10000 == 0 {
			// bar.Add(10000)
		}
	}
	// bar.Finish()

	answer1 = strconv.Itoa(memory[2019])
	answer2 = strconv.Itoa(memory[29999999])

	return
}

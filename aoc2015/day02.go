package aoc2015

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Day02 solves the second day puzzle "I Was Told There Would Be No Math".
//
// Input
//
// A file containing 1000 lines, each of which is in the format LxWxH,
// where L, W, and H are positive integers no more than 32,
// and x is the character 'x'. For example:
//
//	7x4x10
//	22x24x29
//	30x1x2
//	19x2x5
//	11x9x22
//	23x15x10
//	11x11x10
//	30x28x5
//	22x5x4
//	6x26x20
//	16x12x30
//	10x20x5
//	25x14x24
//	16x17x22
//	11x28x26
//	1x11x10
//
// If there is more than or less than three numbers in a row,
// or if there are issues with parsing the input,
// the function will return an error corresponding to the problematic line.
// If the numbers are too large, that may result in an integer overflow.
func Day02(input string) (answer1, answer2 string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	// initialize important variables
	splitDims := make([]string, 3)
	actualDims := make([]int, 3)
	var totalPaper, totalRibbon int // part 1 and part 2 respectively

	// some helper functions
	// it is assumed that dims is of length 3,
	// since we first check whether that is the case.

	// max function for maximum of a slice
	max := func(list []int) int {
		currentMax := list[0]
		for i := range list {
			if list[i] > currentMax {
				currentMax = list[i]
			}
		}
		return currentMax
	}

	// surfaceArea function for total surface area (part 1)
	surfaceArea := func(dims []int) int {
		return 2 * (dims[0]*dims[1] + dims[1]*dims[2] + dims[2]*dims[0])
	}

	// slack function for the little extra paper (part 1)
	slack := func(dims []int) int {
		result := 1
		for _, v := range dims {
			result *= v
		}
		result /= max(dims)
		return result
	}

	// ribbon for amount of ribbon to be used (part 2)
	ribbon := func(dims []int) int {
		result := 0
		for _, v := range dims {
			result += 2 * v
		}
		result -= 2 * max(dims)
		return result
	}
	// bow function for total volume (part 2)
	bow := func(dims []int) int {
		return dims[0] * dims[1] * dims[2]
	}

	// now for the big part.
	for scanner.Scan() {
		readDims := scanner.Text()
		splitDims = strings.Split(readDims, "x")
		if len(splitDims) != 3 {
			// throw an error
			err = fmt.Errorf("invalid dimensions: %v", readDims)
			return
		}
		// Atoi each dimension
		for ind := range actualDims {
			convt, e := strconv.Atoi(splitDims[ind])
			if e != nil {
				err = fmt.Errorf("not an integer: %v (%v)", readDims, splitDims[ind])
				return
			}
			actualDims[ind] = convt
		}
		totalPaper += surfaceArea(actualDims) + slack(actualDims)
		totalRibbon += ribbon(actualDims) + bow(actualDims)
	}

	answer1 = strconv.Itoa(totalPaper)
	answer2 = strconv.Itoa(totalRibbon)
	return
}

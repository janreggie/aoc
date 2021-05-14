package aoc2020

import (
	"bufio"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Day01 solves the first day puzzle "Report Repair".
//
// Input
//
// A file of 200 lines, each line consisting of an integer. For example:
//
//  1721
//  979
//  366
//  299
//  675
//  1456
//
// It is guaranteed that no number is greater than 2020.
func Day01(input string) (answer1, answer2 string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Grab input
	inputs := make([]int, 0, 200) // prealloc 200 because average case
	for scanner.Scan() {
		current, e := strconv.Atoi(scanner.Text())
		if e != nil {
			err = errors.Wrapf(e, "couldn't parse line %v", scanner.Text())
			return
		}
		inputs = append(inputs, current)
	}

	// We can mutate the input (since order ultimately wouldn't matter)
	// and sort the numbers we have.
	// From there, we can sweep from the beginning and from the end,
	// going through the center slowly, until we have a sum of 2020.
	sort.Ints(inputs)
	ii, jj := 0, len(inputs)-1
	for ii < jj {
		sum := inputs[ii] + inputs[jj]
		if sum == 2020 {
			answer1 = strconv.Itoa(inputs[ii] * inputs[jj])
			break
		}
		if sum < 2020 {
			ii++
			continue
		}
		if sum > 2020 {
			jj--
			continue
		}
	}
	if ii == jj {
		err = errors.Errorf("couldn't find two numbers with sum 2020 from %v", inputs)
		return
	}

	// Now for the three number case, which will be a lot more difficult...
	// I doubt there's a better solution than a O(n^3) one.
L1:
	for ii := 0; ii < len(inputs)-2; ii++ {
		for jj := 1; jj < len(inputs)-1; jj++ {
			if inputs[ii]+inputs[jj] > 2020 {
				break // continue with next value for ii
			}

			for kk := 2; kk < len(inputs); kk++ {
				sum := inputs[ii] + inputs[jj] + inputs[kk]
				if sum == 2020 {
					answer2 = strconv.Itoa(inputs[ii] * inputs[jj] * inputs[kk])
					break L1
				}
				if sum > 2020 {
					break // continue with next value for jj
				}
			}
		}
	}
	if answer2 == "" {
		err = errors.Errorf("couldn't find three numbers with sum 2020 from %v", inputs)
		return
	}

	return
}

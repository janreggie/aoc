package aoc2020

import (
	"sort"
	"strconv"

	aoc "github.com/janreggie/aoc/internal"
)

// Day09 solves the ninth day puzzle "Encoding Error"
//
// Input
//
// A file containing positive integers separated by newlines. For example:
//
//   35
//   20
//   15
//   25
//   47
//   40
//   62
//   55
//   65
//   95
//   102
//   117
//   150
//   182
//   127
//
// It is guaranteed that the input is at least 26 lines long.
func Day09(input string) (answer1, answer2 string, err error) {

	allNumbers, err := aoc.SplitLinesToInts(input)
	if err != nil {
		return
	}

	// checkValidSumOfTwo returns true if there are two (distinct) numbers in past that sum to present.
	// This will be for later...
	checkValidSumOfTwo := func(past []int, present int) bool {
		pastSorted := make([]int, len(past))
		copy(pastSorted, past)
		sort.Ints(pastSorted)

		// Consider that there are two indices, ii starting from 0, and jj starting from len-1.
		ii, jj := 0, len(pastSorted)-1
		for ii != jj {
			if sum := pastSorted[ii] + pastSorted[jj]; sum == present {
				return true
			} else if sum < present {
				ii++
			} else if sum > present {
				jj--
			}
		}
		return false
	}

	invalidNumber := 0
	for ii := 0; ii < len(allNumbers)-26; ii++ {
		past := allNumbers[ii : ii+25]
		present := allNumbers[ii+25]
		if !checkValidSumOfTwo(past, present) {
			invalidNumber = present
			break
		}
	}
	answer1 = strconv.Itoa(invalidNumber)

	// Now for the second part. Declare first some helper fxns for later
	sum := func(ss []int) int {
		result := 0
		for _, vv := range ss {
			result += vv
		}
		return result
	}
	min := func(ss []int) int {
		result := ss[0]
		for _, vv := range ss {
			if vv < result {
				result = vv
			}
		}
		return result
	}
	max := func(ss []int) int {
		result := ss[0]
		for _, vv := range ss {
			if vv > result {
				result = vv
			}
		}
		return result
	}

	//Suppose two indices ii and jj exist
	ii, jj := 0, 25
	// and now the sum of allNumbers[ii:jj] is determined.
	// If the number is less than invalidNumber, increase the range by incrementing jj.
	// If greater than, decrease the range by incrementing ii.
	encryptionWeakness := 0
	for ii != jj && jj <= len(allNumbers) {
		rr := allNumbers[ii:jj]
		ss := sum(rr)
		// fmt.Printf("Checking range %d to %d which sum to %d\n", ii, jj, ss)
		if ss == invalidNumber {
			// Note jj-1 since that is the upper bound of the slice
			encryptionWeakness = min(rr) + max(rr)
			break
		} else if ss < invalidNumber {
			jj++
		} else if ss > invalidNumber {
			ii++
		}
	}
	answer2 = strconv.Itoa(encryptionWeakness)

	return
}

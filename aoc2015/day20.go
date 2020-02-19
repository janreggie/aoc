package aoc2015

import (
	"bufio"
	"math"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

// divisors represents the divisors of a number
type divisors []uint

// newDivisors determines the divisors of a number
func newDivisors(number uint) divisors {
	if number == 1 {
		return []uint{1}
	}
	result := make(divisors, 2)
	result[0], result[1] = 1, number
	for ii := uint(2); ii <= uint(math.Sqrt(float64(number))); ii++ {
		if number%ii == 0 {
			result = append(result, ii)
			if ii != number/ii {
				result = append(result, number/ii)
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

// sum returns the sum of divisors
func (dv divisors) sum() uint {
	var result uint
	for _, vv := range dv {
		result += vv
	}
	return result
}

// count returns the number of divisors
func (dv divisors) count() uint {
	return uint(len(dv))
}

// onlyFifty returns a subset of divisors that are less than the number over fifty
func (dv divisors) onlyFifty() divisors {
	result := make(divisors, 0)
	number := dv[len(dv)-1]
	for _, vv := range dv {
		if vv*50 >= number {
			result = append(result, vv)
		}
	}
	return result
}

// Day20 solves the twentieth day puzzle "Infinite Elves and Infinite Houses".
//
// Input
//
// A single line describing a number of presents. For example:
//
//	29000000
//
// It is guaranteed that the input is a non-negative int32.
func Day20(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	scanner.Scan()
	inputRaw, err := strconv.ParseUint(scanner.Text(), 10, 32)
	if err != nil {
		err = errors.Wrapf(err, "could not read inpout %v", scanner.Text())
	}
	input := uint(inputRaw)

	// second half is not implemented yet.
	wroteFirst, wroteSecond := false, false
	for house := uint(1); !wroteFirst || !wroteSecond; house++ {
		// factors := factor(house)
		divisors := newDivisors(house)
		if !wroteFirst && divisors.sum()*10 >= input {
			answer1 = strconv.FormatUint(uint64(house), 10)
			wroteFirst = true
		}
		if !wroteSecond && divisors.onlyFifty().sum()*11 >= input {
			answer2 = strconv.FormatUint(uint64(house), 10)
			wroteSecond = true
		}
	}

	return
}

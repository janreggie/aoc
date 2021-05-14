package aoc2015

import (
	"math"
	"sort"
	"strconv"
)

// divisors represents the divisors of a number
type divisors []uint64

// newDivisors determines the divisors of a number
func newDivisors(number uint64) divisors {
	if number == 1 {
		return []uint64{1}
	}
	result := make(divisors, 2)
	result[0], result[1] = 1, number
	for ii := uint64(2); ii <= uint64(math.Sqrt(float64(number))); ii++ {
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
func (dv divisors) sum() uint64 {
	var result uint64
	for _, vv := range dv {
		result += vv
	}
	return result
}

// count returns the number of divisors
func (dv divisors) count() uint64 {
	return uint64(len(dv))
}

// last returns the final number of divisors, which should be the number pertaining to the divisors
func (dv divisors) last() uint64 {
	return dv[dv.count()-1]
}

// onlyFifty returns a subset of divisors that are greater than the number
// when multiplied by fifty.
// It is guaranteed that dv is sorted.
func (dv divisors) onlyFifty() divisors {
	// result := make(divisors, 0)
	// number := dv[len(dv)-1]
	// for _, vv := range dv {
	// 	if vv*50 >= number {
	// 		result = append(result, vv)
	// 	}
	// }
	// return result
	ind := 0
	basis := dv[len(dv)-1]
	for ; ind < len(dv); ind++ {
		if dv[ind]*50 >= basis {
			break
		}
	}
	// dv[ind:] contains all what you'd want
	result := make(divisors, len(dv[ind:]))
	copy(result, dv[ind:])
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
func Day20(input string) (answer1, answer2 string, err error) {

	realInput, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return
	}

	divisorChan := make(chan divisors, 10)
	done := make(chan struct{})

	// iterator for house counts
	iterator := make(chan uint64, 10)
	go func() {
		for ii := uint64(1); ; ii++ {
			select {
			case <-done:
				return
			default:
				iterator <- ii
			}
		}
	}()

	// generate divisors in parallel
	for ii := 0; ii < 8; ii++ {
		go func() {
			for house := range iterator {
				divisorChan <- newDivisors(house)
			}
		}()
	}

	wroteFirst, wroteSecond := false, false
	for !wroteFirst || !wroteSecond {
		divisors := <-divisorChan
		if !wroteFirst && divisors.sum()*10 >= realInput {
			answer1 = strconv.FormatUint(uint64(divisors.last()), 10)
			wroteFirst = true
		}
		if !wroteSecond && divisors.onlyFifty().sum()*11 >= realInput {
			answer2 = strconv.FormatUint(uint64(divisors.last()), 10)
			wroteSecond = true
		}
	}
	done <- struct{}{}

	return
}

package aoc2015

import (
	"bufio"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

// factors represents the factorization of a number
type factors map[uint]uint // prime factors and their powers

// factor factors a number.
func factor(number uint) factors {
	out := make(factors)
	largestPossible := uint(math.Sqrt(float64(number)))
	for current := uint(2); current < largestPossible; current++ {
		for number%current == 0 {
			number /= current
			out[current]++
		}
	}
	if number != 1 {
		out[number]++ // whatever this could be
	}
	return out
}

// sumOfDivisors returns the sum of divisors of a factorization
func (factors factors) sumOfDivisors() uint {
	var out uint = 1
	for prime, power := range factors {
		var sumOnPrime uint = 1
		var primeToii uint = 1
		for ii := uint(0); ii < power; ii++ {
			primeToii *= prime
			sumOnPrime += primeToii
		}
		out *= sumOnPrime
	}
	return out
}

// countDivisors counts the number of divisors of a factorization
func (factors factors) countDivisors() uint {
	var out uint = 1
	for _, power := range factors {
		out *= (power + 1)
	}
	return out
}

// Day20 solves the twentieth day puzzle
// "Infinite Elves and Infinite Houses"
func Day20(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	scanner.Scan()
	inputRaw, err := strconv.ParseUint(scanner.Text(), 10, 32)
	if err != nil {
		err = errors.Wrapf(err, "could not read inpout %v", scanner.Text())
	}
	input := uint(inputRaw)

	// second half is not implemented yet.
	wroteFirst, wroteSecond := false, true
	for house := uint(1); !wroteFirst || !wroteSecond; house++ {
		factors := factor(house)
		if !wroteFirst && factors.sumOfDivisors()*10 >= input {
			answer1 = strconv.FormatUint(uint64(house), 10)
			wroteFirst = true
		}
		// if !wroteSecond && factors.countDivisors() < 50 && factors.sumOfDivisors()*11 >= input {
		// 	answer2 = strconv.FormatUint(uint64(house), 10)
		// 	wroteSecond = true
		// }
	}

	return
}

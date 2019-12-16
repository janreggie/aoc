package aoc2019

import (
	"bufio"
	"strconv"
)

// Day01 solves the first day puzzle
// "The Tyranny of the Rocket Equation"
func Day01(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	var totalFuel int64 // total fuel required
	var cumuFuel int64  // answer2 (where fuel requires fuel!)
	massToFuel := func(mass int64) int64 {
		return (mass / 3) - 2 // will automatically round down!!
	}
	for scanner.Scan() {
		raw := scanner.Text()
		mass, e := strconv.ParseInt(raw, 10, 64)
		if e != nil {
			err = e
			return
		}
		totalFuel += massToFuel(mass)
		for fuel := massToFuel(mass); fuel > 0; fuel = massToFuel(fuel) {
			cumuFuel += fuel
		}
	}

	answer1 = strconv.FormatInt(totalFuel, 10)
	answer2 = strconv.FormatInt(cumuFuel, 10)
	return
}

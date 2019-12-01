package aoc2019

import (
	"bufio"
	"strconv"
)

func day01MassToFuel(mass int64) int64 {
	return (mass / 3) - 2 // will automatically round down!!
}

// Day01 solves the first day puzzle
// "The Tyranny of the Rocket Equation"
func Day01(scanner *bufio.Scanner) (answer1, answer2 string, e error) {
	var totalFuel int64 // total fuel required
	var cumuFuel int64  // answer2 (where fuel requires fuel!)
	for scanner.Scan() {
		raw := scanner.Text()
		mass, err := strconv.Atoi(raw)
		if err != nil {
			e = err
			return
		}
		totalFuel += day01MassToFuel(int64(mass))
		for fuel := day01MassToFuel(int64(mass)); fuel > 0; fuel = day01MassToFuel(fuel) {
			cumuFuel += fuel
		}
	}

	answer1 = strconv.FormatInt(totalFuel, 10)
	answer2 = strconv.FormatInt(cumuFuel, 10)
	return
}

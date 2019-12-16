package aoc2015

import (
	"bufio"
	"strconv"
)

// Day03 solves the third day puzzle
// "Perfectly Spherical Houses in a Vacuum"
func Day03(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	scanner.Split(bufio.ScanBytes)
	iterateCoordinate := func(coordX, coordY *int32, input string) {
		switch input {
		case "v":
			*coordY--
		case "^":
			*coordY++
		case "<":
			*coordX--
		case ">":
			*coordX++
		}
	}

	// okay how do we "store" these? maybe hash map
	type coord struct {
		x, y int32
	}
	// problem 1
	var currentX, currentY int32
	allResults := make(map[coord]int32) // allResults[coord] = no. of times it occured
	allResults[coord{currentX, currentY}]++
	// problem 2
	var santaX, santaY, roboX, roboY int32
	dualResults := make(map[coord]int32)
	dualResults[coord{santaX, santaY}]++
	isSanta := true // is it santa's turn? will invert for every iteration

	for scanner.Scan() {
		readDirection := scanner.Text()
		iterateCoordinate(&currentX, &currentY, readDirection)
		allResults[coord{currentX, currentY}]++
		if isSanta {
			iterateCoordinate(&santaX, &santaY, readDirection)
			dualResults[coord{santaX, santaY}]++
		} else {
			iterateCoordinate(&roboX, &roboY, readDirection)
			dualResults[coord{roboX, roboY}]++
		}
		isSanta = !isSanta
		if err = scanner.Err(); err != nil {
			return
		}
	}
	answer1 = strconv.Itoa(len(allResults))
	answer2 = strconv.Itoa(len(dualResults))

	return
}

package aoc2015

import (
	"bufio"
	"image"
	"strconv"
)

// iterateCoordinate moves the location using an input (`v`, `^`, `<`, `>`)
func iterateCoordinate(location *image.Point, input string) {
	switch input {
	case "v":
		location.Y--
	case "^":
		location.Y++
	case "<":
		location.X--
	case ">":
		location.X++
	}
}

// Day03 solves the third day puzzle "Perfectly Spherical Houses in a Vacuum".
//
// Input
//
// A single line of length 8192 containing only the characters '<', '>', '^', and 'v',
// each representing direction. For example:
//
//	v^>>v>v>>>>^^>v<^v^>><<^<>>v><^><<^>^<vv^^<><<>><vvvv^>^^<><^^
//
// Any characters that are not any of those will be ignored.
func Day03(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	scanner.Split(bufio.ScanBytes)

	// okay how do we "store" these? maybe hash map
	// part 1
	var santa image.Point               // santa's location (part one)
	houses := make(map[image.Point]int) // houses[coord] = no. of times we've passed by a point
	houses[santa]++
	// part 2
	var nuSanta, roboSanta image.Point      // nuSanta and roboSanta's locations
	dualHouses := make(map[image.Point]int) // dualHouses is houses but for part 2
	dualHouses[nuSanta]++
	isSanta := true // is it nuSanta's turn? will invert for every iteration

	for scanner.Scan() {
		readDirection := scanner.Text()
		iterateCoordinate(&santa, readDirection)
		houses[santa]++
		if isSanta {
			iterateCoordinate(&nuSanta, readDirection)
			dualHouses[nuSanta]++
		} else {
			iterateCoordinate(&roboSanta, readDirection)
			dualHouses[roboSanta]++
		}
		isSanta = !isSanta
	}
	answer1 = strconv.Itoa(len(houses))
	answer2 = strconv.Itoa(len(dualHouses))

	return
}

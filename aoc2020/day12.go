package aoc2020

import (
	"strconv"
	"strings"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/pkg/errors"
)

// navigationInstruction represents an action of navigation
type navigationInstruction struct {
	actn byte // E, S, W, N, L, R, F
	val  int  // by how much was actn executed..?
}

// newNavigationInstruction returns a new navigation instruction with some string input.
// If invalid, return an error
func newNavigationInstruction(input string) (navigationInstruction, error) {
	instr := navigationInstruction{}
	if len(input) < 2 {
		return instr, errors.Errorf("input string too short %s", input)
	}

	instr.actn = input[0]
	if !strings.ContainsRune("ESWNLRF", rune(instr.actn)) {
		return instr, errors.Errorf("could not decode invalid action %c", instr.actn)
	}

	val, err := strconv.Atoi(input[1:])
	if err != nil {
		return instr, errors.Wrapf(err, "could not parse value of input %s", input)
	}
	if (instr.actn == 'L' || instr.actn == 'R') && val%90 != 0 {
		return instr, errors.Errorf("actn was %c and val is %d, not multiple of 90", instr.actn, val)
	}
	instr.val = val
	return instr, nil
}

// generateNavigationInstructions generates a slice of navigationInstruction objects from a Scanner input
func generateNavigationInstructions(input string) ([]navigationInstruction, error) {
	instrs := make([]navigationInstruction, 0)
	for _, line := range aoc.SplitLines(input) {
		instr, err := newNavigationInstruction(line)
		if err != nil {
			return nil, errors.Wrapf(err, "could not parse navigation instruction %s", line)
		}
		instrs = append(instrs, instr)
	}
	return instrs, nil
}

// ship represents a ship. The zero value is ready to use.
// Note that you might have to set up a waypoint position
type ship struct {
	x, y      int // position at x and y
	direction int
	wx, wy    int // waypoint position
}

// directions for ship (increasing: clockwise)
const (
	east int = iota
	south
	west
	north
)

// readInstruction reads a navigationInstruction and the ship moves accordingly (part 1).
// If instr were to be invalid, do nothing.
func (s *ship) readInstruction(instr navigationInstruction) {
	switch instr.actn {
	case 'E':
		s.x += instr.val
	case 'S':
		s.y -= instr.val
	case 'W':
		s.x -= instr.val
	case 'N':
		s.y += instr.val
	case 'L':
		if instr.val%90 != 0 {
			return
		}
		// `4 +` is important to ensure that argument, and hence result, stays positive.
		s.direction = (4 + s.direction - instr.val/90) % 4
	case 'R':
		if instr.val%90 != 0 {
			return
		}
		s.direction = (4 + s.direction + instr.val/90) % 4
	case 'F':
		switch s.direction {
		case east:
			s.x += instr.val
		case south:
			s.y -= instr.val
		case west:
			s.x -= instr.val
		case north:
			s.y += instr.val
		}
	}
}

// readWaypointy reads a navigationInstruction and the ship moves using the waypoint system (part 2).
// If instr were to be invalid, do nothing.
func (s *ship) readWaypointy(instr navigationInstruction) {
	switch instr.actn {
	case 'E':
		s.wx += instr.val
	case 'S':
		s.wy -= instr.val
	case 'W':
		s.wx -= instr.val
	case 'N':
		s.wy += instr.val
	case 'L':
		// might be a bit bad practiceregarding code duplication...
		switch instr.val {
		case 90:
			s.wx, s.wy = -s.wy, s.wx
		case 180:
			s.wx, s.wy = -s.wx, -s.wy
		case 270:
			s.wx, s.wy = s.wy, -s.wx
		}
	case 'R':
		// might be a bit bad practiceregarding code duplication...
		switch instr.val {
		case 90:
			s.wx, s.wy = s.wy, -s.wx
		case 180:
			s.wx, s.wy = -s.wx, -s.wy
		case 270:
			s.wx, s.wy = -s.wy, s.wx
		}
	case 'F':
		s.x += s.wx * instr.val
		s.y += s.wy * instr.val
	}
}

// manhattanDistance returns the Manhattan Distance from the origin to the ship's location
func (s *ship) manhattanDistance() int {
	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}
	return abs(s.x) + abs(s.y)
}

// Day12 solves the twelfth day puzzle "Rain Risk"
//
// Input
//
// A list of commands for a "ship" to execute. For example:
//
//   F10
//   N3
//   F7
//   R90
//   F11
//
// The actions will either be N, S, E, W, L, R, and F;
// and the values are guaranteed to be positive numbers no more than 300.
// The values for L and R will be multiples of 90 (90, 180, 270).
func Day12(input string) (answer1, answer2 string, err error) {

	instrs, err := generateNavigationInstructions(input)
	if err != nil {
		errors.Wrapf(err, "could not generate navigation instructions from input")
		return
	}
	ss := ship{}
	for _, instr := range instrs {
		ss.readInstruction(instr)
	}
	answer1 = strconv.Itoa(ss.manhattanDistance())

	// now for part 2
	ss.wx, ss.wy = 10, 1
	ss.x, ss.y = 0, 0
	ss.direction = east
	for _, instr := range instrs {
		ss.readWaypointy(instr)
	}
	answer2 = strconv.Itoa(ss.manhattanDistance())

	return
}

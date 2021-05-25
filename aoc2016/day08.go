package aoc2016

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/pkg/errors"
)

// littleScreen is a structure representing the screen in Year 2016 Day 8.
// The zero value is ready to use.
type littleScreen struct {
	//lcd [50][6]bool
	lcd [6][50]bool
}

// func newLittleScreen() littleScreen {
// 	return littleScreen{}
// }

func (ls littleScreen) String() string {
	// print out lcd
	var sb strings.Builder
	for ii := 0; ii < 6; ii++ {
		for jj := 0; jj < 50; jj++ {
			if ls.lcd[ii][jj] {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

// rect turns on all pixels in a rectangle at the top-left of the screen
// which is `width` pixels wide and `height` tall.
// Returns an error if width > 50 or height > 6.
func (ls *littleScreen) rect(width, height uint) error {
	// do things....
	if width > 50 || height > 6 {
		return littleScreenError{"rect", width, height}
	}

	for ii := uint(0); ii < height; ii++ {
		for jj := uint(0); jj < width; jj++ {
			ls.lcd[ii][jj] = true
		}
	}
	return nil
}

// rotateRow shifts all pixels in row `row` rightwards by `offset` pixels.
// Returns an error if row >= 6.
func (ls *littleScreen) rotateRow(row, offset uint) error {
	if row >= 6 {
		return littleScreenError{"rotateRow", row, offset}
	}

	offset = offset % 50 // beyond 50 the effect will wrap back to 0
	lastPixels := make([]bool, offset)
	// set last offset number of pixels in ls.lcd(row=row) to lastPixels
	for ii := uint(0); ii < offset; ii++ {
		lastPixels[ii] = ls.lcd[row][50-offset+ii]
	}
	for ii := uint(0); ii < 50-offset; ii++ {
		ls.lcd[row][49-ii] = ls.lcd[row][49-(offset+ii)]
	}
	// get lastPixels back to ls.lcd
	for ii := uint(0); ii < offset; ii++ {
		ls.lcd[row][ii] = lastPixels[ii]
	}

	return nil
}

// rotateColumn shifts all pixels in column `col` downwards by `offset` pixels.
// Returns an error if col > 50.
func (ls *littleScreen) rotateColumn(col, offset uint) error {
	if col >= 50 {
		return littleScreenError{"rotateColumn", col, offset}
	}

	offset = offset % 6
	lastPixels := make([]bool, offset)
	// set last offset number of pixels in ls.lcd(col=col) to lastPixels
	for ii := uint(0); ii < offset; ii++ {
		lastPixels[ii] = ls.lcd[6-offset+ii][col]
	}
	for ii := uint(0); ii < 6-offset; ii++ {
		ls.lcd[5-ii][col] = ls.lcd[5-(offset+ii)][col]
	}
	// get lastPixels back to ls.lcd
	for ii := uint(0); ii < offset; ii++ {
		ls.lcd[ii][col] = lastPixels[ii]
	}

	return nil
}

// countLit counts the number of lights are lit in the LED
func (ls littleScreen) countLit() uint {
	var result uint
	for ii := 0; ii < 6; ii++ {
		for jj := 0; jj < 50; jj++ {
			if ls.lcd[ii][jj] {
				result++
			}
		}
	}
	return result
}

// parseInstruction parses an instruction.
// Examples of instruction are "rect AxB",
// "rotate row y=A by B", or "rotate column x=A by B".
func (ls *littleScreen) parseInstruction(instruction string) error {
	cantParse := fmt.Errorf("could not parse %v", instruction)

	// if len(instruction) < 8 { // less than `rotate ...` or `rect ?x?`
	// 	return errors.Wrap(cantParse, "instruction too short")
	// }

	if strings.HasPrefix(instruction, "rect ") {
		dims := strings.Split(strings.TrimPrefix(instruction, "rect "), "x")
		// did uint64 because either way if the numbers are too big
		// they'll get error'd anyway
		if len(dims) != 2 {
			return errors.Wrapf(cantParse, "could not parse dimensions %v", dims)
		}
		w64, err := strconv.ParseUint(dims[0], 10, 64)
		if err != nil {
			return errors.Wrapf(cantParse, "could not parse width %v: %v", dims[0], err)
		}
		h64, err := strconv.ParseUint(dims[1], 10, 64)
		if err != nil {
			return errors.Wrapf(cantParse, "could not parse height %v: %v", dims[0], err)
		}
		return ls.rect(uint(w64), uint(h64))
	}

	if strings.HasPrefix(instruction, "rotate ") {
		remains := strings.TrimPrefix(instruction, "rotate ")
		var rotRow bool
		var dims []string
		if strings.HasPrefix(remains, "row y=") {
			dims = strings.Split(strings.TrimPrefix(remains, "row y="), " by ")
			rotRow = true
		} else if strings.HasPrefix(remains, "column x=") {
			dims = strings.Split(strings.TrimPrefix(remains, "column x="), " by ")
			rotRow = false
		} else {
			return errors.Wrapf(cantParse, "unknown rotate method %v", remains)
		}

		if len(dims) != 2 {
			return errors.Wrapf(cantParse, "unknown dimensions %v", dims)
		}
		paramA, err := strconv.ParseUint(dims[0], 10, 64)
		if err != nil {
			return errors.Wrapf(cantParse, "could not parse first parameter %v", dims[0])
		}
		paramB, err := strconv.ParseUint(dims[1], 10, 64)
		if err != nil {
			return errors.Wrapf(cantParse, "could not parse second parameter %v", dims[1])
		}
		if rotRow {
			return ls.rotateRow(uint(paramA), uint(paramB))
		}
		return ls.rotateColumn(uint(paramA), uint(paramB))
	}

	return cantParse
}

// littleScreenError is an error from manipulating a littleScreen
type littleScreenError struct {
	function       string // rect, rotateRow, rotateColumn
	paramA, paramB uint   // parameters
}

func (e littleScreenError) Error() string {
	return fmt.Sprintf("could not execute %v(%v,%v)", e.function, e.paramA, e.paramB)
}

// Day08 solves the eight day puzzle "Two-Factor Authentication".
//
// Input
//
// A series of card instructions that look like the following:
//
//	rect AxB
//	rotate row y=A by B
//	rotate column x=A by B
//
// An invalid instruction will cause an error.
func Day08(input string) (answer1, answer2 string, err error) {

	var ls littleScreen
	for _, line := range aoc.SplitLines(input) {
		e := ls.parseInstruction(line)
		if e != nil {
			err = errors.Wrapf(e, "could not parse instruction %v", line)
			return
		}
	}
	answer1 = strconv.FormatUint(uint64(ls.countLit()), 10)
	answer2 = "this:\n" + ls.String()
	return
}

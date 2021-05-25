package aoc2020

import (
	"strconv"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/pkg/errors"
)

// tobogganForest represents a Toboggan forest
type tobogganForest struct {
	width  int      // width of the forest, after which, repeats itself
	height int      // number of rows the forest has
	forest [][]bool // true if a tree exists here
}

// readTobogganForest reads to forest data.
// If it is uneven, reallocate...
func readTobogganForest(input string) (*tobogganForest, error) {

	result := &tobogganForest{}
	result.forest = make([][]bool, 0)

	for _, line := range aoc.SplitLines(input) {
		currentRow := make([]bool, 0)
		for _, unit := range line {
			switch unit {
			case '#':
				currentRow = append(currentRow, true)
			case '.':
				currentRow = append(currentRow, false)
			default:
				return nil, errors.Errorf("couldn't read char %c at row `%s`", unit, line)
			}
		}

		result.height++
		if result.width != 0 && result.width != len(currentRow) {
			return nil, errors.Errorf("width should be %v but row %s is length %v", result.width, line, len(currentRow))
		}
		result.width = len(currentRow)
		result.forest = append(result.forest, currentRow)
	}

	return result, nil
}

// read reads at coordinates x, y where x is distance from first column and y distance from first row
// and returns whether a tree is found there.
// x, y start at zero.
// If y is out of bounds or x is negative, return False.
func (forest *tobogganForest) read(x, y int) bool {
	if x < 0 {
		return false
	}
	if y < 0 || y >= forest.height {
		return false
	}

	// obvious check if width or height are empty
	if forest.width == 0 || forest.height == 0 {
		return false
	}

	return forest.forest[y][x%forest.width]
}

// traverse traverses through the forest starting at (x,y) by taking r steps to the right and d steps downwards
// for every turn, and ultimately counting how many trees have been encountered.
func (forest *tobogganForest) traverse(r, d, x, y int) int {
	if r == 0 && d == 0 { // because this will make traverse run indefinitely
		if forest.read(x, y) {
			return 1
		}
		return 0
	}

	if y > forest.height || y < 0 || x < 0 {
		return 0
	}
	if forest.read(x, y) {
		return 1 + forest.traverse(r, d, x+r, y+d)
	}
	return forest.traverse(r, d, x+r, y+d)
}

// Day03 solves the third day puzzle "Toboggan Trajectory"
//
// Input
//
// A file which contains a description of a forest expressed in dots and hashes. For example:
//
//  some sample input indented to become a code block
//
// The file should only contain `.`, `#`, and newlines. All lines are of equal length.
func Day03(input string) (answer1, answer2 string, err error) {

	forest, err := readTobogganForest(input)
	if err != nil {
		err = errors.Wrap(err, "couldn't read from scanner")
		return
	}

	// While this can be solved concurrently,
	// the overhead of using channels far exceeds the runtime of the five traverse calls below,
	// hence doing so might not be the wisest thing to do.
	r1d1 := forest.traverse(1, 1, 0, 0)
	r3d1 := forest.traverse(3, 1, 0, 0)
	r5d1 := forest.traverse(5, 1, 0, 0)
	r7d1 := forest.traverse(7, 1, 0, 0)
	r1d2 := forest.traverse(1, 2, 0, 0)

	answer1 = strconv.Itoa(r3d1)
	answer2 = strconv.Itoa(r1d1 * r3d1 * r5d1 * r7d1 * r1d2)
	return
}

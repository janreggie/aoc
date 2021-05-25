package aoc2020

import (
	"bufio"
	"strings"
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestTobogganForest(t *testing.T) {
	assert := assert.New(t)

	forest, err := readTobogganForest(bufio.NewScanner(strings.NewReader(day03sampleInput)))
	assert.NoErrorf(err, "no error should be from input")

	// check the upper right corner
	assert.EqualValues(false, forest.read(0, 0))
	assert.EqualValues(true, forest.read(0, 1))
	assert.EqualValues(false, forest.read(1, 0))
	assert.EqualValues(false, forest.read(1, 1))

	// check beyond the bounds
	assert.EqualValues(false, forest.read(11, 0))
	assert.EqualValues(true, forest.read(11, 1))
	assert.EqualValues(false, forest.read(12, 0))
	assert.EqualValues(false, forest.read(12, 1))

	// check traverse
	assert.EqualValues(2, forest.traverse(1, 1, 0, 0))
	assert.EqualValues(7, forest.traverse(3, 1, 0, 0))
	assert.EqualValues(3, forest.traverse(5, 1, 0, 0))
	assert.EqualValues(4, forest.traverse(7, 1, 0, 0))
	assert.EqualValues(2, forest.traverse(1, 2, 0, 0))
}

func TestDay03(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{
			Details: "Y2020D03 sample input",
			Input:   day03sampleInput,
			Result1: "7",
			Result2: "336",
		},
		{
			Details: "Y2020D03 my input",
			Input:   day03myInput,
			Result1: "200",
			Result2: "3737923200",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day03, assert)
	}
}

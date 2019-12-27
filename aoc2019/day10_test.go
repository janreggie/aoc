package aoc2019

import (
	"testing"

	"github.com/janreggie/AdventOfCode/structs"
	"github.com/stretchr/testify/assert"
)

func TestDay10(t *testing.T) {
	assert := assert.New(t)
	testCases := []structs.TestCase{
		{Input: "......#.#.\n" +
			"#..#.#....\n" +
			"..#######.\n" +
			".#.#.###..\n" +
			".#..#.....\n" +
			"..#....#.#\n" +
			"#..#....#.\n" +
			".##.#..###\n" +
			"##...#..#.\n" +
			".#....####\n",
			Result1: "33",
			Result2: ""},
	}
	for _, eachCase := range testCases {
		eachCase.Test(Day10, assert)
	}
}

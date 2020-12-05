package aoc2019

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay03(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Input: "R8,U5,L5,D3\nU7,R6,D4,L4",
			Result1: "6",
			Result2: "30"},
		{Input: "R75,D30,R83,U83,L12,D49,R71,U7,L72\n" +
			"U62,R66,U55,R34,D71,R55,D58,R83",
			Result1: "159",
			Result2: "610"},
		{Input: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\n" +
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			Result1: "135",
			Result2: "410"},
		{Details: "Y2019D03 my input",
			Input:   day03myInput,
			Result1: "1431",
			Result2: "48012"},
	}
	for _, tt := range testCases {
		tt.Test(Day03, assert)
	}
}

func BenchmarkDay03(b *testing.B) {
	internal.Benchmark(Day03, b, day03myInput)
}

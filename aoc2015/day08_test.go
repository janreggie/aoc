package aoc2015

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay08(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Input: `""`,
			Result1: "2",
			Result2: "4"},
		{Input: `"abc"`,
			Result1: "2",
			Result2: "4"},
		{Input: `"aaa\"aaa"`,
			Result1: "3",
			Result2: "6"},
		{Input: `"\x27"`,
			Result1: "5",
			Result2: "5"},
		{Details: "Y2015D08 my input",
			Input:   day08myInput,
			Result1: "1342",
			Result2: "2074"},
	}
	for _, tt := range testCases {
		tt.Test(Day08, assert)
	}
}

func BenchmarkDay08(b *testing.B) {
	aoc.Benchmark(Day08, b, day08myInput)
}

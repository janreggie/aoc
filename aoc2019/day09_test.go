package aoc2019

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay09(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Details: "Y2019D09 my input",
			Input:   day09myInput,
			Result1: "3345854957",
			Result2: "68938"},
	}
	for _, tt := range testCases {
		tt.Test(Day09, assert)
	}
}

func BenchmarkDay09(b *testing.B) {
	aoc.Benchmark(Day09, b, day09myInput)
}

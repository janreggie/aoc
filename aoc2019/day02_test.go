package aoc2019

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay02(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Details: "Y2019D02 my input",
			Input:   day02myInput,
			Result1: "3562672",
			Result2: "8250"}, // my puzzle input
	}
	for _, tt := range testCases {
		tt.Test(Day02, assert)
	}
}

func BenchmarkDay02(b *testing.B) {
	aoc.Benchmark(Day02, b, day02myInput)
}

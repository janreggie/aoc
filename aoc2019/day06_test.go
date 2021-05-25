package aoc2019

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay06(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Input: day06sampleInput,
			Result1: "54",
			Result2: "4"},
		{Details: "Y2019D06 my input",
			Input:   day06myInput,
			Result1: "154386",
			Result2: "346"},
	}
	for _, tt := range testCases {
		tt.Test(Day06, assert)
	}
}

func BenchmarkDay06(b *testing.B) {
	aoc.Benchmark(Day06, b, day06myInput)
}

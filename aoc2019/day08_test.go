package aoc2019

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay08(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Details: "Y2019D08 my input (no Part 2 assert)",
			Input:   day08myInput,
			Result1: "1560"},
	}
	for _, tt := range testCases {
		tt.Test(Day08, assert)
	}
}

func BenchmarkDay08(b *testing.B) {
	internal.Benchmark(Day08, b, day08myInput)
}

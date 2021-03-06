package aoc2016

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay05(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Input: "abc",
			Result1: "18f47a30",
			Result2: "05ace8e3"},
		{Details: "Y2016D05 my input",
			Input:   day05myInput,
			Result1: "c6697b55",
			Result2: "8c35d1ab"},
	}

	for _, tt := range testCases {
		tt.Test(Day05, assert)
	}
}

func BenchmarkDay05(b *testing.B) {
	aoc.Benchmark(Day05, b, day05myInput)
}

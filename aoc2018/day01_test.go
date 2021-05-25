package aoc2018

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Input: "+1\n-2\n+3\n+1",
			Result1: "3",
			Result2: "2"},
		{Details: "Y2018D01 my input",
			Input:   day01myInput,
			Result1: "420",
			Result2: "227"},
	}
	for _, tt := range testCases {
		tt.Test(Day01, assert)
	}
}

func BenchmarkDay01(b *testing.B) {
	aoc.Benchmark(Day01, b, day01myInput)
}

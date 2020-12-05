package aoc2018

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Input: "+1 -2 +3 +1",
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
	internal.Benchmark(Day01, b, day01myInput)
}

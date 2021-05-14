package aoc2016

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Input: "R8, R4, R4, R8",
			Result1: "8",
			Result2: "4"},
		{Details: "Y2016D01 my input",
			Input:   day01myInput,
			Result1: "291",
			Result2: "159"},
	}
	for _, tt := range testCases {
		tt.Test(Day01, assert)
	}
}

func BenchmarkDay01(b *testing.B) {
	internal.Benchmark(Day01, b, day01myInput)
}

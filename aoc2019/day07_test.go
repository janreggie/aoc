package aoc2019

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay07(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Details: "Y2019D07 my input",
			Input:   day07myInput,
			Result1: "262086",
			Result2: ""},
	}
	for _, tt := range testCases {
		tt.Test(Day07, assert)
	}
}

func BenchmarkDay07(b *testing.B) {
	aoc.Benchmark(Day07, b, day07myInput)
}

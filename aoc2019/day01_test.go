package aoc2019

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Input: "14",
			Result1: "2",
			Result2: "2"},
		{Input: "1969",
			Result1: "654",
			Result2: "966"},
		{Input: "100756",
			Result1: "33583",
			Result2: "50346"},
		{Details: "Y2019D01 my input",
			Input:   day01myInput,
			Result1: "3330521",
			Result2: "4992931"},
	}
	for _, tt := range testCases {
		tt.Test(Day01, assert)
	}
}

func BenchmarkDay01(b *testing.B) {
	internal.Benchmark(Day01, b, day01myInput)
}

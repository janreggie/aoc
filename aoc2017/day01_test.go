package aoc2017

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Input: "1122",
			Result1: "3",
			Result2: "0"},
		{Input: "1212",
			Result1: "0",
			Result2: "6"},
		{Input: "1234",
			Result1: "0",
			Result2: "0"},
		{Input: "91212129",
			Result1: "9",
			Result2: "6"},
		{Details: "Y2017D01 my input",
			Input:   day01myInput,
			Result1: "1182",
			Result2: "1152",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day01, assert)
	}
}

func BenchmarkDay01(b *testing.B) {
	internal.Benchmark(Day01, b, day01myInput)
}

package aoc2019

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay10(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Input: day10sampleInput,
			Result1: "210",
			Result2: "802"},
		{Details: "Y2019D10 my input",
			Input:   day10myInput,
			Result1: "274",
			Result2: "305"},
	}
	for _, tt := range testCases {
		tt.Test(Day10, assert)
	}
}

func Benchmark_generateSisonPoints(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		channel := generateSisonPoints(10, 15)
		for range channel {
			// do nothing
		}
	}
}
func Benchmark_generateSisonPointsSlow(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		channel := generateSisonPointsSlow(10, 15)
		for range channel {
			// do nothing
		}
	}
}

func BenchmarkDay10(b *testing.B) {
	aoc.Benchmark(Day10, b, day10myInput)
}

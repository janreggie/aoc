package aoc2015

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay20(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Details: "Y2015D20 my input",
			Input:   day20myInput,
			Result1: "665280",
			Result2: "705600",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day20, assert)
	}
}

func BenchmarkDay20(b *testing.B) {
	aoc.Benchmark(Day20, b, day20myInput)
}

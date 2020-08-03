package aoc2015

import (
	"testing"

	"github.com/janreggie/AdventOfCode/internal"
	"github.com/stretchr/testify/assert"
)

const day20myInput = `29000000`

func TestDay20(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
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
	internal.Benchmark(Day20, b, day20myInput)
}

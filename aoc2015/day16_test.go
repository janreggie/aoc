package aoc2015

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay16(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Details: "Y2015D16 my test case",
			Input:   day16myInput,
			Result1: "40",
			Result2: "241"},
	}
	for _, tt := range testCases {
		tt.Test(Day16, assert)
	}
}

func BenchmarkDay16(b *testing.B) {
	aoc.Benchmark(Day16, b, day16myInput)
}

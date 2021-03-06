package aoc2015

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Input: ")",
			Result1: "-1",
			Result2: "1"},
		{Input: "()())",
			Result1: "-1",
			Result2: "5"},
		{Details: "Y2015D01 my input",
			Input:   day01myInput,
			Result1: "138",
			Result2: "1771"},
		{Details: "Y2015D01 erroneous input",
			Input:   "(hi)",
			WantErr: true},
	}
	for _, tt := range testCases {
		tt.Test(Day01, assert)
	}
}

func BenchmarkDay01(b *testing.B) {
	aoc.Benchmark(Day01, b, day01myInput)
}

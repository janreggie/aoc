package aoc2015

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay03(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Input: "^v",
			Result1: "2",
			Result2: "3"},
		{Input: "^>v<",
			Result1: "4",
			Result2: "3"},
		{Input: "^v^v^v^v^v",
			Result1: "2",
			Result2: "11"},
		{Details: "Y2015D03 my input",
			Input:   day03myInput,
			Result1: "2081",
			Result2: "2341"},
	}
	for _, tt := range testCases {
		tt.Test(Day03, assert)
	}
}
func BenchmarkDay03(b *testing.B) {
	internal.Benchmark(Day03, b, day01myInput)
}

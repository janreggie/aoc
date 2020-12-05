package aoc2016

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay03(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Details: "Y2016D03 Part 1",
			Input:   "5 10 25",
			Result1: "0"},
		{Details: "Y2016D03 Part 2",
			Input: "101 301 501\n" +
				"102 302 502\n" +
				"103 303 503\n" +
				"201 401 601\n" +
				"202 402 602\n" +
				"203 403 603\n",
			Result2: "6"},
		{Details: "Y2016D03 my input",
			Input:   day03myInput,
			Result1: "869",
			Result2: "1544"},
	}

	for _, tt := range testCases {
		tt.Test(Day03, assert)
	}
}

func BenchmarkDay03(b *testing.B) {
	internal.Benchmark(Day03, b, day03myInput)
}

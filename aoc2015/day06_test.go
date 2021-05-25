package aoc2015

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay06(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Input: "toggle 936,774 through 937,775\n" + // ans1 += 4; ans2 += 8
			"turn off 116,843 through 533,934\n" + // nothing
			"turn on 950,906 through 986,993\n", // ans1 += 3256; ans2 += 3256
			Result1: "3260",
			Result2: "3264"},
		{Input: "toggle 0,0 through 999,999\n" + // ans1 += 1000000; ans2 += 2000000
			"turn on 936,388 through 948,560\n" + // ans1 += 0; ans2 += 2249
			"turn off 485,17 through 655,610\n", // ans1 -= 101574; ans2 -= 101574
			Result1: "898426",
			Result2: "1900675"},
		{Details: "Y2015D06 my input",
			Input:   day06myInput,
			Result1: "543903",
			Result2: "14687245"},
	}
	for _, tt := range testCases {
		tt.Test(Day06, assert)
	}
}
func TestDay06Deux(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Input: "toggle 936,774 through 937,775\n" + // ans1 += 4; ans2 += 8
			"turn off 116,843 through 533,934\n" + // nothing
			"turn on 950,906 through 986,993\n", // ans1 += 3256; ans2 += 3256
			Result1: "3260",
			Result2: "3264"},
		{Input: "toggle 0,0 through 999,999\n" + // ans1 += 1000000; ans2 += 2000000
			"turn on 936,388 through 948,560\n" + // ans1 += 0; ans2 += 2249
			"turn off 485,17 through 655,610\n", // ans1 -= 101574; ans2 -= 101574
			Result1: "898426",
			Result2: "1900675"},
		{Details: "Y2015D06 my input",
			Input:   day06myInput,
			Result1: "543903",
			Result2: "14687245"},
	}
	for _, tt := range testCases {
		tt.Test(Day06, assert)
	}
}

func BenchmarkDay06(b *testing.B) {
	aoc.Benchmark(Day06, b, day06myInput)
}

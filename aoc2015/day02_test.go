package aoc2015

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay02(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Input: "2x3x4",
			Result1: "58",
			Result2: "34"},
		{Input: "1x1x10",
			Result1: "43",
			Result2: "14"},
		{Input: "2x3x4\n1x1x10",
			Result1: "101",
			Result2: "48"},
		{Details: "Y2015D02 my input",
			Input:   day02myInput,
			Result1: "1606483",
			Result2: "3842356"},
		{Details: "Y2015D02 erroneous Input: AxB only",
			Input:   "4x2",
			WantErr: true},
		{Details: "Y2015D02 erroneous Input: AxBxCxD",
			Input:   "4x2x4x5",
			WantErr: true},
	}
	for _, tt := range testCases {
		t.Logf("testing case %#v", tt)
		tt.Test(Day02, assert)
	}
}

func BenchmarkDay02(b *testing.B) {
	internal.Benchmark(Day02, b, day01myInput)
}

package aoc2016

import (
	"bufio"
	"strings"
	"testing"

	"github.com/janreggie/AdventOfCode/tools"
	"github.com/stretchr/testify/assert"
)

const day05myInput = `ffykfhsq`

func TestDay05(t *testing.T) {
	assert := assert.New(t)
	testCases := []tools.TestCase{
		{Input: "abc",
			Result1: "18f47a30",
			Result2: "05ace8e3"},
		{Details: "Y2016D05 my input",
			Input:   day05myInput,
			Result1: "c6697b55",
			Result2: "8c35d1ab"},
	}

	for _, tt := range testCases {
		tt.Test(Day05, assert)
	}
}

func BenchmarkDay05(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		Day05(bufio.NewScanner(strings.NewReader(day05myInput)))
	}
}

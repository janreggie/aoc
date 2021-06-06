package aoc2017

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_spreadsheet(t *testing.T) {

	assert := assert.New(t)

	testCases := []struct {
		input    string
		expected [][]int64
		checksum int64
	}{
		{
			day02sampleInput,
			[][]int64{{5, 1, 9, 5},
				{7, 5, 3},
				{2, 4, 6, 8}},
			18},
	}

	for _, tc := range testCases {
		sheet, err := newSpreadsheet(tc.input)
		assert.NoError(err)
		assert.Equal(tc.expected, sheet.nums)
		assert.Equal(tc.checksum, sheet.checksum())
	}
}

func TestDay02(t *testing.T) {

	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{
			Details: "Y2017D02 sample input",
			Input:   day02sampleInput,
			Result1: "18",
		},
		{
			Details: "Y2017D02 my input",
			Input:   day02myInput,
			Result1: "41887",
			Result2: "226",
		},
	}

	for _, tc := range testCases {
		tc.Test(Day02, assert)
	}
}

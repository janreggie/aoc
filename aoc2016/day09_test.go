package aoc2016

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_decompress(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		input string
		want  string
	}{
		{"ADVENT", "ADVENT"},
		{"A(1x5)BC", "ABBBBBC"},
		{"(3x3)XYZ", "XYZXYZXYZ"},
		{"A(2x2)BCD(2x2)EFG", "ABCBCDEFEFG"},
		{"(6x1)(1x3)A", "(1x3)A"},
		{"X(8x2)(3x3)ABCY", "X(3x3)ABC(3x3)ABCY"},
	}

	for _, tt := range testCases {
		result, err := decompress(tt.input)
		assert.NoError(err, tt)
		assert.Equal(tt.want, result, tt)
	}
}

func TestDay09(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Input: "ADVENT",
			Result1: "6"},
		{Input: "A(1x5)BC",
			Result1: "7"},
		{Input: "(3x3)XYZ",
			Result1: "9"},
		{Input: "A(2x2)BCD(2x2)EFG",
			Result1: "11"},
		{Input: "(6x1)(1x3)A",
			Result1: "6"},
		{Input: "X(8x2)(3x3)ABCY",
			Result1: "18"},
	}
	for _, tt := range testCases {
		tt.Test(Day09, assert)
	}
}

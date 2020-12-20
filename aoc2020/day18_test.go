package aoc2020

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_equation(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		eqn      string
		simple   int
		advanced int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71, 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51, 51},
		{"2 * 3 + (4 * 5)", 26, 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437, 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240, 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632, 23340},
	}

	for _, tt := range testCases {
		simpleAnswer, err := simpleMath(parseEquation(tt.eqn))
		assert.NoError(err)
		assert.Equal(tt.simple, simpleAnswer)
		advancedAnswer, err := advancedMath(parseEquation(tt.eqn))
		assert.NoError(err)
		assert.Equal(tt.advanced, advancedAnswer)
	}
}

func TestDay18(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{
			Details: "Y2020D18 sample input",
			Input:   day18sampleInput,
			Result1: "26457",
			Result2: "694173",
		},
		{
			Details: "Y2020D18 my input",
			Input:   day18myInput,
			Result1: "1890866893020",
			Result2: "34646237037193",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day18, assert)
	}
}

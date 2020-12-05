package aoc2020

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay02(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{
			Details: "Y2020D02 sample input",
			Input:   day02sampleInput,
			Result1: "2",
			Result2: "1",
		},
		{
			Details: "Y2020D02 my input",
			Input:   day02myInput,
			Result1: "396",
			Result2: "428",
		},
	}

	for _, tt := range testCases {
		tt.Test(Day02, assert)
	}
}

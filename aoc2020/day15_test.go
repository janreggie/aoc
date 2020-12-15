package aoc2020

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay15(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{
			Details: "Y2020D15 sample input",
			Input:   day15sampleInput,
			Result1: "436",
			Result2: "175594",
		},
		{
			Details: "Y2020D15 my input",
			Input:   day15myInput,
			Result1: "260",
			Result2: "950",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day15, assert)
	}
}

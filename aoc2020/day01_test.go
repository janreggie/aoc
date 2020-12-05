package aoc2020

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{
			Details: "Y2020D01 sample input",
			Input:   day01sampleInput,
			Result1: "514579",
			Result2: "241861950",
		},
		{
			Details: "Y2020D01 my input",
			Input:   day01myInput,
			Result1: "806656",
			Result2: "230608320",
		},
	}

	for _, tt := range testCases {
		tt.Test(Day01, assert)
	}
}

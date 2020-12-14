package aoc2020

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay13(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{
			Details: "Y2020D13 sample input",
			Input:   day13sampleInput,
			Result1: "295",
			Result2: "1068781",
		},
		{
			Details: "Y2020D13 my input",
			Input:   day13myInput,
			Result1: "370",
			Result2: "894954360381385",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day13, assert)
	}
}

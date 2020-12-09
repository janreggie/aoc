package aoc2020

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay09(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{
			Details: "Y2020D09 my input",
			Input:   day09myInput,
			Result1: "257342611",
			Result2: "35602097",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day09, assert)
	}
}

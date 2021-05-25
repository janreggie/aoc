package aoc2020

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_readBoardingPass(t *testing.T) {
	assert := assert.New(t)

	// use example board passes
	testCases := []struct {
		input  string
		seat   planeSeat
		seatID int
	}{
		{"FBFBBFFRLR", planeSeat{row: 44, col: 5}, 357},
		{"BFFFBBFRRR", planeSeat{row: 70, col: 7}, 567},
		{"FFFBBBFRRR", planeSeat{row: 14, col: 7}, 119},
		{"BBFFBBFRLL", planeSeat{row: 102, col: 4}, 820},
	}

	for _, tt := range testCases {
		actual, err := readBoardingPass(tt.input)
		assert.NoError(err)
		assert.Equal(tt.seat, actual)
		assert.Equal(tt.seatID, actual.seatID())
	}
}

func TestDay05(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{
			Details: "Y2020D05 my input",
			Input:   day05myInput,
			Result1: "930",
			Result2: "515",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day05, assert)
	}
}

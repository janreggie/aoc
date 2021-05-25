package aoc2020

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_waitingArea(t *testing.T) {
	assert := assert.New(t)
	// variables for later
	const e, o, f = empty, occupied, floor
	area, err := generateWaitingArea(day11sampleInput)
	assert.NoError(err)
	assert.Equal(waitingArea{
		width:  10,
		height: 10,
		stable: false,
		representation: [][]seat{
			{e, f, e, e, f, e, e, f, e, e},
			{e, e, e, e, e, e, e, f, e, e},
			{e, f, e, f, e, f, f, e, f, f},
			{e, e, e, e, f, e, e, f, e, e},
			{e, f, e, e, f, e, e, f, e, e},
			{e, f, e, e, e, e, e, f, e, e},
			{f, f, e, f, e, f, f, f, f, f},
			{e, e, e, e, e, e, e, e, e, e},
			{e, f, e, e, e, e, e, e, f, e},
			{e, f, e, e, e, e, e, f, e, e},
		},
	}, area)

	area = area.iterateSimple()
	assert.Equal(waitingArea{
		width:  10,
		height: 10,
		stable: false,
		representation: [][]seat{
			{o, f, o, o, f, o, o, f, o, o},
			{o, o, o, o, o, o, o, f, o, o},
			{o, f, o, f, o, f, f, o, f, f},
			{o, o, o, o, f, o, o, f, o, o},
			{o, f, o, o, f, o, o, f, o, o},
			{o, f, o, o, o, o, o, f, o, o},
			{f, f, o, f, o, f, f, f, f, f},
			{o, o, o, o, o, o, o, o, o, o},
			{o, f, o, o, o, o, o, o, f, o},
			{o, f, o, o, o, o, o, f, o, o},
		},
	}, area)

	area = area.iterateSimple()
	assert.Equal(waitingArea{
		width:  10,
		height: 10,
		stable: false,
		representation: [][]seat{
			{o, f, e, e, f, e, o, f, o, o},
			{o, e, e, e, e, e, e, f, e, o},
			{e, f, e, f, e, f, f, e, f, f},
			{o, e, e, e, f, e, e, f, e, o},
			{o, f, e, e, f, e, e, f, e, e},
			{o, f, e, e, e, e, o, f, o, o},
			{f, f, e, f, e, f, f, f, f, f},
			{o, e, e, e, e, e, e, e, e, o},
			{o, f, e, e, e, e, e, e, f, e},
			{o, f, o, e, e, e, e, f, o, o},
		},
	}, area)

	for !area.isStable() {
		area = area.iterateSimple()
	}
	assert.Equal(waitingArea{
		width:  10,
		height: 10,
		stable: true,
		representation: [][]seat{
			{o, f, o, e, f, e, o, f, o, o},
			{o, e, e, e, o, e, e, f, e, o},
			{e, f, o, f, e, f, f, o, f, f},
			{o, e, o, o, f, o, o, f, e, o},
			{o, f, o, e, f, e, e, f, e, e},
			{o, f, o, e, o, e, o, f, o, o},
			{f, f, e, f, e, f, f, f, f, f},
			{o, e, o, e, o, o, e, o, e, o},
			{o, f, e, e, e, e, e, e, f, e},
			{o, f, o, e, o, e, o, f, o, o},
		},
	}, area)
}

func TestDay11(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{
			Details: "Y2020D11 sample input",
			Input:   day11sampleInput,
			Result1: "37",
			Result2: "26",
		},
		{
			Details: "Y2020D11 my input",
			Input:   day11myInput,
			Result1: "2453",
			Result2: "",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day11, assert)
	}
}

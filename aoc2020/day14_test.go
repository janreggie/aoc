package aoc2020

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_mask(t *testing.T) {
	assert := assert.New(t)
	mask, err := newMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	assert.NoError(err)
	assert.Equal("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", mask.String())
	insOuts := [][2]uint64{
		{11, 73},
		{101, 101},
		{0, 64},
	}
	for _, io := range insOuts {
		assert.Equal(io[1], mask.apply(io[0]))
	}

	mask, err = newMask("000000000000000000000000000000X1001X")
	assert.NoError(err)
	assert.ElementsMatch([]uint64{26, 27, 58, 59}, mask.v2floaters(42))
}

func TestDay14(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		// This will be horribly slow...
		// {
		// 	Details: "Y2020D14 sample input",
		// 	Input:   day14sampleInput,
		// 	Result1: "165",
		// 	Result2: "",
		// },
		{
			Details: "Y2020D14 my input",
			Input:   day14myInput,
			Result1: "11327140210986",
			Result2: "",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day14, assert)
	}
}

package aoc2016

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

const day08myInput = `rect 1x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 5
rect 1x1
rotate row y=0 by 3
rect 1x1
rotate row y=0 by 3
rect 2x1
rotate row y=0 by 5
rect 1x1
rotate row y=0 by 5
rect 4x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 5
rect 4x1
rotate row y=0 by 3
rect 2x1
rotate row y=0 by 5
rect 4x1
rotate row y=0 by 2
rect 1x2
rotate row y=1 by 6
rotate row y=0 by 2
rect 1x2
rotate column x=32 by 1
rotate column x=23 by 1
rotate column x=13 by 1
rotate row y=0 by 6
rotate column x=0 by 1
rect 5x1
rotate row y=0 by 2
rotate column x=30 by 1
rotate row y=1 by 20
rotate row y=0 by 18
rotate column x=13 by 1
rotate column x=10 by 1
rotate column x=7 by 1
rotate column x=2 by 1
rotate column x=0 by 1
rect 17x1
rotate column x=16 by 3
rotate row y=3 by 7
rotate row y=0 by 5
rotate column x=2 by 1
rotate column x=0 by 1
rect 4x1
rotate column x=28 by 1
rotate row y=1 by 24
rotate row y=0 by 21
rotate column x=19 by 1
rotate column x=17 by 1
rotate column x=16 by 1
rotate column x=14 by 1
rotate column x=12 by 2
rotate column x=11 by 1
rotate column x=9 by 1
rotate column x=8 by 1
rotate column x=7 by 1
rotate column x=6 by 1
rotate column x=4 by 1
rotate column x=2 by 1
rotate column x=0 by 1
rect 20x1
rotate column x=47 by 1
rotate column x=40 by 2
rotate column x=35 by 2
rotate column x=30 by 2
rotate column x=10 by 3
rotate column x=5 by 3
rotate row y=4 by 20
rotate row y=3 by 10
rotate row y=2 by 20
rotate row y=1 by 16
rotate row y=0 by 9
rotate column x=7 by 2
rotate column x=5 by 2
rotate column x=3 by 2
rotate column x=0 by 2
rect 9x2
rotate column x=22 by 2
rotate row y=3 by 40
rotate row y=1 by 20
rotate row y=0 by 20
rotate column x=18 by 1
rotate column x=17 by 2
rotate column x=16 by 1
rotate column x=15 by 2
rotate column x=13 by 1
rotate column x=12 by 1
rotate column x=11 by 1
rotate column x=10 by 1
rotate column x=8 by 3
rotate column x=7 by 1
rotate column x=6 by 1
rotate column x=5 by 1
rotate column x=3 by 1
rotate column x=2 by 1
rotate column x=1 by 1
rotate column x=0 by 1
rect 19x1
rotate column x=44 by 2
rotate column x=40 by 3
rotate column x=29 by 1
rotate column x=27 by 2
rotate column x=25 by 5
rotate column x=24 by 2
rotate column x=22 by 2
rotate column x=20 by 5
rotate column x=14 by 3
rotate column x=12 by 2
rotate column x=10 by 4
rotate column x=9 by 3
rotate column x=7 by 3
rotate column x=3 by 5
rotate column x=2 by 2
rotate row y=5 by 10
rotate row y=4 by 8
rotate row y=3 by 8
rotate row y=2 by 48
rotate row y=1 by 47
rotate row y=0 by 40
rotate column x=47 by 5
rotate column x=46 by 5
rotate column x=45 by 4
rotate column x=43 by 2
rotate column x=42 by 3
rotate column x=41 by 2
rotate column x=38 by 5
rotate column x=37 by 5
rotate column x=36 by 5
rotate column x=33 by 1
rotate column x=28 by 1
rotate column x=27 by 5
rotate column x=26 by 5
rotate column x=25 by 1
rotate column x=23 by 5
rotate column x=22 by 1
rotate column x=21 by 2
rotate column x=18 by 1
rotate column x=17 by 3
rotate column x=12 by 2
rotate column x=11 by 2
rotate column x=7 by 5
rotate column x=6 by 5
rotate column x=5 by 4
rotate column x=3 by 5
rotate column x=2 by 5
rotate column x=1 by 3
rotate column x=0 by 4
`

func Test_littleScreen_rect(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		width, height uint
		want          littleScreen
		wantErr       bool
	}{
		{
			width:  3,
			height: 2,
			want: littleScreen{
				lcd: [6][50]bool{
					{true, true, true},
					{true, true, true},
				}, // the rest is false!
			},
		},
		{
			width:   51,
			height:  1,
			wantErr: true,
		},
		{
			width:   1,
			height:  7,
			wantErr: true,
		},
	}

	for _, tt := range testCases {
		var ls littleScreen
		err := ls.rect(tt.width, tt.height)
		if tt.wantErr {
			assert.Error(err, err)
			return
		}
		assert.NoError(err, err)
		assert.Equal(tt.want, ls)
	}
}

func Test_littleScreen_rotateRow(t *testing.T) {
	assert := assert.New(t)
	var ls littleScreen
	assert.NoError(ls.rect(3, 2), "created 3x2 rectangle")
	assert.Equal(
		littleScreen{
			lcd: [6][50]bool{
				{true, true, true},
				{true, true, true},
			},
		},
		ls,
		"created 3x2 rectangle",
	)

	// rotate row y=0 by 4
	assert.NoError(ls.rotateRow(0, 4), "rotate y=0 by 4")
	assert.Equal(
		littleScreen{
			lcd: [6][50]bool{
				{false, false, false, false, true, true, true},
				{true, true, true},
			},
		},
		ls,
		"rotate y=0 by 4",
	)

	// rotate row y=6 by 4
	assert.Error(ls.rotateRow(6, 4), "should not be able to rotate y=6")
}

func Test_littleScreen_rotateColumn(t *testing.T) {
	assert := assert.New(t)
	var ls littleScreen
	assert.NoError(ls.rect(3, 2), "created 3x2 rectangle")
	assert.Equal(
		littleScreen{
			lcd: [6][50]bool{
				{true, true, true},
				{true, true, true},
			},
		},
		ls,
		"created 3x2 rectangle",
	)

	// rotate col x=1 by 1
	assert.NoError(ls.rotateColumn(1, 1), "rotate col x=1 by 1")
	assert.Equal(
		littleScreen{
			lcd: [6][50]bool{
				{true, false, true},
				{true, true, true},
				{false, true, false},
			},
		},
		ls,
		"rotate col x=1 by 1",
	)

	// rotate col x=50 by 1
	assert.Error(ls.rotateColumn(50, 1), "should not be able to rotate x=50")
}

func Test_littleScreen_parseInstruction(t *testing.T) {
	assert := assert.New(t)
	var ls littleScreen

	// rect 3x2
	instr := "rect 3x2"
	assert.NoError(ls.parseInstruction(instr), "should be able to create 3x2")
	assert.Equal(
		littleScreen{
			lcd: [6][50]bool{
				{true, true, true},
				{true, true, true},
			},
		},
		ls,
		instr,
	)

	// rotate column x=1 by 1
	instr = "rotate column x=1 by 1"
	assert.NoError(ls.parseInstruction(instr), "should be able to ", instr)
	assert.Equal(
		littleScreen{
			lcd: [6][50]bool{
				{true, false, true},
				{true, true, true},
				{false, true},
			},
		},
		ls,
		instr,
	)

	// rotate row y=0 by 4
	instr = "rotate row y=0 by 4"
	assert.NoError(ls.parseInstruction(instr), "should be able to ", instr)
	assert.Equal(
		littleScreen{
			lcd: [6][50]bool{
				{false, false, false, false, true, false, true},
				{true, true, true},
				{false, true, false},
			},
		},
		ls,
		instr,
	)
}

func TestDay08(t *testing.T) {
	assert := assert.New(t)
	testCase := internal.TestCase{
		Details: "Y2016D08 my input",
		Input:   day08myInput,
		Result1: `110`,
		// Result2 difficult to formalize
	}
	testCase.Test(Day08, assert)
}

func BenchmarkDay08(b *testing.B) {
	internal.Benchmark(Day08, b, day08myInput)
}

package aoc2016

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

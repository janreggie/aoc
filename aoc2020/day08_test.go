package aoc2020

import (
	"bufio"
	"strings"
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_instruction(t *testing.T) {
	assert := assert.New(t)
	instrs, err := generateInstructionList(bufio.NewScanner(strings.NewReader(day08sampleInput)))
	assert.NoError(err)
	assert.Equal([]instruction{
		{"nop", +0},
		{"acc", +1},
		{"jmp", +4},
		{"acc", +3},
		{"jmp", -3},
		{"acc", -99},
		{"acc", +1},
		{"jmp", -4},
		{"acc", +6},
	}, instrs)
}

func TestDay08(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{
			Details: "Y2020D08 sample input",
			Input:   day08sampleInput,
			Result1: "5",
			Result2: "8",
		},
		{
			Details: "Y2020D08 my input",
			Input:   day08myInput,
			Result1: "2058",
			Result2: "1000",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day08, assert)
	}
}

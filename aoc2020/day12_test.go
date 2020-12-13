package aoc2020

import (
	"bufio"
	"strings"
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_navigationInstruction(t *testing.T) {
	assert := assert.New(t)
	instrs, err := generateNavigationInstructions(bufio.NewScanner(strings.NewReader(day12sampleInput)))
	assert.NoError(err)
	assert.Equal([]navigationInstruction{
		{'F', 10},
		{'N', 3},
		{'F', 7},
		{'R', 90},
		{'F', 11},
	}, instrs)
}

func Test_ship(t *testing.T) {
	assert := assert.New(t)
	// guaranteed no error (from test above)
	instrs, _ := generateNavigationInstructions(bufio.NewScanner(strings.NewReader(day12sampleInput)))
	ss := ship{}
	assert.Equal(ship{x: 0, y: 0, direction: east}, ss)
	// states represent the state that should exist after launching every value at instrs.
	states := []ship{
		{x: 10, y: 0, direction: east},   // F10
		{x: 10, y: 3, direction: east},   // N3
		{x: 17, y: 3, direction: east},   // F7
		{x: 17, y: 3, direction: south},  // R90
		{x: 17, y: -8, direction: south}, // F11
	}
	for ii, instr := range instrs {
		ss.readInstruction(instr)
		assert.Equal(states[ii], ss)
	}

	ss = ship{}
	instrStates := []struct {
		instr navigationInstruction
		state ship
	}{
		{navigationInstruction{'N', 10}, ship{x: 0, y: 10, direction: east}},
		{navigationInstruction{'W', 10}, ship{x: -10, y: 10, direction: east}},
		{navigationInstruction{'S', 15}, ship{x: -10, y: -5, direction: east}},
		{navigationInstruction{'E', 18}, ship{x: 8, y: -5, direction: east}},
		{navigationInstruction{'F', 3}, ship{x: 11, y: -5, direction: east}},
		{navigationInstruction{'R', 90}, ship{x: 11, y: -5, direction: south}},
		{navigationInstruction{'F', 4}, ship{x: 11, y: -9, direction: south}},
		{navigationInstruction{'R', 90}, ship{x: 11, y: -9, direction: west}},
		{navigationInstruction{'F', 3}, ship{x: 8, y: -9, direction: west}},
		{navigationInstruction{'W', 3}, ship{x: 5, y: -9, direction: west}},
	}
	for _, iiss := range instrStates {
		ss.readInstruction(iiss.instr)
		assert.Equal(iiss.state, ss)
	}

	ss = ship{wx: 10, wy: 1}
	states = []ship{
		{100, 10, east, 10, 1},   // F10
		{100, 10, east, 10, 4},   // N3
		{170, 38, east, 10, 4},   // F7
		{170, 38, east, 4, -10},  // R90
		{214, -72, east, 4, -10}, // F11
	}
	for ii, instr := range instrs {
		ss.readWaypointy(instr)
		assert.Equal(states[ii], ss)
	}
}

func TestDay12(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{
			Details: "Y2020D12 sample input",
			Input:   day12sampleInput,
			Result1: "25",
			Result2: "286",
		},
		{
			Details: "Y2020D12 my input",
			Input:   day12myInput,
			Result1: "441",
			Result2: "40014",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day12, assert)
	}
}

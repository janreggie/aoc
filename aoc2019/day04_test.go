package aoc2019

import (
	"testing"

	"github.com/janreggie/AdventOfCode/structs"
	"github.com/stretchr/testify/assert"
)

func Test_checkAdjacent(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"12", 12, false},
		{"11", 11, true},
		{"123455", 123455, true},
		{"124464", 124464, true},
		{"123456", 123456, false},
	}
	for _, tt := range tests {
		assert.Equal(checkAdjacent(tt.input), tt.want, tt.name)
	}
}

func Test_checkNonDecreasing(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"12", 12, true},
		{"11", 11, true},
		{"10", 10, false},
		{"123455", 123455, true},
		{"124464", 124464, false},
		{"123456", 123456, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		assert.Equal(checkNonDecreasing(tt.input), tt.want, tt.name)
	}
}

func Test_checkDuality(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"112233", 112233, true},
		{"123444", 123444, false},
		{"111122", 111122, true},
		{"123456", 123456, false},
		{"123345", 123345, true},
		{"122333", 122333, true},
		{"114444", 114444, true},
	}
	for _, tt := range tests {
		assert.Equal(checkDuality(tt.input), tt.want, tt.name)
	}
}

func TestDay04(t *testing.T) {
	assert := assert.New(t)
	testCases := []structs.TestCase{
		{Details: "Y2019D04 my input",
			Input:   "271973-785961",
			Result1: "925",
			Result2: "607"},
	}
	for _, tt := range testCases {
		tt.Test(Day04, assert)
	}
}

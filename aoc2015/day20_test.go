package aoc2015

import (
	"testing"

	"github.com/janreggie/AdventOfCode/tools"
	"github.com/stretchr/testify/assert"
)

func Test_factor(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		input uint
		want  factors
	}{
		{2, factors{2: 1}},
		{7, factors{7: 1}},
		{144, factors{2: 4, 3: 2}},
	}

	for _, tt := range testCases {
		assert.Equal(tt.want, factor(tt.input), tt.input)
	}
}

func Test_factors_sumOfDivisors(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		input uint
		want  uint
	}{
		{2, 3},
		{7, 8},
		{144, 403},
	}

	for _, tt := range testCases {
		factors := factor(tt.input)
		assert.Equal(tt.want, factors.sumOfDivisors(), tt.input)
	}
}

func Test_factors_countDivisors(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		input uint
		want  uint
	}{
		{2, 2},
		{7, 2},
		{144, 15},
	}

	for _, tt := range testCases {
		factors := factor(tt.input)
		assert.Equal(tt.want, factors.countDivisors(), tt.input)
	}
}

func TestDay20(t *testing.T) {
	assert := assert.New(t)
	testCases := []tools.TestCase{
		{Details: "Y2015D20 my input",
			Input:   "29000000",
			Result1: "665280",
			// Result2 unimplemented
		},
	}
	for _, tt := range testCases {
		tt.Test(Day20, assert)
	}
}

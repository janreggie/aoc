package aoc2018_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2018"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{"+1 -2 +3 +1", "3", "2"},
	}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2018.Day01, assert)
	}
}

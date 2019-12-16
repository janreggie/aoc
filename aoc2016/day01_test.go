package aoc2016_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2016"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{"R8, R4, R4, R8\n", "8", "4"},
	}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2016.Day01, assert)
	}
}

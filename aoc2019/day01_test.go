package aoc2019_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2019"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{"14", "2", "2"},
		{"1969", "654", "966"},
		{"100756", "33583", "50346"},
	}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2019.Day01, assert)
	}
}

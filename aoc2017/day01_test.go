package aoc2017_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2017"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{"1122\n", "3", "0"},
		{"1212\n", "0", "6"},
		{"1234\n", "0", "0"},
		{"91212129\n", "9", "6"},
	}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2017.Day01, assert)
	}
}

package aoc2015_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2015"
	"github.com/stretchr/testify/assert"
)

func TestDay02(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{"2x3x4", "58", "34"},
		{"1x1x10", "43", "14"},
		{"2x3x4\n1x1x10", "101", "48"}}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2015.Day02, assert)
	}
}

package aoc2015_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2015"
	"github.com/stretchr/testify/assert"
)

func TestDay08(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{`""`, "2", "4"},
		{`"abc"`, "2", "4"},
		{`"aaa\"aaa"`, "3", "6"},
		{`"\x27"`, "5", "5"},
	}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2015.Day08, assert)
	}
}

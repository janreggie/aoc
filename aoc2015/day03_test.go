package aoc2015_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2015"
	"github.com/stretchr/testify/assert"
)

func TestDay03(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{"^v", "2", "3"},
		{"^>v<", "4", "3"},
		{"^v^v^v^v^v", "2", "11"},
	}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2015.Day03, assert)
	}
}

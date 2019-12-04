package aoc2019_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2019"
	"github.com/stretchr/testify/assert"
)

func TestDay04(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{"271973-785961", "925", "607"},
	}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2019.Day04, assert)
	}
}

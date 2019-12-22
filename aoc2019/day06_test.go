package aoc2019_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2019"
	"github.com/stretchr/testify/assert"
)

func TestDay06(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{`COM)B
		B)C
		C)D
		D)E
		E)F
		B)G
		G)H
		D)I
		E)J
		J)K
		K)L
		K)YOU
		I)SAN
		`, "54", "4"},
	}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2019.Day06, assert)
	}
}

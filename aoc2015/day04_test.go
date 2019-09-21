package aoc2015_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2015"
	"github.com/stretchr/testify/assert"
)

func TestDay04(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{"abcdef", "609043", "6742839"},
		{"pqrstuv", "1048970", "5714438"},
	}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2015.Day04, assert)
	}
}

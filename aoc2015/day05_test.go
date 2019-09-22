package aoc2015_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2015"
	"github.com/stretchr/testify/assert"
)

func TestDay05(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{"ugknbfddgicrmopn", "1", "0"},
		{"aaa", "1", "0"},
		{"jchzalrnumimnmhp", "0", "0"},
		{"haegwjzuvuyypxyu", "0", "0"},
		{"qjhvhtzxzqqjkmpb", "0", "1"},
		{"xxyxx", "0", "1"},
		{"uurcxstgmygtbstg", "0", "0"},
		{"ieodomkazucvgmuy", "0", "0"},
	}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2015.Day05, assert)
	}
}

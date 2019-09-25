package aoc2015_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2015"
	"github.com/stretchr/testify/assert"
)

func TestDay06(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{"toggle 936,774 through 937,775\n" + // ans1 += 4; ans2 += 8
			"turn off 116,843 through 533,934\n" + // nothing
			"turn on 950,906 through 986,993\n", // ans1 += 3256; ans2 += 3256
			"3260", "3264"},
		{"toggle 0,0 through 999,999\n" + // ans1 += 1000000; ans2 += 2000000
			"turn on 936,388 through 948,560\n" + // ans1 += 0; ans2 += 2249
			"turn off 485,17 through 655,610\n", // ans1 -= 101574; ans2 -= 101574
			"898426", "1900675"},
	}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2015.Day06, assert)
	}
}

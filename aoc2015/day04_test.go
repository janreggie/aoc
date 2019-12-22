package aoc2015_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2015"
	"github.com/stretchr/testify/assert"
)

func TestDay04(t *testing.T) {
	assert := assert.New(t)
	testCases := []testCase{
		{"abcdef", "609043", "6742839"},
		{"pqrstuv", "1048970", "5714438"},
		{"iwrupvqb", "346386", "9958218"},
	}
	for _, eachCase := range testCases {
		testTestCase(eachCase, aoc2015.Day04, assert)
	}
}

func BenchmarkDay04(b *testing.B) {
	var ans1, ans2 string
	var err error
	for n := 0; n < b.N; n++ {
		ans1, ans2, err = aoc2015.Day04(bufio.NewScanner(strings.NewReader("iwrupvqb")))
		ans1, ans2, err = aoc2015.Day04(bufio.NewScanner(strings.NewReader("abcdef")))
		ans1, ans2, err = aoc2015.Day04(bufio.NewScanner(strings.NewReader("pqrstuv")))
	}
	_, _, _ = ans1, ans2, err
}

func BenchmarkDay04ST(b *testing.B) {
	var ans1, ans2 string
	var err error
	for n := 0; n < b.N; n++ {
		ans1, ans2, err = aoc2015.Day04ST(bufio.NewScanner(strings.NewReader("iwrupvqb")))
		ans1, ans2, err = aoc2015.Day04ST(bufio.NewScanner(strings.NewReader("abcdef")))
		ans1, ans2, err = aoc2015.Day04ST(bufio.NewScanner(strings.NewReader("pqrstuv")))
	}
	_, _, _ = ans1, ans2, err
}

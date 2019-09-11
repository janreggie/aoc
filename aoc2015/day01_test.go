package aoc2015_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2015"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input   string
	result1 string
	result2 string
}

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCase := []testCase{
		testCase{")", "-1", "1"},
		testCase{"()())", "-1", "5"},
	}
	for _, eachCase := range testCase {
		scanner := bufio.NewScanner(strings.NewReader(eachCase.input))
		answer1, answer2, err := aoc2015.Day01(scanner)
		assert.NoError(err)
		assert.Equal(eachCase.result1, answer1)
		assert.Equal(eachCase.result2, answer2)
	}
}

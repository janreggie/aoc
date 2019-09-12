package aoc2015_test

import (
	"bufio"
	"strings"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input   string
	result1 string
	result2 string
}

func testTestCase(expected testCase, f func(*bufio.Scanner) (string, string, error), a *assert.Assertions) {
	scanner := bufio.NewScanner(strings.NewReader(expected.input))
	answer1, answer2, err := f(scanner)
	a.NoError(err)
	a.Equal(expected.result1, answer1)
	a.Equal(expected.result2, answer2)
}

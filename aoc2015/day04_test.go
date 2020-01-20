package aoc2015

import (
	"bufio"
	"strings"
	"testing"

	"github.com/janreggie/AdventOfCode/structs"
	"github.com/stretchr/testify/assert"
)

func TestDay04(t *testing.T) {
	assert := assert.New(t)
	testCases := []structs.TestCase{
		// {Input: "abcdef",
		// 	Result1: "609043",
		// 	Result2: "6742839"},
		// {Input: "pqrstuv",
		// 	Result1: "1048970",
		// 	Result2: "5714438"},
		{Details: "Y2015D04 my input",
			Input:   "iwrupvqb",
			Result1: "346386",
			Result2: "9958218"},
	}
	for _, tt := range testCases {
		tt.Test(Day04, assert)
	}
}

func BenchmarkDay04(b *testing.B) {
	var ans1, ans2 string
	var err error
	for n := 0; n < b.N; n++ {
		ans1, ans2, err = Day04(bufio.NewScanner(strings.NewReader("iwrupvqb")))
	}
	_, _, _ = ans1, ans2, err
}

func BenchmarkDay04ST(b *testing.B) {
	var ans1, ans2 string
	var err error
	for n := 0; n < b.N; n++ {
		ans1, ans2, err = Day04ST(bufio.NewScanner(strings.NewReader("iwrupvqb")))
	}
	_, _, _ = ans1, ans2, err
}

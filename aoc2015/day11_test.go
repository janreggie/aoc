package aoc2015

import (
	"bufio"
	"strings"
	"testing"

	"github.com/janreggie/AdventOfCode/structs"
	"github.com/stretchr/testify/assert"
)

func Test_newPassword(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		input string
		want  string
	}{
		{"heqaabcc", "heqaabcc"},
		{"heqa", "aaaaheqa"},
		{"abcdefgH", "abcdefga"},
		{"abcde~Gh", "abcdezah"},
		{"posdfidpoi", "posdfidp"},
	}
	for _, tt := range testCases {
		assert.Equal(tt.want, newPassword(tt.input).string(), tt.input)
	}
}

func Test_increment(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		input string
		want  string
	}{
		{"kayatazy", "kayatazz"},
		{"kayatazz", "kayatbaa"},
		{"zzzzzzzz", "aaaaaaaa"},
	}
	for _, tt := range testCases {
		ps := newPassword(tt.input)
		ps.increment()
		assert.Equal(tt.want, ps.string(), tt.input)
	}
}

func TestDay11(t *testing.T) {
	assert := assert.New(t)
	testCases := []structs.TestCase{
		{Input: "abcdefgh", Result1: "abcdffaa"},
		{Input: "ghijklmn", Result1: "ghjaabcc"},
		{Details: "Y2015D11 my input", Input: "hepxcrrq", Result1: "hepxxyzz", Result2: "heqaabcc"},
	}
	for _, eachCase := range testCases {
		eachCase.Test(Day11, assert)
	}
}

func BenchmarkDay11(b *testing.B) {
	input := "hepxcrrq"
	for ii := 0; ii < b.N; ii++ {
		Day11(bufio.NewScanner(strings.NewReader(input)))
	}
}

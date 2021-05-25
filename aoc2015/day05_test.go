package aoc2015

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func TestDay05(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		// Note that the results represent how many nice strings there are.
		// And since there is only one input,
		// a result of 0 means that the string is naughty (at least in Part One or Two)
		// and a result of 1 means that the string is nice.
		{Input: "ugknbfddgicrmopn",
			Result1: "1",
			Result2: "0"},
		{Input: "aaa",
			Result1: "1",
			Result2: "0"},
		{Input: "jchzalrnumimnmhp",
			Result1: "0",
			Result2: "0"},
		{Input: "haegwjzuvuyypxyu",
			Result1: "0",
			Result2: "0"},
		{Input: "qjhvhtzxzqqjkmpb",
			Result1: "0",
			Result2: "1"},
		{Input: "xxyxx",
			Result1: "0",
			Result2: "1"},
		{Input: "uurcxstgmygtbstg",
			Result1: "0",
			Result2: "0"},
		{Input: "ieodomkazucvgmuy",
			Result1: "0",
			Result2: "0"},
		{Details: "Y2015D05 my input",
			Input:   day05myInput,
			Result1: "238",
			Result2: "69"}, // nice
	}
	for _, tt := range testCases {
		t.Logf("Testing %#v", tt)
		tt.Test(Day05, assert)
	}
}

func Test_checkThreeVowels(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input string
		want  bool
	}{
		{input: "xazegov", want: true},
		{input: "dvszwmarrgswjxmb", want: false},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, checkThreeVowels(tt.input), tt.input)
	}
}

func Test_checkTwiceInARow(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input string
		want  bool
	}{
		{input: "abcdde", want: true},
		{input: "abcdee", want: true},
		{input: "asldkfj", want: false},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, checkTwiceInARow(tt.input), tt.input)
	}
}

func Test_checkSanity(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input string
		want  bool
	}{
		{input: "haegwjzuvuyypxyu", want: false},
		{input: "ugknbfddgicrmopn", want: true},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, checkSanity(tt.input), tt.input)
	}
}

func Test_checkRepeatPair(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input string
		want  bool
	}{
		{input: "aabcdefgaa", want: true},
		{input: "aaa", want: false},
		{input: "ssdflkjghouiwe", want: false},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, checkRepeatPair(tt.input), tt.input)
	}
}

func Test_checkRepeatChar(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input string
		want  bool
	}{
		{input: "aba", want: true},
		{input: "akdirnwlx", want: false},
		{input: "aabbccdd", want: false},
		{input: "aabcbcdd", want: true},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, checkRepeatChar(tt.input), tt.input)
	}
}
func Test_checkNiceOne(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input string
		want  bool
	}{
		{input: "ugknbfddgicrmopn", want: true},
		{input: "aaa", want: true},
		{input: "jchzalrnumimnmhp", want: false},
		{input: "haegwjzuvuyypxyu", want: false},
		{input: "dvszwmarrgswjxmb", want: false},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, checkNiceOne(tt.input), tt.input)
	}
}

func Test_checkNiceTwo(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input string
		want  bool
	}{
		{input: "qjhvhtzxzqqjkmpb", want: true},
		{input: "xxyxx", want: true},
		{input: "uurcxstgmygtbstg", want: false},
		{input: "ieodomkazucvgmuy", want: false},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, checkNiceTwo(tt.input), tt.input)
	}
}

func BenchmarkDay05(b *testing.B) {
	internal.Benchmark(Day05, b, day05myInput)
}

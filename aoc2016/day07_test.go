package aoc2016

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_newIpv7Address(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		input string
		want  ipv7Address
	}{
		{input: "abba[mnop]qrst",
			want: []struct {
				raw       string
				bracketed bool
			}{
				{"abba", false},
				{"mnop", true},
				{"qrst", false},
			}},
	}

	for _, tt := range testCases {
		assert.Equal(tt.want, newIpv7Address(tt.input), tt.input)
	}
}

func Test_ipv7Address_supportTLS(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		input string
		want  bool
	}{
		{input: "abba[mnop]qrst", want: true},
		{input: "abcd[bddb]xyyx", want: false},
		{input: "aaaa[qwer]tyui", want: false},
		{input: "ioxxoj[asdfgh]zxcvbn", want: true},
	}

	for _, tt := range testCases {
		ipv7 := newIpv7Address(tt.input)
		assert.Equal(tt.want, ipv7.supportTLS(), tt.input)
	}
}

func Test_ipv7Address_supportSSL(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		input string
		want  bool
	}{
		{input: "aba[bab]xyz", want: true},
		{input: "xyx[xyx]xyx", want: false},
		{input: "aaa[kek]eke", want: true},
		{input: "zazbz[bzb]cdb", want: true},
	}

	for _, tt := range testCases {
		ipv7 := newIpv7Address(tt.input)
		assert.Equal(tt.want, ipv7.supportSSL(), tt.input)
	}
}

func TestDay07(t *testing.T) {
	assert := assert.New(t)
	testCase := internal.TestCase{
		Details: "Y2016D07 my input",
		Input:   day07myInput,
		Result1: `110`,
		Result2: `242`,
	}
	testCase.Test(Day07, assert)
}

func BenchmarkDay07(b *testing.B) {
	internal.Benchmark(Day07, b, day07myInput)
}

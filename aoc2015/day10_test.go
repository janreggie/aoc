package aoc2015

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

const day10myInput = `3113322113`

func Test_lookAndSay(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input string
		want  string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, lookAndSay(tt.input))
	}
}

func TestDay10(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Details: "Y2015D10 my input",
			Input:   day10myInput,
			Result1: "329356",
			Result2: "4666278"},
	}
	for _, tt := range testCases {
		tt.Test(Day10, assert)
	}
}

func BenchmarkDay10(b *testing.B) {
	internal.Benchmark(Day10, b, day10myInput)
}

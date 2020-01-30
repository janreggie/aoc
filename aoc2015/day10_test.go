package aoc2015

import (
	"testing"

	"github.com/janreggie/AdventOfCode/tools"
	"github.com/stretchr/testify/assert"
)

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
	testCases := []tools.TestCase{
		{Details: "Y2015D10 my input",
			Input:   "3113322113",
			Result1: "329356",
			Result2: "4666278"},
	}
	for _, tt := range testCases {
		tt.Test(Day10, assert)
	}
}

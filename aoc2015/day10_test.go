package aoc2015

import (
	"testing"

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

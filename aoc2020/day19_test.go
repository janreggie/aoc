package aoc2020

import (
	"strings"
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_messageRuleset(t *testing.T) {
	assert := assert.New(t)
	rules := strings.Split(strings.Split(day19sampleInput, "\n\n")[0], "\n") // this works.
	ruleset := newMessageRuleset()
	for _, rule := range rules {
		assert.NoError(ruleset.addRule(rule))
	}
	assert.Equal(messageRuleset{
		{[][]int{{4, 1, 5}}, ""},
		{[][]int{{2, 3}, {3, 2}}, ""},
		{[][]int{{4, 4}, {5, 5}}, ""},
		{[][]int{{4, 5}, {5, 4}}, ""},
		{nil, "a"},
		{nil, "b"},
	}, ruleset)

	// Check if rule 4 matches...
	assert.Equal(true, ruleset.check(4, "a"))
	assert.Equal(false, ruleset.check(4, "b"))

	// Check if rule 3 matches...
	assert.Equal(true, ruleset.check(3, "ab"))
	assert.Equal(true, ruleset.check(3, "ba"))
	assert.Equal(false, ruleset.check(3, "aa"))

	messages := []struct {
		msg   string
		match bool
	}{
		{"ababbb", true},
		{"bababa", false},
		{"abbbab", true},
		{"aaabbb", false},
		{"aaaabbb", false},
	}
	for _, mm := range messages {
		assert.Equal(mm.match, ruleset.check(0, mm.msg))
	}
}

func TestDay19(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{
			Details: "Y2020D19 sample input",
			Input:   day19sampleInput,
			Result1: "3",
			Result2: "12",
		},
		{
			Details: "Y2020D19 my input",
			Input:   day19myInput,
			Result1: "192",
			Result2: "",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day19, assert)
	}
}

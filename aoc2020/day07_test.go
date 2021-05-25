package aoc2020

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_bagRuleset(t *testing.T) {
	assert := assert.New(t)
	ruleset, err := generateBagRuleset(day07sampleInput)
	assert.NoError(err)
	assert.Equal(map[bag][]bagRule{
		{"light", "red"}:    {{1, bag{"bright", "white"}}, {2, bag{"muted", "yellow"}}},
		{"dark", "orange"}:  {{3, bag{"bright", "white"}}, {4, bag{"muted", "yellow"}}},
		{"bright", "white"}: {{1, bag{"shiny", "gold"}}},
		{"muted", "yellow"}: {{2, bag{"shiny", "gold"}}, {9, bag{"faded", "blue"}}},
		{"shiny", "gold"}:   {{1, bag{"dark", "olive"}}, {2, bag{"vibrant", "plum"}}},
		{"dark", "olive"}:   {{3, bag{"faded", "blue"}}, {4, bag{"dotted", "black"}}},
		{"vibrant", "plum"}: {{5, bag{"faded", "blue"}}, {6, bag{"dotted", "black"}}},
		{"faded", "blue"}:   {},
		{"dotted", "black"}: {},
	}, ruleset.rr)

	assert.ElementsMatch(
		[]bag{
			{"bright", "white"},
			{"muted", "yellow"},
		},
		ruleset.findContaining(bag{"shiny", "gold"}),
	)

	assert.ElementsMatch(
		[]bag{
			{"bright", "white"},
			{"muted", "yellow"},
			{"dark", "orange"},
			{"light", "red"},
		},
		ruleset.findAllContaining(bag{"shiny", "gold"}),
	)

	assert.Equal(11, ruleset.countBags(bag{"vibrant", "plum"}))
	assert.Equal(7, ruleset.countBags(bag{"dark", "olive"}))
	assert.Equal(32, ruleset.countBags(bag{"shiny", "gold"}))
}

func TestDay07(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{
			Details: "Y2020D07 sample input",
			Input:   day07sampleInput,
			Result1: "4",
			Result2: "32",
		},
		{
			Details: "Y2020D07 my input",
			Input:   day07myInput,
			Result1: "142",
			Result2: "10219",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day07, assert)
	}
}

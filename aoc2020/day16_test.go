package aoc2020

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_fieldRange(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		input    string
		expected fieldRange
	}{
		{"arrival platform: 36-453 or 473-963", fieldRange{"arrival platform", [2]int{36, 453}, [2]int{473, 963}}},
		{"arrival track: 35-912 or 924-951", fieldRange{"arrival track", [2]int{35, 912}, [2]int{924, 951}}},
		{"class: 39-376 or 396-968", fieldRange{"class", [2]int{39, 376}, [2]int{396, 968}}},
		{"duration: 31-686 or 697-974", fieldRange{"duration", [2]int{31, 686}, [2]int{697, 974}}},
	}

	for _, tt := range testCases {
		actual, err := newFieldRange(tt.input)
		assert.NoError(err)
		assert.Equal(tt.expected, actual)
	}
}

func Test_ticketNotes(t *testing.T) {
	assert := assert.New(t)
	notes, err := generateTicketNotes(bufio.NewScanner(strings.NewReader(day16sampleInput)))
	assert.NoError(err)
	fmt.Println(notes)
	assert.Equal(ticketNotes{
		fields: []fieldRange{
			{"class", [2]int{1, 3}, [2]int{5, 7}},
			{"row", [2]int{6, 11}, [2]int{33, 44}},
			{"seat", [2]int{13, 40}, [2]int{45, 50}},
		},
		yourTicket: []int{7, 1, 14},
		allTickets: [][]int{
			{7, 3, 47},
			{40, 4, 50},
			{55, 2, 20},
			{38, 6, 12},
		},
	}, notes)
	invalidSums := []int{0, 4, 55, 12}
	for ii := range invalidSums {
		assert.Equal(invalidSums[ii], notes.invalidValuesAt(ii))
	}

	const secondExample = `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
`
	notes, err = generateTicketNotes(bufio.NewScanner(strings.NewReader(secondExample)))
	assert.NoError(err)
	notes.discardInvalids() // shouldn't change nearby tickets
	assert.Equal(ticketNotes{
		fields: []fieldRange{
			{"class", [2]int{0, 1}, [2]int{4, 19}},
			{"row", [2]int{0, 5}, [2]int{8, 19}},
			{"seat", [2]int{0, 13}, [2]int{16, 19}},
		},
		yourTicket: []int{11, 12, 13},
		allTickets: [][]int{
			{3, 9, 18},
			{15, 1, 5},
			{5, 14, 9},
		},
	}, notes)

	validRanges := notes.findFieldRanges()
	class, row, seat := fieldRange{"class", [2]int{0, 1}, [2]int{4, 19}},
		fieldRange{"row", [2]int{0, 5}, [2]int{8, 19}},
		fieldRange{"seat", [2]int{0, 13}, [2]int{16, 19}}
	assert.Equal([][]fieldRange{
		{row},
		{class, row},
		{class, row, seat},
	}, validRanges)

	deduced := deducePossible(validRanges)
	assert.Equal([][]fieldRange{{row, class, seat}}, deduced)
}

func TestDay16(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{
			Details: "Y2020D16 sample input",
			Input:   day16sampleInput,
			Result1: "71",
			Result2: "1", // we can't really do anything with it
		},
		{
			Details: "Y2020D16 my input",
			Input:   day16myInput,
			Result1: "18227",
			Result2: "2355350878831",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day16, assert)
	}
}

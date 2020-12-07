package aoc2020

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_customsResponse(t *testing.T) {
	assert := assert.New(t)
	response, err := newCustomsResponse("vknjbcuqrxshzemtayow")
	assert.NoError(err)
	assert.Equal(customsResponse{
		true,  // a
		true,  // b
		true,  // c
		false, // d
		true,  // e
		false, // f
		false, // g
		true,  // h
		false, // i
		true,  // j
		true,  // k
		false, // l
		true,  // m
		true,  // n
		true,  // o
		false, // p
		true,  // q
		true,  // r
		true,  // s
		true,  // t
		true,  // u
		true,  // v
		true,  // w
		true,  // x
		true,  // y
		true,  // z
	}, response)
}

func Test_customsResponseGroups(t *testing.T) {
	assert := assert.New(t)
	responseGroups, err := generateCustomsResponseGroups(bufio.NewScanner(strings.NewReader(day06sampleInput)))
	assert.NoError(err)
	assert.Equal(5, len(responseGroups))
	for ii, grp := range responseGroups {
		fmt.Printf("Group %v with %v members\n", ii, len(grp))
		for _, resp := range grp {
			fmt.Println(resp)
		}
	}

	assert.Equal([]customsResponseGroup{
		{{true, true, true}},
		{{true}, {false, true}, {false, false, true}},
		{{true, true}, {true, false, true}},
		{{true}, {true}, {true}, {true}},
		{{false, true}},
	}, responseGroups)

	atLeastOnes := []customsResponse{
		{true, true, true},
		{true, true, true},
		{true, true, true},
		{true},
		{false, true},
	}
	for ii, grp := range responseGroups {
		assert.Equal(atLeastOnes[ii], grp.atLeastOne())
	}

	everyones := []customsResponse{
		{true, true, true},
		{},
		{true},
		{true},
		{false, true},
	}
	for ii, grp := range responseGroups {
		assert.Equal(everyones[ii], grp.everyone())
	}
}

func TestDay06(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{
			Details: "Y2020D06 sample input",
			Input:   day06sampleInput,
			Result1: "11",
			Result2: "6",
		},
		{
			Details: "Y2020D06 my input",
			Input:   day06myInput,
			Result1: "6416",
			Result2: "3050",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day06, assert)
	}
}

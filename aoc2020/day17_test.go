package aoc2020

import (
	"fmt"
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_conwayCube(t *testing.T) {
	// aliases for later
	T, F := true, false
	assert := assert.New(t)
	cube, err := newConwayCube(day17sampleInput)
	assert.NoError(err)
	assert.Equal(conwayCube{{
		{F, T, F},
		{F, F, T},
		{T, T, T},
	}}, cube)

	surrounded := make([][][]int, 3)
	for zz := range surrounded {
		surrounded[zz] = make([][]int, 5)
		for yy := range surrounded[zz] {
			surrounded[zz][yy] = make([]int, 5)
		}
	}
	for zz := range surrounded {
		for yy := range surrounded[zz] {
			for xx := range surrounded[zz][yy] {
				surrounded[zz][yy][xx] = cube.surrounds(xx-1, yy-1, zz-1)
			}
		}
	}
	assert.Equal([][][]int{
		{
			{0, 1, 1, 1, 0},
			{0, 1, 2, 2, 1},
			{1, 3, 5, 4, 2},
			{1, 2, 4, 3, 2},
			{1, 2, 3, 2, 1},
		}, {
			{0, 1, 1, 1, 0},
			{0, 1, 1, 2, 1},
			{1, 3, 5, 3, 2},
			{1, 1, 3, 2, 2},
			{1, 2, 3, 2, 1},
		},
		{
			{0, 1, 1, 1, 0},
			{0, 1, 2, 2, 1},
			{1, 3, 5, 4, 2},
			{1, 2, 4, 3, 2},
			{1, 2, 3, 2, 1},
		},
	}, surrounded)

	cube = cube.iterate()
	fmt.Println(cube)
	assert.Equal(conwayCube{
		{
			{T, F, F},
			{F, F, T},
			{F, T, F},
		}, {
			{T, F, T},
			{F, T, T},
			{F, T, F},
		}, {
			{T, F, F},
			{F, F, T},
			{F, T, F},
		},
	}, cube)

	hcube, err := newConwayHypercube(day17sampleInput)
	assert.NoError(err)
	assert.Equal(1, hcube.hyperdepth())
	assert.Equal(1, hcube.depth())
	assert.Equal([][]bool{
		{F, T, F},
		{F, F, T},
		{T, T, T},
	}, hcube[0][0])

	hcube = hcube.iterate()
	assert.Equal(conwayHypercube{
		{
			{ // z=-1, w=-1
				{T, F, F},
				{F, F, T},
				{F, T, F},
			},
			{ // z=0, w=-1
				{T, F, F},
				{F, F, T},
				{F, T, F},
			},
			{ // z=1, w=-1
				{T, F, F},
				{F, F, T},
				{F, T, F},
			},
		},
		{
			{ // z=-1, w=0
				{T, F, F},
				{F, F, T},
				{F, T, F},
			},
			{ // z=0, w=0
				{T, F, T},
				{F, T, T},
				{F, T, F},
			},
			{ // z=1, w=-1
				{T, F, F},
				{F, F, T},
				{F, T, F},
			},
		},
		{
			{ // z=-1, w=1
				{T, F, F},
				{F, F, T},
				{F, T, F},
			},
			{ // z=0, w=1
				{T, F, F},
				{F, F, T},
				{F, T, F},
			},
			{ // z=1, w=1
				{T, F, F},
				{F, F, T},
				{F, T, F},
			},
		},
	}, hcube)
}

func TestDay17(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{
			Details: "Y2020D17 sample input",
			Input:   day17sampleInput,
			Result1: "112",
			Result2: "848",
		},
		{
			Details: "Y2020D17 my input",
			Input:   day17myInput,
			Result1: "273",
			Result2: "1504",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day17, assert)
	}
}

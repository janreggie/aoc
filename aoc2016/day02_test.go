package aoc2016

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_number_iterate(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		name    string
		initial keypadNumber
		input   rune
		final   keypadNumber
	}{
		{"UP on 1", 1, 'U', 1},
		{"UP on 2", 2, 'U', 2},
		{"UP on 3", 3, 'U', 3},
		{"UP on 4", 4, 'U', 1},
		{"UP on 5", 5, 'U', 2},
		{"UP on 6", 6, 'U', 3},
		{"UP on 7", 7, 'U', 4},
		{"UP on 8", 8, 'U', 5},
		{"UP on 9", 9, 'U', 6},

		{"DOWN on 1", 1, 'D', 4},
		{"DOWN on 2", 2, 'D', 5},
		{"DOWN on 3", 3, 'D', 6},
		{"DOWN on 4", 4, 'D', 7},
		{"DOWN on 5", 5, 'D', 8},
		{"DOWN on 6", 6, 'D', 9},
		{"DOWN on 7", 7, 'D', 7},
		{"DOWN on 8", 8, 'D', 8},
		{"DOWN on 9", 9, 'D', 9},

		{"LEFT on 1", 1, 'L', 1},
		{"LEFT on 2", 2, 'L', 1},
		{"LEFT on 3", 3, 'L', 2},
		{"LEFT on 4", 4, 'L', 4},
		{"LEFT on 5", 5, 'L', 4},
		{"LEFT on 6", 6, 'L', 5},
		{"LEFT on 7", 7, 'L', 7},
		{"LEFT on 8", 8, 'L', 7},
		{"LEFT on 9", 9, 'L', 8},

		{"RIGHT on 1", 1, 'R', 2},
		{"RIGHT on 2", 2, 'R', 3},
		{"RIGHT on 3", 3, 'R', 3},
		{"RIGHT on 4", 4, 'R', 5},
		{"RIGHT on 5", 5, 'R', 6},
		{"RIGHT on 6", 6, 'R', 6},
		{"RIGHT on 7", 7, 'R', 8},
		{"RIGHT on 8", 8, 'R', 9},
		{"RIGHT on 9", 9, 'R', 9},
	}
	for _, tt := range testCases {
		input := tt.initial
		input.iterate(tt.input)
		assert.Equal(tt.final, input, tt.name, input)
	}
}

func Test_number_iterateDeux(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		name    string
		initial keypadNumber
		input   rune
		final   keypadNumber
	}{
		{"UP on 1", 1, 'U', 1},
		{"UP on 2", 2, 'U', 2},
		{"UP on 3", 3, 'U', 1},
		{"UP on 4", 4, 'U', 4},
		{"UP on 5", 5, 'U', 5},
		{"UP on 6", 6, 'U', 2},
		{"UP on 7", 7, 'U', 3},
		{"UP on 8", 8, 'U', 4},
		{"UP on 9", 9, 'U', 9},
		{"UP on A", 10, 'U', 6},
		{"UP on B", 11, 'U', 7},
		{"UP on C", 12, 'U', 8},
		{"UP on D", 13, 'U', 11},

		{"DOWN on 1", 1, 'D', 3},
		{"DOWN on 2", 2, 'D', 6},
		{"DOWN on 3", 3, 'D', 7},
		{"DOWN on 4", 4, 'D', 8},
		{"DOWN on 5", 5, 'D', 5},
		{"DOWN on 6", 6, 'D', 10},
		{"DOWN on 7", 7, 'D', 11},
		{"DOWN on 8", 8, 'D', 12},
		{"DOWN on 9", 9, 'D', 9},
		{"DOWN on A", 10, 'D', 10},
		{"DOWN on B", 11, 'D', 13},
		{"DOWN on C", 12, 'D', 12},
		{"DOWN on D", 13, 'D', 13},

		{"LEFT on 1", 1, 'L', 1},
		{"LEFT on 2", 2, 'L', 2},
		{"LEFT on 3", 3, 'L', 2},
		{"LEFT on 4", 4, 'L', 3},
		{"LEFT on 5", 5, 'L', 5},
		{"LEFT on 6", 6, 'L', 5},
		{"LEFT on 7", 7, 'L', 6},
		{"LEFT on 8", 8, 'L', 7},
		{"LEFT on 9", 9, 'L', 8},
		{"LEFT on A", 10, 'L', 10},
		{"LEFT on B", 11, 'L', 10},
		{"LEFT on C", 12, 'L', 11},
		{"LEFT on D", 13, 'L', 13},

		{"RIGHT on 1", 1, 'R', 1},
		{"RIGHT on 2", 2, 'R', 3},
		{"RIGHT on 3", 3, 'R', 4},
		{"RIGHT on 4", 4, 'R', 4},
		{"RIGHT on 5", 5, 'R', 6},
		{"RIGHT on 6", 6, 'R', 7},
		{"RIGHT on 7", 7, 'R', 8},
		{"RIGHT on 8", 8, 'R', 9},
		{"RIGHT on 9", 9, 'R', 9},
		{"RIGHT on A", 10, 'R', 11},
		{"RIGHT on B", 11, 'R', 12},
		{"RIGHT on C", 12, 'R', 12},
		{"RIGHT on D", 13, 'R', 13},
	}
	for _, tt := range testCases {
		input := tt.initial
		input.iterateDeux(tt.input)
		assert.Equal(tt.final, input, tt.name, input)
	}

}

func TestDay02(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Input: "ULL\n" +
			"RRDDD\n" +
			"LURDL\n" +
			"UUUUD\n",
			Result1: "1985",
			Result2: "5DB3",
		},
		{Details: "Y2016D02 my input",
			Input:   day02myInput,
			Result1: "76792",
			Result2: "A7AC3"},
	}
	for _, tt := range testCases {
		tt.Test(Day02, assert)
	}
}

func BenchmarkDay02(b *testing.B) {
	aoc.Benchmark(Day01, b, day02myInput)
}

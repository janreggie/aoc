package aoc2016

import (
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_repertoire(t *testing.T) {
	assert := assert.New(t)

	rp := newRepertoire(0)
	assert.EqualValues(0, rp.length(), "length of repertoire must be zero")

	rp.reallocate(3)
	assert.EqualValues(3, rp.length(), "repertoire reallocated")

	rp.add("eedadn")
	assert.Equal(repertoire{
		{'e': 1},
		{'e': 1},
		{'d': 1},
		{'a': 1},
		{'d': 1},
		{'n': 1},
	}, rp, "added one string")

	rp.add("eedbdxnc")
	assert.Equal(repertoire{
		{'e': 2},
		{'e': 2},
		{'d': 2},
		{'a': 1, 'b': 1},
		{'d': 2},
		{'n': 1, 'x': 1},
		{'n': 1},
		{'c': 1},
	}, rp, "added another string")

	rp.add("abcbex")
	assert.Equal(repertoire{
		{'e': 2, 'a': 1},
		{'e': 2, 'b': 1},
		{'d': 2, 'c': 1},
		{'b': 2, 'a': 1},
		{'d': 2, 'e': 1},
		{'x': 2, 'n': 1},
		{'n': 1},
		{'c': 1},
	}, rp, "added third string")
	assert.Equal("eedbdxnc", rp.mostCommon(), "most common error corrected version")
	assert.Equal("abcaennc", rp.leastCommon(), "least common error corrected version")
}

func TestDay06(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Details: "Y2016D06 sample input",
			Input:   day06sampleInput,
			Result1: "easter",
			Result2: "advent"},
		{Details: "Y2016D06 my input",
			Input:   day06myInput,
			Result1: "usccerug",
			Result2: "cnvvtafc"},
	}
	for _, tt := range testCases {
		tt.Test(Day06, assert)
	}
}

func BenchmarkDay06(b *testing.B) {
	internal.Benchmark(Day06, b, day06myInput)
}

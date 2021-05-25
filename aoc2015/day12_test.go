package aoc2015

import (
	"testing"

	"github.com/antonholmquist/jason"
	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_extractSum(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		value string // to be marshalled using NewValueFromBytes
		want  int64
	}{
		// TODO: Add test cases.
		{"int64", `43`, 43},
		{"array", `[1,2,3]`, 6},
		{"object", `{"a":2,"b":4}`, 6},
		{"nested array", `[[[3]]]`, 3},
		{"nested object", `{"a":{"b":4},"c":-1}`, 3},
		{"object with array value", `{"a":[-1,1]}`, 0},
		{"array of int and object", `[-1,{"a":1}]`, 0},
	}
	for _, tt := range tests {
		value, err := jason.NewValueFromBytes([]byte(tt.value))
		assert.NoError(err, tt.name)
		assert.Equal(tt.want, extractSum(value), tt.name)
	}
}

func Test_extractSumButRed(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		value string // to be marshalled using NewValueFromBytes
		want  int64
	}{
		// TODO: Add test cases.
		{"int64", `43`, 43},
		{"array but red object", `[1,{"c":"red","b":2},3]`, 4},
		{"object with red", `{"d":"red","e":[1,2,3,4],"f":5}`, 0},
		{"array with red string", `[1,"red",5]`, 6},
	}
	for _, tt := range tests {
		value, err := jason.NewValueFromBytes([]byte(tt.value))
		assert.NoError(err, tt.name)
		assert.Equal(tt.want, extractSumButRed(value), tt.name)
	}
}
func TestDay12(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Details: "Y2015D12 my input",
			Input:   day12myInput,
			Result1: "111754",
			Result2: "65402"},
	}
	for _, tt := range testCases {
		tt.Test(Day12, assert)
	}
}

func BenchmarkDay12(b *testing.B) {
	aoc.Benchmark(Day12, b, day12myInput)
}

package aoc2020

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_countAdapterArrrangement(t *testing.T) {
	assert := assert.New(t)
	adapters := make([]int, 0)
	for _, vv := range strings.Split(day10sampleInput, "\n") {
		jolts, err := strconv.Atoi(vv)
		assert.NoError(err)
		adapters = append(adapters, jolts)
	}
	sort.Ints(adapters)
	masterChain := make([]int, len(adapters)+2)
	masterChain[0] = 0
	copy(masterChain[1:], adapters)
	masterChain[len(masterChain)-1] = adapters[len(adapters)-1] + 3
	fmt.Println("masterChain:", masterChain)
	assert.Equal(19208, countAdapterArrangement(masterChain))
}

func TestDay10(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{
			Details: "Y2020D10 sample input",
			Input:   day10sampleInput,
			Result1: "220",
			Result2: "",
		},
		{
			Details: "Y2020D10 my input",
			Input:   day10myInput,
			Result1: "2470",
			Result2: "",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day10, assert)
	}
}

func Benchmark_countAdapterArrangement(b *testing.B) {
	adapters := make([]int, 0)
	for _, vv := range strings.Split(day10sampleInput, "\n") {
		jolts, _ := strconv.Atoi(vv) // assert to be true
		adapters = append(adapters, jolts)
	}
	sort.Ints(adapters)
	masterChain := make([]int, len(adapters)+2)
	masterChain[0] = 0
	copy(masterChain[1:], adapters)
	masterChain[len(masterChain)-1] = adapters[len(adapters)-1] + 3

	for ii := 0; ii < b.N; ii++ {
		countAdapterArrangement(masterChain)
	}
}

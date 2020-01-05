package aoc2015

import (
	"bufio"
	"strings"
	"testing"

	"github.com/janreggie/AdventOfCode/structs"
	"github.com/stretchr/testify/assert"
)

func TestDay09(t *testing.T) {
	assert := assert.New(t)
	testCases := []structs.TestCase{
		{Input: "London to Dublin = 464\n" +
			"London to Belfast = 518\n" +
			"Dublin to Belfast = 141\n",
			Result1: "605",
			Result2: "804"},
		{Details: "Y2019D09 my input",
			Input: "Faerun to Norrath = 129\n" +
				"Faerun to Tristram = 58\n" +
				"Faerun to AlphaCentauri = 13\n" +
				"Faerun to Arbre = 24\n" +
				"Faerun to Snowdin = 60\n" +
				"Faerun to Tambi = 71\n" +
				"Faerun to Straylight = 67\n" +
				"Norrath to Tristram = 142\n" +
				"Norrath to AlphaCentauri = 15\n" +
				"Norrath to Arbre = 135\n" +
				"Norrath to Snowdin = 75\n" +
				"Norrath to Tambi = 82\n" +
				"Norrath to Straylight = 54\n" +
				"Tristram to AlphaCentauri = 118\n" +
				"Tristram to Arbre = 122\n" +
				"Tristram to Snowdin = 103\n" +
				"Tristram to Tambi = 49\n" +
				"Tristram to Straylight = 97\n" +
				"AlphaCentauri to Arbre = 116\n" +
				"AlphaCentauri to Snowdin = 12\n" +
				"AlphaCentauri to Tambi = 18\n" +
				"AlphaCentauri to Straylight = 91\n" +
				"Arbre to Snowdin = 129\n" +
				"Arbre to Tambi = 53\n" +
				"Arbre to Straylight = 40\n" +
				"Snowdin to Tambi = 15\n" +
				"Snowdin to Straylight = 99\n" +
				"Tambi to Straylight = 70\n",
			Result1: "207",
			Result2: "804",
		},
	}
	for _, eachCase := range testCases {
		eachCase.Test(Day09, assert)
	}
}

func Benchmark_graph_findMinimumHamiltonian(b *testing.B) {
	const input = "Faerun to Norrath = 129\n" +
		"Faerun to Tristram = 58\n" +
		"Faerun to AlphaCentauri = 13\n" +
		"Faerun to Arbre = 24\n" +
		"Faerun to Snowdin = 60\n" +
		"Faerun to Tambi = 71\n" +
		"Faerun to Straylight = 67\n" +
		"Norrath to Tristram = 142\n" +
		"Norrath to AlphaCentauri = 15\n" +
		"Norrath to Arbre = 135\n" +
		"Norrath to Snowdin = 75\n" +
		"Norrath to Tambi = 82\n" +
		"Norrath to Straylight = 54\n" +
		"Tristram to AlphaCentauri = 118\n" +
		"Tristram to Arbre = 122\n" +
		"Tristram to Snowdin = 103\n" +
		"Tristram to Tambi = 49\n" +
		"Tristram to Straylight = 97\n" +
		"AlphaCentauri to Arbre = 116\n" +
		"AlphaCentauri to Snowdin = 12\n" +
		"AlphaCentauri to Tambi = 18\n" +
		"AlphaCentauri to Straylight = 91\n" +
		"Arbre to Snowdin = 129\n" +
		"Arbre to Tambi = 53\n" +
		"Arbre to Straylight = 40\n" +
		"Snowdin to Tambi = 15\n" +
		"Snowdin to Straylight = 99\n" +
		"Tambi to Straylight = 70\n"
	gr, _ := newGraph(bufio.NewScanner(strings.NewReader(input))) // guaranteed no error
	for ii := 0; ii < b.N; ii++ {
		gr.findMinimumHamiltonian()
	}
}

func Benchmark_graph_findMaximumHamiltonian(b *testing.B) {
	const input = "Faerun to Norrath = 129\n" +
		"Faerun to Tristram = 58\n" +
		"Faerun to AlphaCentauri = 13\n" +
		"Faerun to Arbre = 24\n" +
		"Faerun to Snowdin = 60\n" +
		"Faerun to Tambi = 71\n" +
		"Faerun to Straylight = 67\n" +
		"Norrath to Tristram = 142\n" +
		"Norrath to AlphaCentauri = 15\n" +
		"Norrath to Arbre = 135\n" +
		"Norrath to Snowdin = 75\n" +
		"Norrath to Tambi = 82\n" +
		"Norrath to Straylight = 54\n" +
		"Tristram to AlphaCentauri = 118\n" +
		"Tristram to Arbre = 122\n" +
		"Tristram to Snowdin = 103\n" +
		"Tristram to Tambi = 49\n" +
		"Tristram to Straylight = 97\n" +
		"AlphaCentauri to Arbre = 116\n" +
		"AlphaCentauri to Snowdin = 12\n" +
		"AlphaCentauri to Tambi = 18\n" +
		"AlphaCentauri to Straylight = 91\n" +
		"Arbre to Snowdin = 129\n" +
		"Arbre to Tambi = 53\n" +
		"Arbre to Straylight = 40\n" +
		"Snowdin to Tambi = 15\n" +
		"Snowdin to Straylight = 99\n" +
		"Tambi to Straylight = 70\n"
	gr, _ := newGraph(bufio.NewScanner(strings.NewReader(input))) // guaranteed no error
	for ii := 0; ii < b.N; ii++ {
		gr.findMaximumHamiltonian()
	}
}

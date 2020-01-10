package aoc2015

import (
	"bufio"
	"strings"
	"testing"

	"github.com/janreggie/AdventOfCode/structs"
	"github.com/stretchr/testify/assert"
)

var day09sampleGraph = func() graph {
	const input = "London to Dublin = 464\n" +
		"London to Belfast = 518\n" +
		"Dublin to Belfast = 141\n"

	// guaranteed not to error
	gr, _ := newGraph(bufio.NewScanner(strings.NewReader(input)))
	return gr
}()

var day09myGraph = func() graph {
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

	// guaranteed not to error
	gr, _ := newGraph(bufio.NewScanner(strings.NewReader(input)))
	return gr
}()

func Test_graph_get(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		gr   *graph
		a, b string
		want uint
	}{
		{&day09sampleGraph, "London", "Belfast", 518},
		{&day09sampleGraph, "Belfast", "London", 518},
		{&day09sampleGraph, "London", "London", 0},
		{&day09sampleGraph, "London", "Somewhere", disconnected},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, tt.gr.get(tt.a, tt.b), tt.a, tt.b, tt.want)
	}
}

func Test_graph_distance(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		gr    *graph
		towns []string
		want  uint
	}{
		{&day09sampleGraph, []string{"Dublin", "London", "Belfast"}, 982},
		{&day09sampleGraph, []string{"London", "Dublin", "Belfast"}, 605},
		{&day09sampleGraph, []string{"London", "Belfast", "Dublin"}, 659},
		{&day09sampleGraph, []string{"Dublin", "Belfast", "London"}, 659},
		{&day09sampleGraph, []string{"Belfast", "Dublin", "London"}, 605},
		{&day09sampleGraph, []string{"Belfast", "London", "Dublin"}, 982},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, tt.gr.distanceSimple(tt.towns), tt.towns)
	}
}

func Test_graph_greedyShortestPathFrom(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		gr   *graph
		town string
		want uint
	}{
		{&day09sampleGraph, "Dublin", 659},
		{&day09sampleGraph, "London", 605},
		{&day09sampleGraph, "Belfast", 605},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, tt.gr.shortestPathGreedyFrom(tt.town), tt.town)
	}
}

func Test_graph_shortestPathGreedy(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		gr   *graph
		want uint
	}{
		{&day09sampleGraph, 605},
		{&day09myGraph, 207}, // the shortest one we have.
	}
	for _, tt := range tests {
		assert.Equal(tt.want, tt.gr.shortestPathGreedy(), tt.gr)
	}
}

func TestDay09(t *testing.T) {
	assert := assert.New(t)
	testCases := []structs.TestCase{
		{Input: "London to Dublin = 464\n" +
			"London to Belfast = 518\n" +
			"Dublin to Belfast = 141\n",
			Result1: "605",
			Result2: "982"},
		{Details: "Y2015D09 my input",
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

func BenchmarkDay09(b *testing.B) {
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
	for ii := 0; ii < b.N; ii++ {
		Day09(bufio.NewScanner(strings.NewReader(input)))
	}
}

func Benchmark_graph_shortestPathPermutative(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		day09myGraph.shortestPathPermutative()
	}
}

func Benchmark_graph_shortestPathGreedy(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		day09myGraph.shortestPathGreedy()
	}
}

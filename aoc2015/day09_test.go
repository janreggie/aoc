package aoc2015

import (
	"bufio"
	"strings"
	"testing"

	"github.com/janreggie/AdventOfCode/tools"
	"github.com/stretchr/testify/assert"
)

var day09sampleGraph = func() townGraph {
	const input = "London to Dublin = 464\n" +
		"London to Belfast = 518\n" +
		"Dublin to Belfast = 141\n"

	// guaranteed not to error
	gr, _ := newTownGraph(bufio.NewScanner(strings.NewReader(input)))
	return gr
}()

var day09myGraph = func() townGraph {
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
	gr, _ := newTownGraph(bufio.NewScanner(strings.NewReader(input)))
	return gr
}()

func Test_graph_get(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		gr   *townGraph
		a, b town
		want townDistance
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
		gr    *townGraph
		towns []town
		want  townDistance
	}{
		{&day09sampleGraph, []town{"Dublin", "London", "Belfast"}, 982},
		{&day09sampleGraph, []town{"London", "Dublin", "Belfast"}, 605},
		{&day09sampleGraph, []town{"London", "Belfast", "Dublin"}, 659},
		{&day09sampleGraph, []town{"Dublin", "Belfast", "London"}, 659},
		{&day09sampleGraph, []town{"Belfast", "Dublin", "London"}, 605},
		{&day09sampleGraph, []town{"Belfast", "London", "Dublin"}, 982},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, tt.gr.distanceSimple(tt.towns), tt.towns)
	}
}

func Test_graph_shortestPathGreedyFrom(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		gr     *townGraph
		origin town
		want   townDistance
	}{
		{&day09sampleGraph, "Dublin", 659},
		{&day09sampleGraph, "London", 605},
		{&day09sampleGraph, "Belfast", 605},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, tt.gr.shortestPathGreedyFrom(tt.origin), tt.origin)
	}
}

func Test_graph_shortestPathCleverFrom(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		gr     *townGraph
		origin town
		want   townDistance
	}{
		{&day09sampleGraph, "Dublin", 659},
		{&day09sampleGraph, "London", 605},
		{&day09sampleGraph, "Belfast", 605},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, tt.gr.shortestPathCleverFrom(tt.origin), tt.origin)
	}
}

func Test_graph_shortestPathGreedy(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		gr   *townGraph
		want townDistance
	}{
		{&day09sampleGraph, 605},
		// {&day09myGraph, 207}, // the shortest one we have. may not work.
	}
	for _, tt := range tests {
		assert.Equal(tt.want, tt.gr.shortestPathGreedy(), tt.gr)
	}
}

func Test_graph_shortestPathClever(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		gr   *townGraph
		want townDistance
	}{
		{&day09sampleGraph, 605},
		{&day09myGraph, 207}, // the shortest one we have.
	}
	for _, tt := range tests {
		assert.Equal(tt.want, tt.gr.shortestPathClever(), tt.gr)
	}
}
func Test_graph_shortestPathPermutative(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		gr   *townGraph
		want townDistance
	}{
		{&day09sampleGraph, 605},
		{&day09myGraph, 207}, // the shortest one we have.
	}
	for _, tt := range tests {
		assert.Equal(tt.want, tt.gr.shortestPathPermutative(), tt.gr)
	}
}
func Test_graph_longestPathGreedy(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		gr   *townGraph
		want townDistance
	}{
		{&day09sampleGraph, 982},
		// {&day09myGraph, 804}, // the longest one we have. may not work.
	}
	for _, tt := range tests {
		assert.Equal(tt.want, tt.gr.longestPathGreedy(), tt.gr)
	}
}
func Test_graph_longestPathNotSoCleverFrom(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		gr     *townGraph
		origin town
		want   townDistance
	}{
		{&day09sampleGraph, "Dublin", 982},
		{&day09sampleGraph, "London", 659},
		{&day09sampleGraph, "Belfast", 982},
	}
	for _, tt := range tests {
		assert.Equal(tt.want, tt.gr.longestPathNotSoCleverFrom(tt.origin), tt.origin)
	}
}

func Test_graph_longestPathNotSoClever(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		gr   *townGraph
		want townDistance
	}{
		{&day09sampleGraph, 982},
		{&day09myGraph, 804}, // the longest one we have.
	}
	for _, tt := range tests {
		assert.Equal(tt.want, tt.gr.longestPathNotSoClever(), tt.gr)
	}
}
func Test_graph_longestPathPermutative(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		gr   *townGraph
		want townDistance
	}{
		{&day09sampleGraph, 982},
		{&day09myGraph, 804}, // the longest one we have.
	}
	for _, tt := range tests {
		assert.Equal(tt.want, tt.gr.longestPathPermutative(), tt.gr)
	}
}

func Test_townPath_newAndCopy(t *testing.T) {
	assert := assert.New(t)
	gr := &day09sampleGraph
	tests := []struct {
		a, b      town
		wantErr   bool
		distance  townDistance
		remaining []town
	}{
		{"London", "Dublin", false, 464, []town{"Belfast"}},
		{"London", "Nowhere", true, disconnected, []town{}},
		{"Nowhere", "London", true, disconnected, []town{}},
	}
	for _, tt := range tests {
		path, err := newTownPath(tt.a, tt.b, gr)
		if tt.wantErr {
			assert.Error(err, tt)
			continue
		}
		assert.NoError(err, tt)
		assert.Equal([]town{tt.a, tt.b}, path.raw, tt)
		assert.Equal(tt.distance, path.distance, tt)
		assert.Equal(tt.remaining, path.remaining, tt)
		// now for copying
		otherPath := path.copy()
		assert.Equal([]town{tt.a, tt.b}, otherPath.raw, tt)
		assert.Equal(tt.distance, otherPath.distance, tt)
		assert.Equal(tt.remaining, otherPath.remaining, tt)
	}
}

func Test_townPath_add(t *testing.T) {
	assert := assert.New(t)
	gr := &day09sampleGraph

	testCases := []struct {
		next      town
		wantErr   bool
		distance  townDistance
		remaining []town
	}{
		{"Dublin", false, 659, []town{}},
		{"Nowhere", true, disconnected, []town{}},
		{"London", true, disconnected, []town{}},
	}
	for _, tt := range testCases {
		path, err := newTownPath("London", "Belfast", gr)
		assert.NoError(err)
		newPath, err := path.add(tt.next)
		if tt.wantErr {
			assert.Error(err, tt)
			continue
		}
		assert.NoError(err, tt)
		assert.Equal(tt.distance, newPath.distance, tt)
		assert.Equal(tt.remaining, newPath.remaining, tt)
	}
}

func Test_townPathQueue_pushAndPop(t *testing.T) {
	assert := assert.New(t)
	gr := &day09sampleGraph

	// let's initialize the path
	LD, err := newTownPath("London", "Dublin", gr) // LD: 464
	assert.NoError(err, gr)
	LDB, err := LD.add("Belfast") // LDB: 605
	assert.NoError(err, LD)

	// let's make that queue
	tpq := newTownPathQueue()
	tpq.push(LD) // tpq: [LondonDublin]
	assert.ElementsMatch(townPathQueue{LD}, tpq, tpq)
	tpq.push(LDB) // tpq: [LondonDublin, LondonDublinBelfast]
	assert.ElementsMatch(townPathQueue{LD, LDB}, tpq, tpq)

	// now let's another one
	DB, err := newTownPath("Dublin", "Belfast", gr) // DB: 141
	assert.NoError(err, gr)
	DBL, err := DB.add("London") // DBL: 659
	assert.NoError(err, DB)
	tpq.push(DB) // tpq: [DB, LD, LDB]
	assert.ElementsMatch(townPathQueue{DB, LD, LDB}, tpq, tpq)
	tpq.push(DBL) // tpq: [DB, LD, LDB, DBL]
	assert.ElementsMatch(townPathQueue{DB, LD, LDB, DBL}, tpq, tpq)

	// now let's start popping...
	pop, err := tpq.pop() // pop, tpq: DB, [LD, LDB, DBL]
	assert.NoError(err, tpq)
	assert.Equal(DB, pop, tpq)
	assert.ElementsMatch(townPathQueue{LD, LDB, DBL}, tpq, tpq)

	pop, err = tpq.pop() // pop, tpq: LD, [LDB, DBL]
	assert.NoError(err, tpq)
	assert.Equal(LD, pop, tpq)
	assert.ElementsMatch(townPathQueue{LDB, DBL}, tpq, tpq)

	pop, err = tpq.pop() // pop, tpq: LDB, [DBL]
	assert.NoError(err, tpq)
	assert.Equal(LDB, pop, tpq)
	assert.ElementsMatch(townPathQueue{DBL}, tpq, tpq)

	pop, err = tpq.pop() // pop, tpq: DBL, []
	assert.NoError(err, tpq)
	assert.Equal(DBL, pop, tpq)
	assert.ElementsMatch(townPathQueue{}, tpq, tpq)

	pop, err = tpq.pop() // can pop no more
	assert.Error(err, tpq)
}

func TestDay09(t *testing.T) {
	assert := assert.New(t)
	testCases := []tools.TestCase{
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
	for _, tt := range testCases {
		tt.Test(Day09, assert)
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

func Benchmark_graph_shortestPathClever(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		day09myGraph.shortestPathClever()
	}
}
func Benchmark_graph_longestPathPermutative(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		day09myGraph.longestPathPermutative()
	}
}

func Benchmark_graph_longestPathGreedy(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		day09myGraph.longestPathGreedy()
	}
}

func Benchmark_graph_longestPathNotSoClever(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		day09myGraph.longestPathNotSoClever()
	}
}

package aoc2015

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// exclude returns the same slice of towns but without some specified value.
func exclude(from []town, except town) []town {
	out := make([]town, 0, len(from)-1)
	for _, vv := range from {
		if vv != except {
			out = append(out, vv)
		}
	}
	return out
}

// disconnected is a townDistance that is unfathomably large.
const disconnected = townDistance(math.MaxUint32)

// town represents a town
type town string

// townPair represents a pair of two towns
// where a is lexographically less than or equal to b.
type townPair struct {
	a, b town
}

// townDistance represents the distance between towns
type townDistance uint

// townGraph represents a graph of towns
type townGraph struct {
	links map[townPair]townDistance // pls no negative distance
	towns []town
}

// townPath is a path that connects multiple towns together.
// The path can only be added to.
type townPath struct {
	raw      []town
	distance townDistance
}

// newTownPath creates a new townPath construct
// that starts with two paths.

// newGraph creates a graph construct
func newGraph(scanner *bufio.Scanner) (townGraph, error) {
	out := townGraph{links: make(map[townPair]townDistance), towns: make([]town, 0)}

	// now fill it in for each line
	for scanner.Scan() {
		// assume format is
		// `AlphaNumericStringA to AlphaNumericStringB = UnsignedInteger`
		text := strings.Fields(scanner.Text())
		if len(text) != 5 {
			return townGraph{}, fmt.Errorf("could not parse %v", scanner.Text())
		}
		townOne, townTwo := town(text[0]), town(text[2])
		dist, err := strconv.Atoi(text[4])
		if err != nil {
			return townGraph{}, errors.Wrapf(err, "could not convert %v in %v", text[4], scanner.Text())
		}
		if dist < 0 {
			return townGraph{}, errors.Wrapf(err, "could not fathom distance %v in %v", dist, scanner.Text())
		}
		out.set(townOne, townTwo, townDistance(dist))
	}

	return out, nil
}

// set adds the edge that connects a and b with distance dist.
// Make sure that neither a nor b contain \x00.
// If a == b then it doesn't do anything.
func (tg *townGraph) set(a, b town, dist townDistance) {
	if a == b {
		return
	}
	if a > b {
		a, b = b, a
	}
	tg.links[townPair{a, b}] = dist
	writeA, writeB := true, true
	for _, v := range tg.towns {
		if v == a {
			writeA = false
		}
		if v == b {
			writeB = false
		}
	}
	if writeA {
		tg.towns = append(tg.towns, a)
	}
	if writeB {
		tg.towns = append(tg.towns, b)
	}
}

// get returns the distance between two towns.
// If a == b then return 0.
// If link doesn't exist then return disconnected.
func (tg townGraph) get(a, b town) townDistance {
	if a == b {
		return 0
	}
	if a > b {
		a, b = b, a
	}
	if dist, ok := tg.links[townPair{a, b}]; ok {
		return dist
	}
	return disconnected
}

// in returns true if town is in the graph's list of towns.
func (tg *townGraph) in(town town) bool {
	for _, tt := range tg.towns {
		if tt == town {
			return true
		}
	}
	return false
}

// isValid returns whether the path is a isValid path in graph.
// If path is empty it will return true.
func (tg *townGraph) isValid(path []town) bool {
	if len(path) == 0 {
		return true
	}
	if len(path) == 1 {
		return tg.in(path[0])
	}

	firstTwo := tg.get(path[0], path[1])
	if len(path) == 2 {
		return firstTwo != disconnected
	}

	return tg.isValid(path[1:])
}

// distance determines the distance returned after traversing all towns in path.
// Will also return a boolean signifying if path is a valid path.
// If path is not valid it will return disconnected and false.
// If path is empty it will return 0 and true.
func (tg *townGraph) distance(path []town) (townDistance, bool) {
	if len(path) == 0 {
		return 0, true
	}
	if len(path) == 1 {
		return 0, tg.in(path[0])
	}

	firstTwo := tg.get(path[0], path[1])
	if len(path) == 2 {
		return firstTwo, firstTwo != disconnected
	}

	remaining, areValid := tg.distance(path[1:])
	if !areValid {
		return disconnected, false
	}
	return firstTwo + remaining, true
}

// distanceSimple traces the distance returned after traversing all towns in path.
// Does not check if the path is even valid.
// If it isn't then distanceSimple may return a number that is at least disconnectedLink.
func (tg *townGraph) distanceSimple(path []town) townDistance {
	if len(path) <= 1 {
		return 0 // there is no distance to be travelled
	}
	if len(path) == 2 {
		return tg.get(path[0], path[1])
	}
	return tg.get(path[0], path[1]) + tg.distanceSimple(path[1:])
}

// permutations creates a channel that contains all the possible paths in the graph
// and will close said channel.
func (tg *townGraph) permutations() <-chan []town {
	permutate := func(c chan []town, inputs []town) {
		output := make([]town, len(inputs))
		copy(output, inputs)
		c <- output

		size := len(inputs)
		p := make([]int, size+1)
		for i := 0; i < size+1; i++ {
			p[i] = i
		}
		for i := 1; i < size; {
			p[i]--
			j := 0
			if i%2 == 1 {
				j = p[i]
			}
			tmp := inputs[j]
			inputs[j] = inputs[i]
			inputs[i] = tmp
			output := make([]town, len(inputs))
			copy(output, inputs)
			c <- output
			for i = 1; p[i] == 0; i++ {
				p[i] = i
			}
		}
	}

	c := make(chan []town)
	go func(c chan []town) {
		defer close(c)
		permutate(c, tg.towns)
	}(c)
	return c
}

// shortestPathPermutative returns the shortest path distance
// by checking through all permutations of the towns in the graph.
// This is an expensive operation.
// If gr has no valid paths it will return disconnected.
func (tg *townGraph) shortestPathPermutative() townDistance {
	allPermuts := tg.permutations()

	record := disconnected
	for path := range allPermuts {
		// guaranteed to be valid.
		if pathDist := tg.distanceSimple(path); pathDist < record {
			record = pathDist
		}
	}

	return record
}

// longestPathPermutative returns the longest path distance
// by checking through all permutations of the towns in the graph.
// This is an expensive operation.
// If gr has no valid paths it will return disconnected.
func (tg *townGraph) longestPathPermutative() townDistance {
	allPermuts := tg.permutations()

	record := townDistance(0)
	for path := range allPermuts {
		// guaranteed to be valid
		if pathDist := tg.distanceSimple(path); pathDist > record {
			record = pathDist
		}
	}

	if record == 0 {
		return disconnected
	}
	return record
}

// shortestPathGreedyFrom returns the shortest path distance from some town.
// If town does not exist return MaxUint32.
// If the graph is not complete it may perform undefined behavior.
// This function has issues (see below).
//
// Rationale:
//
// How about instead of permutating all values we do something a bit better...
// Suppose there are seven towns A to G that Santa has to go through.
// Suppose Santa would start at town A.
// From here, determine which town from the six left is the closest.
// Suppose town B is the closest. Then the distance travelled so far would be A+B.
// Continue with the other towns until all towns would have been used.
//
// The shortest travelled path from town A to all the other towns
// will be what shortestPathGreedyFrom will return
//
// Such a greedy algorithm may run into issues.
// Consider a graph that is defined by the following:
//
//	A to B = 1
//	A to C = 4
//	A to D = 4
//	B to C = 1
//	B to D = 10
//	C to D = 100
//
// Suppose we start at A.
// The shortest path from A would be A -> B -> C -> D, whose total distance is 102.
// This is not the best choice as it is possible to use A -> C -> B -> D,
// whose total distance is 15.
// If the path starts from C however, the path taken would have been C -> B -> A -> D,
// which would have a total distance of 6.
func (tg *townGraph) shortestPathGreedyFrom(from town) townDistance {
	// check first if town is in towns
	if !tg.in(from) {
		return disconnected
	}

	remaining := exclude(tg.towns, from)
	current := from
	var total townDistance

	for len(remaining) > 0 {
		// check distances from current to everything in remaining
		recordDist, recordTown := disconnected, town("")
		for _, eachDestination := range remaining {
			if eachDistance := tg.get(current, eachDestination); eachDistance < recordDist {
				recordDist, recordTown = eachDistance, eachDestination
			}
		}
		// now recordTown sounds like the winner.
		remaining = exclude(remaining, recordTown)
		current = recordTown
		total += recordDist
	}

	return total
}

// shortestPathGreedy checks all towns in a graph's towns list to see which is the shortest to go to.
func (tg *townGraph) shortestPathGreedy() townDistance {
	record := disconnected
	for _, town := range tg.towns {
		if pathDistance := tg.shortestPathGreedyFrom(town); pathDistance < record {
			record = pathDistance
		}
	}
	return record
}

// shortestPathCleverFrom determines the shortest path from a town
// using an algorithm a bit smarter than the previous one,
// although this algorithm is still exhaustive.
// If there is no valid path from town, return disconnected.
func (tg *townGraph) shortestPathCleverFrom(town town) townDistance {
	// do things...

	return 0
}

// Day09 solves the ninth day puzzle
// "All in a Single Night"
func Day09(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	santasGraph, err := newGraph(scanner)
	if err != nil {
		return
	}

	answer1 = strconv.FormatUint(uint64(santasGraph.shortestPathPermutative()), 10)
	answer2 = strconv.FormatUint(uint64(santasGraph.longestPathPermutative()), 10)

	return
}

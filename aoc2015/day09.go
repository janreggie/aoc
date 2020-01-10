package aoc2015

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

// stringSliceBut returns the same slice of strings but without a specified value.
func stringSliceBut(from []string, except string) []string {
	out := make([]string, 0, len(from)-1)
	for _, vv := range from {
		if vv != except {
			out = append(out, vv)
		}
	}
	return out
}

// graph represents a graph
type graph struct {
	links map[string]uint // pls no negative distance
	towns []string
}

// disconnected is a distance that is unfathomably large
const disconnected = uint(math.MaxUint32)

func newGraph(scanner *bufio.Scanner) (graph, error) {
	out := graph{links: make(map[string]uint), towns: make([]string, 0)}

	// now fill it in for each line
	for scanner.Scan() {
		// assume format is
		// `AlphaNumericStringA to AlphaNumericStringB = UnsignedInteger`
		text := strings.Fields(scanner.Text())
		if len(text) != 5 {
			return graph{}, fmt.Errorf("could not parse %v", scanner.Text())
		}
		townOne, townTwo := text[0], text[2]
		distance, err := strconv.Atoi(text[4])
		if err != nil {
			return graph{}, errors.Wrapf(err, "could not convert %v in %v", text[4], scanner.Text())
		}
		if distance < 0 {
			return graph{}, errors.Wrapf(err, "could not fathom distance %v in %v", distance, scanner.Text())
		}
		out.set(townOne, townTwo, uint(distance))
	}

	return out, nil
}

// set adds the edge that connects a and b with distance dist.
// Make sure that neither a nor b contain \x00.
// If a == b then it doesn't do anything.
func (gr *graph) set(a, b string, dist uint) {
	if a == b {
		return
	}
	if a > b {
		a, b = b, a
	}
	gr.links[a+"\x00"+b] = dist
	writeA, writeB := true, true
	for _, v := range gr.towns {
		if v == a {
			writeA = false
		}
		if v == b {
			writeB = false
		}
	}
	if writeA {
		gr.towns = append(gr.towns, a)
	}
	if writeB {
		gr.towns = append(gr.towns, b)
	}
}

// get returns the distance between two towns.
// If a == b then return 0.
// If link doesn't exist then return disconnected.
func (gr graph) get(a, b string) uint {
	if a == b {
		return 0
	}
	if a > b {
		a, b = b, a
	}
	link := a + "\x00" + b
	if dist, ok := gr.links[link]; ok {
		return dist
	}
	return disconnected
}

// in returns true if town is in the graph's list of towns.
func (gr *graph) in(town string) bool {
	for _, tt := range gr.towns {
		if tt == town {
			return true
		}
	}
	return false
}

// isValid returns whether the path is a isValid path in graph.
// If path is empty it will return true.
func (gr *graph) isValid(path []string) bool {
	if len(path) == 0 {
		return true
	}
	if len(path) == 1 {
		return gr.in(path[0])
	}

	firstTwo := gr.get(path[0], path[1])
	if len(path) == 2 {
		return firstTwo != disconnected
	}

	return gr.isValid(path[1:])
}

// distance determines the distance returned after traversing all towns in path.
// Will also return a boolean signifying if path is a valid path.
// If path is not valid it will return disconnected and false.
// If path is empty it will return 0 and true.
func (gr *graph) distance(path []string) (uint, bool) {
	if len(path) == 0 {
		return 0, true
	}
	if len(path) == 1 {
		return 0, gr.in(path[0])
	}

	firstTwo := gr.get(path[0], path[1])
	if len(path) == 2 {
		return firstTwo, firstTwo != disconnected
	}

	remaining, areValid := gr.distance(path[1:])
	if !areValid {
		return disconnected, false
	}
	return firstTwo + remaining, true
}

// distanceSimple traces the distance returned after traversing all towns in path.
// Does not check if the path is even valid.
// If it isn't then distanceSimple may return a number that is at least disconnectedLink.
func (gr *graph) distanceSimple(path []string) uint {
	if len(path) <= 1 {
		return 0 // there is no distance to be travelled
	}
	if len(path) == 2 {
		return gr.get(path[0], path[1])
	}
	return gr.get(path[0], path[1]) + gr.distanceSimple(path[1:])
}

// permutations creates a channel that contains all the possible paths in the graph
// and will close said channel.
func (gr *graph) permutations() <-chan []string {
	permutate := func(c chan []string, inputs []string) {
		output := make([]string, len(inputs))
		copy(output, inputs)
		if gr.isValid(output) {
			c <- output
		}

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
			output := make([]string, len(inputs))
			copy(output, inputs)
			if gr.isValid(output) {
				c <- output
			}
			for i = 1; p[i] == 0; i++ {
				p[i] = i
			}
		}
	}

	c := make(chan []string)
	go func(c chan []string) {
		defer close(c)
		permutate(c, gr.towns)
	}(c)
	return c
}

// shortestPathPermutative returns the shortest path distance
// by checking through all permutations of the towns in the graph.
// This is an expensive operation.
// If gr has no valid paths it will return disconnected.
func (gr *graph) shortestPathPermutative() uint {
	allPermuts := gr.permutations()

	record := disconnected
	for path := range allPermuts {
		// guaranteed to be valid.
		if pathDist := gr.distanceSimple(path); pathDist < record {
			record = pathDist
		}
	}

	return record
}

// longestPathPermutative returns the longest path distance
// by checking through all permutations of the towns in the graph.
// This is an expensive operation.
// If gr has no valid paths it will return disconnected.
func (gr *graph) longestPathPermutative() uint {
	allPermuts := gr.permutations()

	record := uint(0)
	for path := range allPermuts {
		// guaranteed to be valid
		if pathDist := gr.distanceSimple(path); pathDist > record {
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
func (gr *graph) shortestPathGreedyFrom(town string) uint {
	// check first if town is in towns
	if !gr.in(town) {
		return disconnected
	}

	remaining := stringSliceBut(gr.towns, town)
	current := town
	var totalDistance uint

	for len(remaining) > 0 {
		// check distances from current to everything in remaining
		recordDistance, recordTown := disconnected, ""
		for _, eachDestination := range remaining {
			if eachDistance := gr.get(current, eachDestination); eachDistance < recordDistance {
				recordDistance, recordTown = eachDistance, eachDestination
			}
		}
		// now recordTown sounds like the winner.
		remaining = stringSliceBut(remaining, recordTown)
		current = recordTown
		totalDistance += recordDistance
	}

	return totalDistance
}

// shortestPathGreedy checks all towns in a graph's towns list to see which is the shortest to go to.
func (gr *graph) shortestPathGreedy() uint {
	record := disconnected
	for _, town := range gr.towns {
		if pathDistance := gr.shortestPathGreedyFrom(town); pathDistance < record {
			record = pathDistance
		}
	}
	return record
}

// Day09 solves the ninth day puzzle
// "All in a Single Night"
func Day09(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	santasGraph, err := newGraph(scanner)
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		answer1 = strconv.FormatUint(uint64(santasGraph.shortestPathPermutative()), 10)
		wg.Done()
	}()
	go func() {
		answer2 = strconv.FormatUint(uint64(santasGraph.longestPathPermutative()), 10)
		wg.Done()
	}()
	wg.Wait()
	return
}

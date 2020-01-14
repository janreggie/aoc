package aoc2015

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

// excludeTown returns the same slice of towns but without some specified value.
func excludeTown(from []town, except ...town) []town {
	out := make([]town, 0, len(from)/2) // don't worry it'll realloc.
	for _, vv := range from {
		// check if vv equals any of except
		isInExcept := false
		for _, exception := range except {
			if vv == exception {
				isInExcept = true
			}
		}
		if !isInExcept {
			out = append(out, vv)
		}
	}
	return out
}

// disconnected is a townDistance that is unfathomably large.
const disconnected townDistance = 4294967295

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
// The path can only be added to, and all towns must be unique.
// The path depends on a graph.
type townPath struct {
	raw       []town
	distance  townDistance
	graph     *townGraph
	remaining []town // the ones that could still be checked (from graph)
}

// townPathQueue is a queue of townPaths.
// The values in the queue are arranged by their distances,
// used to perform a breadth-first-search through the paths of a graph.
type townPathQueue []townPath

// uint64 converts townDistance to uint64
func (distance townDistance) uint64() uint64 {
	return uint64(distance)
}

// newGraph creates a graph construct
func newTownGraph(scanner *bufio.Scanner) (townGraph, error) {
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
func (graph *townGraph) set(a, b town, dist townDistance) {
	if a == b {
		return
	}
	if a > b {
		a, b = b, a
	}
	graph.links[townPair{a, b}] = dist
	writeA, writeB := true, true
	for _, v := range graph.towns {
		if v == a {
			writeA = false
		}
		if v == b {
			writeB = false
		}
	}
	if writeA {
		graph.towns = append(graph.towns, a)
	}
	if writeB {
		graph.towns = append(graph.towns, b)
	}
}

// get returns the distance between two towns.
// If a == b then return 0.
// If link doesn't exist then return disconnected.
func (graph townGraph) get(a, b town) townDistance {
	if a == b {
		return 0
	}
	if a > b {
		a, b = b, a
	}
	if dist, ok := graph.links[townPair{a, b}]; ok {
		return dist
	}
	return disconnected
}

// in returns true if town is in the graph's list of towns.
func (graph townGraph) in(town town) bool {
	for _, tt := range graph.towns {
		if tt == town {
			return true
		}
	}
	return false
}

// isValid returns whether the path is a isValid path in graph.
// If path is empty it will return true.
func (graph townGraph) isValid(path []town) bool {
	if len(path) == 0 {
		return true
	}
	if len(path) == 1 {
		return graph.in(path[0])
	}

	firstTwo := graph.get(path[0], path[1])
	if len(path) == 2 {
		return firstTwo != disconnected
	}

	return graph.isValid(path[1:])
}

// distance determines the distance returned after traversing all towns in path.
// Will also return a boolean signifying if path is a valid path.
// If path is not valid it will return disconnected and false.
// If path is empty it will return 0 and true.
func (graph townGraph) distance(path []town) (townDistance, bool) {
	if len(path) == 0 {
		return 0, true
	}
	if len(path) == 1 {
		return 0, graph.in(path[0])
	}

	firstTwo := graph.get(path[0], path[1])
	if len(path) == 2 {
		return firstTwo, firstTwo != disconnected
	}

	remaining, areValid := graph.distance(path[1:])
	if !areValid {
		return disconnected, false
	}
	return firstTwo + remaining, true
}

// distanceSimple traces the distance returned after traversing all towns in path.
// Does not check if the path is even valid.
// If it isn't then distanceSimple may return a number that is at least disconnectedLink.
func (graph townGraph) distanceSimple(path []town) townDistance {
	if len(path) <= 1 {
		return 0 // there is no distance to be travelled
	}
	if len(path) == 2 {
		return graph.get(path[0], path[1])
	}
	return graph.get(path[0], path[1]) + graph.distanceSimple(path[1:])
}

// permutations creates a channel that contains all the possible paths in the graph
// and will close said channel.
func (graph townGraph) permutations() <-chan []town {
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
		towns := make([]town, len(graph.towns))
		copy(towns, graph.towns)
		permutate(c, towns)
	}(c)
	return c
}

// shortestPathPermutative returns the shortest path distance
// by checking through all permutations of the towns in the graph.
// This is an expensive operation.
// If gr has no valid paths it will return disconnected.
func (graph townGraph) shortestPathPermutative() townDistance {
	allPermuts := graph.permutations()

	record := disconnected
	for path := range allPermuts {
		// guaranteed to be valid.
		if pathDist := graph.distanceSimple(path); pathDist < record {
			record = pathDist
		}
	}

	return record
}

// longestPathPermutative returns the longest path distance
// by checking through all permutations of the towns in the graph.
// This is an expensive operation.
// If gr has no valid paths it will return disconnected.
func (graph townGraph) longestPathPermutative() townDistance {
	allPermuts := graph.permutations()

	record := townDistance(0)
	for path := range allPermuts {
		// guaranteed to be valid
		if pathDist := graph.distanceSimple(path); pathDist > record {
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
// Rationale
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
func (graph townGraph) shortestPathGreedyFrom(from town) townDistance {
	// check first if town is in towns
	if !graph.in(from) {
		return disconnected
	}

	remaining := excludeTown(graph.towns, from)
	current := from
	var total townDistance

	for len(remaining) > 0 {
		// check distances from current to everything in remaining
		recordDist, recordTown := disconnected, town("")
		for _, eachDestination := range remaining {
			if eachDistance := graph.get(current, eachDestination); eachDistance < recordDist {
				recordDist, recordTown = eachDistance, eachDestination
			}
		}
		// now recordTown sounds like the winner.
		remaining = excludeTown(remaining, recordTown)
		current = recordTown
		total += recordDist
	}

	return total
}

// shortestPathGreedy checks all towns in a graph's towns list to see which is the shortest to go to.
func (graph townGraph) shortestPathGreedy() townDistance {
	record := disconnected
	for _, town := range graph.towns {
		if pathDistance := graph.shortestPathGreedyFrom(town); pathDistance < record {
			record = pathDistance
		}
	}
	return record
}

// longestPathGreedyFrom is like shortestPathGreedyFrom but returns
// the longest path it could trace from town using a greedy algorithm
func (graph townGraph) longestPathGreedyFrom(from town) townDistance {
	// check first if town is in towns
	if !graph.in(from) {
		return disconnected
	}

	remaining := excludeTown(graph.towns, from)
	current := from
	var total townDistance

	for len(remaining) > 0 {
		// check distances from current to everything in remaining
		recordDist, recordTown := townDistance(0), town("")
		for _, eachDestination := range remaining {
			if eachDistance := graph.get(current, eachDestination); eachDistance > recordDist {
				recordDist, recordTown = eachDistance, eachDestination
			}
		}
		// now recordTown sounds like the winner.
		remaining = excludeTown(remaining, recordTown)
		current = recordTown
		total += recordDist
	}

	return total
}

// longestPathGreedy checks all towns in a graph's towns list to see which is the shortest to go to.
func (graph townGraph) longestPathGreedy() townDistance {
	var record townDistance
	for _, town := range graph.towns {
		if pathDistance := graph.longestPathGreedyFrom(town); pathDistance > record {
			record = pathDistance
		}
	}
	return record
}

// shortestPathCleverFrom determines the shortest path from a town
// using an algorithm a bit smarter than the previous one,
// although this algorithm is still exhaustive.
// This uses the concept of townPaths and a townPathQueue to record all paths.
// If there is no valid path from town, return disconnected.
func (graph townGraph) shortestPathCleverFrom(town town) townDistance {
	// we should be able to solve this...
	allPaths := newTownPathQueue()

	// let's use tg.towns...
	for _, tt := range graph.towns {
		path, err := newTownPath(town, tt, &graph)
		if err != nil { // maybe same town? or no direct path?
			continue
		}
		allPaths.push(path)
	}

	for len(allPaths) > 0 {
		path, _ := allPaths.pop() // guaranteed no error
		if len(path.remaining) == 0 {
			// i guess that's our winner
			return path.distance
		}

		// otherwise let's consider all remainings
		for _, next := range path.remaining {
			nextPath, err := path.add(next)
			if err != nil {
				continue // there should be no error...
			}
			allPaths.push(nextPath)
		}
	}

	return disconnected // well we tried...
}

// longestPathNotSoCleverFrom computes the longest path from a town.
func (graph townGraph) longestPathNotSoCleverFrom(town town) townDistance {
	allPaths := newTownPathQueue()

	// let's use tg.towns...
	for _, tt := range graph.towns {
		path, err := newTownPath(town, tt, &graph)
		if err != nil { // maybe same town? or no direct path?
			continue
		}
		allPaths.push(path)
	}

	var record townDistance

	for len(allPaths) > 0 {
		path, _ := allPaths.pop() // guaranteed no error
		if len(path.remaining) == 0 {
			// this is part of what makes this algo take so long.
			// we're exhausting all potential paths
			// from shortest to longest.
			if path.distance > record {
				record = path.distance
			}
		}

		// otherwise let's consider all remainings
		for _, next := range path.remaining {
			nextPath, err := path.add(next)
			if err != nil {
				continue // there should be no error...
			}
			allPaths.push(nextPath)
		}
	}

	if record == 0 {
		return disconnected // there are no paths...
	}
	return record
}

// shortestPathClever spins up shortestPathCleverFrom for all towns in a graph
func (graph townGraph) shortestPathClever() townDistance {
	record := disconnected
	for _, town := range graph.towns {
		if pathDistance := graph.shortestPathCleverFrom(town); pathDistance < record {
			record = pathDistance
		}
	}
	return record

}

// longestPathNotSoClever spins up longestPathCleverFrom for all towns in a graph
func (graph townGraph) longestPathNotSoClever() townDistance {
	var record townDistance
	for _, town := range graph.towns {
		if pathDistance := graph.longestPathNotSoCleverFrom(town); pathDistance > record {
			record = pathDistance
		}
	}
	return record

}

// newTownPath creates a new townPath construct that starts with two towns.
// Will return an error if a and b are not towns in townGraph.
// Will also return an error if a and b are equal.
func newTownPath(a, b town, graph *townGraph) (townPath, error) {
	if a == b {
		return townPath{}, fmt.Errorf("could not create path to itself")
	}
	dist, isValid := graph.distance([]town{a, b})
	if !isValid {
		return townPath{}, fmt.Errorf("could not create path between %v and %v", a, b)
	}
	return townPath{raw: []town{a, b}, distance: dist, graph: graph, remaining: excludeTown(graph.towns, a, b)}, nil
}

// copy copies the townPath construct
func (path townPath) copy() townPath {
	output := townPath{}
	output.raw = make([]town, len(path.raw))
	copy(output.raw, path.raw)
	output.distance = path.distance
	output.graph = path.graph
	output.remaining = make([]town, len(path.remaining))
	copy(output.remaining, path.remaining)
	return output
}

// head and tail returns the first and last towns of a townPath
func (path *townPath) head() town {
	return path.raw[0]
}

func (path *townPath) tail() town {
	return path.raw[len(path.raw)-1]
}

// add adds a town to a town path and returns said path.
// Will return an error if a path between the tail and the next town could not be found.
// Will also return an error if next is already in the path.
func (path townPath) add(next town) (townPath, error) {
	// check if tp.remaining is empty
	if len(path.remaining) == 0 {
		return townPath{}, fmt.Errorf("but there are no potential towns left")
	}

	// check if next is in raw
	for _, town := range path.raw {
		if town == next {
			return townPath{}, fmt.Errorf("%v already in path", next)
		}
	}

	// check distance bet. last() and next
	tail := path.tail()
	nextDistance := path.graph.get(tail, next)
	if nextDistance == disconnected {
		return townPath{}, fmt.Errorf("no link between %v and %v", tail, next)
	}

	newTp := path.copy()
	newTp.raw = append(newTp.raw, next)
	newTp.distance += nextDistance
	newTp.remaining = excludeTown(newTp.remaining, next)
	return newTp, nil
}

// newTownPathQueue creates a townPathQueue construct.
func newTownPathQueue() townPathQueue {
	return make(townPathQueue, 0) // it's that simple
}

// push pushes a townPath to the queue.
func (queue *townPathQueue) push(path townPath) {
	if len(*queue) == 0 {
		*queue = append(*queue, path.copy())
		return
	}

	ind := 0
	for ind < len(*queue) && (*queue)[ind].distance < path.distance {
		ind++
	}
	// append a value to tpq. Just any value..
	*queue = append(*queue, townPath{})
	// move all elements...
	for ii := len(*queue) - 1; ii > ind; ii-- {
		(*queue)[ii] = (*queue)[ii-1]
	}
	(*queue)[ind] = path
}

// pop pops a value from the queue.
// Will return an error if the queue is empty.
func (queue *townPathQueue) pop() (townPath, error) {
	if len(*queue) == 0 {
		return townPath{}, errors.New("queue is empty")
	}
	out := (*queue)[0]
	*queue = (*queue)[1:]
	return out, nil
}

// Day09 solves the ninth day puzzle
// "All in a Single Night"
func Day09(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	santasGraph, err := newTownGraph(scanner)
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		answer1 = strconv.FormatUint(santasGraph.shortestPathClever().uint64(), 10)
		wg.Done()
	}()
	go func() {
		// use permutative. clever's not really clever.
		answer2 = strconv.FormatUint(santasGraph.longestPathPermutative().uint64(), 10)
		wg.Done()
	}()
	wg.Wait()

	return
}

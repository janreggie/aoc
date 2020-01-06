package aoc2015

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// graph represents a graph
type graph struct {
	links map[string]uint // pls no negative distance
	towns []string
}

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
// If link doesn't exist then return MaxUint.
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
	return uint(math.MaxUint32)
}

// distance traces the distance returned after traversing all towns in []string.
// Does not check if towns are in graph.
func (gr graph) distance(towns []string) uint {
	if len(towns) <= 1 {
		return 0 // there is no distance to be travelled
	}
	if len(towns) == 2 {
		return gr.get(towns[0], towns[1])
	}
	return gr.get(towns[0], towns[1]) + gr.distance(towns[1:])
}

// generatePermutations generates a channel containing permutations of a string slice
// and closes once all permutations have been exhausted
func generatePermutations(data []string) <-chan []string {
	c := make(chan []string)
	go func(c chan []string) {
		defer close(c)
		permutate(c, data)
	}(c)
	return c
}

func permutate(c chan []string, inputs []string) {
	output := make([]string, len(inputs))
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
		output := make([]string, len(inputs))
		copy(output, inputs)
		c <- output
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

// Day09 solves the ninth day puzzle
// "All in a Single Night"
func Day09(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	santasGraph, err := newGraph(scanner)
	if err != nil {
		return
	}

	allPermuts := generatePermutations(santasGraph.towns)

	highest := uint(0)
	lowest := uint(math.MaxUint32)
	for path := range allPermuts {
		if pathDist := santasGraph.distance(path); pathDist > highest {
			highest = pathDist
		} else if pathDist < lowest {
			lowest = pathDist
		}
	}

	answer1 = strconv.FormatUint(uint64(lowest), 10)
	answer2 = strconv.FormatUint(uint64(highest), 10)
	return
}

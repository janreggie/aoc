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
type graph map[string]map[string]uint // pls no negative distance

func newGraph(scanner *bufio.Scanner) (graph, error) {
	out := make(graph)

	// now fill it in for each line
	for scanner.Scan() {
		// assume format is
		// `AlphaNumericStringA to AlphaNumericStringB = UnsignedInteger`
		text := strings.Fields(scanner.Text())
		if len(text) != 5 {
			return nil, fmt.Errorf("could not parse %v", scanner.Text())
		}
		townOne, townTwo := text[0], text[2]
		distance, err := strconv.Atoi(text[4])
		if err != nil {
			return nil, errors.Wrapf(err, "could not convert %v in %v", text[4], scanner.Text())
		}
		if distance < 0 {
			return nil, errors.Wrapf(err, "could not fathom distance %v in %v", distance, scanner.Text())
		}
		out.add(townOne, townTwo, uint(distance))
	}

	return out, nil
}

// add adds the edge that connects a and b with distance dist
func (gr graph) add(a, b string, dist uint) {
	if _, ok := gr[a]; !ok {
		gr[a] = make(map[string]uint)
	}
	if _, ok := gr[b]; !ok {
		gr[b] = make(map[string]uint)
	}

	gr[a][b] = dist
	gr[b][a] = dist
	gr[a][a] = 0 // self distance!
	gr[b][b] = 0 // self distance!
}

// distance traces the distance returned after traversing all towns in []string.
// Does not check if towns are in graph.
func (gr graph) distance(towns []string) uint {
	if len(towns) <= 1 {
		return 0 // there is no distance to be travelled
	}
	if len(towns) == 2 {
		return gr[towns[0]][towns[1]]
	}
	return gr[towns[0]][towns[1]] + gr.distance(towns[1:])
}

// findMinimumHamiltonian returns the smallest Hamiltonian distance there is for a graph
// by exhausting all permutations.
func (gr graph) findMinimumHamiltonian() uint {
	// what are all the Towns again?
	// permute from https://www.golangprograms.com/golang-program-to-generate-slice-permutations-of-number-entered-by-user.html
	permuteStrings := func(xs []string) (permuts [][]string) {
		var rc func([]string, int)
		rc = func(a []string, k int) {
			if k == len(a) {
				permuts = append(permuts, append([]string{}, a...))
			} else {
				for i := k; i < len(xs); i++ {
					a[k], a[i] = a[i], a[k]
					rc(a, k+1)
					a[k], a[i] = a[i], a[k]
				}
			}
		}
		rc(xs, 0)
		return permuts
	}

	allTowns := []string{}
	for town := range gr {
		allTowns = append(allTowns, town)
	}
	allPermuts := permuteStrings(allTowns)

	// check lowest
	lowest := uint(math.MaxUint32)
	for _, path := range allPermuts {
		if pathDist := gr.distance(path); pathDist < lowest {
			lowest = pathDist
		}
	}
	return lowest
}

// findMaximumHamiltonian returns the largest Hamiltonian distance there is for a graph
// by exhausting all permutations.
func (gr graph) findMaximumHamiltonian() uint {
	// what are all the Towns again?
	// permute from https://www.golangprograms.com/golang-program-to-generate-slice-permutations-of-number-entered-by-user.html
	permuteStrings := func(xs []string) (permuts [][]string) {
		var rc func([]string, int)
		rc = func(a []string, k int) {
			if k == len(a) {
				permuts = append(permuts, append([]string{}, a...))
			} else {
				for i := k; i < len(xs); i++ {
					a[k], a[i] = a[i], a[k]
					rc(a, k+1)
					a[k], a[i] = a[i], a[k]
				}
			}
		}
		rc(xs, 0)
		return permuts
	}

	allTowns := []string{}
	for town := range gr {
		allTowns = append(allTowns, town)
	}
	allPermuts := permuteStrings(allTowns)

	// check highest
	highest := uint(0)
	for _, path := range allPermuts {
		if pathDist := gr.distance(path); pathDist > highest {
			highest = pathDist
		}
	}
	return highest
}

// Day09 solves the ninth day puzzle
// "All in a Single Night"
func Day09(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	santasGraph, err := newGraph(scanner)
	if err != nil {
		return
	}
	answer1 = strconv.FormatUint(uint64(santasGraph.findMinimumHamiltonian()), 10)
	answer2 = strconv.FormatUint(uint64(santasGraph.findMaximumHamiltonian()), 10)
	return
}

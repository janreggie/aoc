package aoc2019

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type orbit struct {
	object       string   // the name of the object itself
	parent       *orbit   // the one that it orbits
	satellites   []string // those that orbit around the object
	orbiterCount int      // how many satellites, direct or indirect, does object have?
}

// newOrbit creates an orbit struct and adds to some census (map[string]*orbit)
func newOrbit(input string, census map[string]*orbit) (err error) {
	// input could be, for instance, "A)B"
	// check first if "A" already exists in census
	raw := strings.Split(input, ")")
	if len(raw) != 2 {
		// wrong input!
		err = fmt.Errorf("invalid input %v", input)
		return
	}

	// check raw[0]'s existence
	if center, ok := census[raw[0]]; !ok {
		// let's create a new object
		census[raw[0]] = &orbit{
			object:       raw[0],
			satellites:   []string{raw[1]},
			orbiterCount: 0, // we'll deal with this later
		}
	} else {
		// if it does exist!
		center.satellites = append(center.satellites, raw[1])
	}

	// check raw[1]'s existence
	if satellite, ok := census[raw[1]]; !ok {
		census[raw[1]] = &orbit{
			object:       raw[1],
			parent:       census[raw[0]], // it should exist by now!
			satellites:   []string{},
			orbiterCount: 0,
		}
	} else {
		// what happens if it does?
		satellite.parent = census[raw[0]]
	}

	// now the satellite's parent has one more satellite
	// as well as that parent's parent
	// and that one's parent
	// ...and so on.
	for ptr := census[raw[0]]; ptr != nil; ptr = ptr.parent {
		// ptr.orbiterCount++ // increments
		ptr.orbiterCount = 0
		for _, child := range ptr.satellites {
			ptr.orbiterCount += census[child].orbiterCount + 1
		}
	}
	return nil
}

// orbitalTransfers computes for how many orbital transfers two nodes require
func orbitalTransfers(ob1, ob2 *orbit) int {
	var lca *orbit // lowest common ancestor
	ob1parents := make([]*orbit, 0)
	ob1gen := 0 // how many steps from lca to ob1
	for ptr := ob1; ptr != nil; ptr = ptr.parent {
		ob1parents = append(ob1parents, ptr)
	}
	ob2parents := make([]*orbit, 0)
exterloop:
	for ptr := ob2; ptr != nil; ptr = ptr.parent {
		for count, key := range ob1parents {
			if key == ptr {
				lca = ptr // well ptr is the lca!
				ob1gen = count
				break exterloop
			}
		}
		ob2parents = append(ob2parents, ptr)
	}
	_ = lca // we didn't really use this did we?
	return ob1gen + len(ob2parents) - 2
}

// Day06 solves the sixth day puzzle "Universal Orbit Map".
//
// Input
//
// A file containing several lines each representing an orbit,
// where each orbit is of the form IDENT)IDENT,
// where IDENT is an alphanumeric string of at most length 3.
// For example:
//
//	4PT)ZHN
//	MHG)8YQ
//	J9M)B25
//	KGG)845
//
// It is assumed that there are no more than 1400 combinations of orbits,
// and that there is always an IDENT that is "YOU" and "SAN" in the input.
func Day06(input string) (answer1, answer2 string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	allSatellites := make(map[string]*orbit, 0)
	for scanner.Scan() {
		raw := strings.Fields(scanner.Text())
		if len(raw) != 1 {
			continue
		}
		if err = newOrbit(raw[0], allSatellites); err != nil {
			return
		}
	}

	// orbiterCount is number of orbits there are with object as center
	// total number of orbits == sum of orbiterCount
	totalOrbits := 0
	for _, val := range allSatellites {
		totalOrbits += val.orbiterCount
	}
	answer1 = strconv.Itoa(totalOrbits)
	answer2 = strconv.Itoa(orbitalTransfers(allSatellites["YOU"], allSatellites["SAN"]))

	return
}

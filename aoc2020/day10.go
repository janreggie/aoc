package aoc2020

import (
	"bufio"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// countAdapterArrangement counts the number of arrangements for a given slice of jolt output voltages.
// An arrangement is such if the difference between all its numbers is 1 or 3.
func countAdapterArrangement(arr []int) int {
	// cache is a cache to remember what we have since countAdapterArrangement branches like hell
	cache := make([]struct {
		val     int
		written bool
	}, len(arr))
	// readCache reads the ind'th entry from the cache. Returns -1 if value has yet to be written.
	readCache := func(ind int) int {
		cur := cache[ind]
		if cur.written {
			return cur.val
		}
		return -1
	}
	// writeCache writes a value to cache
	writeCache := func(ind, val int) {
		cache[ind].val = val
		cache[ind].written = true
	}

	// recurse is the one that's actually in charge of counting the arrangements.
	// Takes in as well an index to remember which in the parent array are we.
	var recurse func([]int, int) int
	recurse = func(aa []int, ind int) int {
		if len(aa) <= 1 {
			return 1 // it is an arrangement in its own right
		}
		if rr := readCache(ind); rr != -1 {
			return rr
		}

		direct, skipOne, skipTwo := 0, 0, 0
		if len(aa) >= 2 && aa[1]-aa[0] <= 3 {
			direct = recurse(aa[1:], ind+1)
		}
		if len(aa) >= 3 && aa[2]-aa[0] <= 3 {
			skipOne = recurse(aa[2:], ind+2)

		}
		if len(aa) >= 4 && aa[3]-aa[0] <= 3 {
			skipTwo = recurse(aa[3:], ind+3)
		}
		result := direct + skipOne + skipTwo
		writeCache(ind, result)
		return result
	}
	return recurse(arr, 0)
}

// Day10 solves the tenth day puzzle "Adapter Array"
//
// Input
//
// A list of jolt ratings, separated by newlines. For example:
//
//   16
//   10
//   15
//   5
//   1
//   11
//   7
//   19
//   6
//   12
//   4
//
// It is guaranteed that all entries are nonnegative integers no more than 200.
func Day10(input string) (answer1, answer2 string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	// adapters is the slice of adapters (puzzle input)
	adapters := make([]int, 0)
	for scanner.Scan() {
		aa, e := strconv.Atoi(scanner.Text())
		if e != nil {
			err = errors.Wrapf(e, "couldn't parse %s from input", scanner.Text())
			return
		}
		adapters = append(adapters, aa)
	}

	adaptersSorted := make([]int, len(adapters))
	copy(adaptersSorted, adapters)
	sort.Ints(adaptersSorted)

	oneVoltDiffs, threeVoltDiffs := 0, 0
	prev := 0 // the zeroth one
	for _, curr := range adaptersSorted {
		if curr-prev == 1 {
			oneVoltDiffs++
		} else if curr-prev == 3 {
			threeVoltDiffs++
		}
		prev = curr
	}
	deviceBuiltin := prev + 3
	threeVoltDiffs++ // your device's built-in adapter
	answer1 = strconv.Itoa(oneVoltDiffs * threeVoltDiffs)

	// Now for the second part...
	masterChain := make([]int, len(adaptersSorted)+2)
	masterChain[0] = 0
	copy(masterChain[1:], adaptersSorted)
	masterChain[len(masterChain)-1] = deviceBuiltin
	answer2 = strconv.Itoa(countAdapterArrangement(masterChain))
	return
}

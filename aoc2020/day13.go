package aoc2020

import (
	"sort"
	"strconv"
	"strings"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/pkg/errors"
)

// Day13 solves the thirteenth day puzzle "Shuttle Search"
//
// Input
//
// The earliest time you could depart from a bus, and then a list of bus IDs with some `x`s.
// For example:
//
//   939
//   7,13,x,x,59,x,31,19
//
// The actions will either be N, S, E, W, L, R, and F;
// and the values are guaranteed to be positive numbers no more than 300.
// The values for L and R will be multiples of 90 (90, 180, 270).
func Day13(input string) (answer1, answer2 string, err error) {

	lines := aoc.SplitLines(input)
	if len(lines) != 2 {
		err = errors.Errorf("expected number of lines to be 2, got %v instead", len(lines))
		return
	}

	timestamp, err := strconv.Atoi(lines[0])
	if err != nil {
		err = errors.Wrapf(err, "could not read earliest timestamp %s", lines[0])
		return
	}

	rawBuses := strings.Split(lines[1], ",")
	buses := make([]int, len(rawBuses))
	for ii := range rawBuses {
		var e error
		if buses[ii], e = strconv.Atoi(rawBuses[ii]); e != nil { // x
			buses[ii] = 0
		}
	}

	// max returns the maximum of a list of integers
	max := func(ints []int) int {
		cand := ints[0]
		for _, vv := range ints {
			if vv > cand {
				cand = vv
			}
		}
		return cand
	}

	// Now for some maths magic
	// recordDeparture is the earliest possible time one can depart.
	recordDeparture, recordBusID := timestamp+max(buses), 0
	for _, bus := range buses {
		if bus == 0 {
			continue
		}
		current := timestamp + bus - (timestamp % bus)
		if current < recordDeparture {
			recordDeparture = current
			recordBusID = bus
		}
	}
	answer1 = strconv.Itoa(recordBusID * (recordDeparture - timestamp))

	// Now for the second one.
	// Note: Chinese remainder theorem

	// Check if buses are coprime
	// greatest common divisor (GCD) via Euclidean algorithm
	gcd := func(a, b int) int {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		return a
	}

	// find Least Common Multiple (LCM) via GCD
	var lcm func(a, b int, integers ...int) int
	lcm = func(a, b int, integers ...int) int {
		result := a * b / gcd(a, b)

		for i := 0; i < len(integers); i++ {
			result = lcm(result, integers[i])
		}

		return result
	}
	product := func(integers ...int) int {
		res := 1
		for _, vv := range integers {
			res *= vv
		}
		return res
	}

	busesWithoutZero := make([]int, 0)
	for _, bus := range buses {
		if bus == 0 {
			continue
		}
		busesWithoutZero = append(busesWithoutZero, bus)
	}
	if product(busesWithoutZero...) != lcm(busesWithoutZero[0], busesWithoutZero[1], busesWithoutZero[2:]...) {
		err = errors.Errorf("input bus IDs %v are not coprime", busesWithoutZero)
		return
	}

	busesAndDelays := make([][2]int64, 0) // 0: bus ID, 1: delay
	for ii, bus := range buses {
		if bus == 0 {
			continue
		}
		busesAndDelays = append(busesAndDelays, [2]int64{int64(bus), int64(ii)}) // don't do ii % bus just yet...
	}
	// Now do some sieving. Would be faster if busesAndDelays are sorted by bb[0] descending
	sort.Slice(busesAndDelays, func(ii, jj int) bool { return busesAndDelays[ii][0] > busesAndDelays[jj][0] })

	var step int64 = 1
	var guess int64 = 1
	for _, bd := range busesAndDelays {
		for (guess+bd[1])%bd[0] != 0%bd[0] {
			guess += step
		}
		step *= bd[0]
	}

	answer2 = strconv.FormatInt(guess, 10)

	return
}

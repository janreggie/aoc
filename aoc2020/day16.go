package aoc2020

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// fieldRange represents a field and its (two!) ranges
type fieldRange struct {
	name string // what is the name of the field?
	r1   [2]int // first range
	r2   [2]int // second range
}

func (rng fieldRange) String() string {
	return fmt.Sprintf("%s: %d-%d or %d-%d", rng.name, rng.r1[0], rng.r1[1], rng.r2[0], rng.r2[1])
}

// newFieldRange generates a fieldRange object from a string
func newFieldRange(fr string) (fieldRange, error) {
	frSplit := strings.Split(fr, ": ")
	if len(frSplit) != 2 {
		return fieldRange{}, errors.Errorf("could not parse %s: lacking colon and space", fr)
	}
	var result fieldRange
	result.name = frSplit[0]
	if _, err := fmt.Sscanf(frSplit[1], "%d-%d or %d-%d", &result.r1[0], &result.r1[1], &result.r2[0], &result.r2[1]); err != nil {
		return fieldRange{}, errors.Wrapf(err, "could not parse bounds %s well", frSplit[0])
	}
	return result, nil
}

// eq returns true if its argument is equal to itself
func (rng fieldRange) eq(compare fieldRange) bool {
	return rng.name == compare.name &&
		rng.r1 == compare.r1 &&
		rng.r2 == compare.r2
}

// in returns if a number is in fieldRange
func (rng fieldRange) in(num int) bool {
	if rng.r1[0] <= num && num <= rng.r1[1] {
		return true
	}
	if rng.r2[0] <= num && num <= rng.r2[1] {
		return true
	}
	return false
}

// deducePossible reduces a 2D slice of fieldRanges to a slice of possible solutions.
// For instance, if rngs={{A,B},{C},{A,B,C}},
// its output should be {{A,C,B},{B,C,A}}.
// For efficiency sake, rngs ought to be sorted ascending by length.
func deducePossible(rngs [][]fieldRange) [][]fieldRange {
	if len(rngs) == 0 {
		return [][]fieldRange{{}}
	}
	// if len(rngs) == 1 {
	// output := make([][]fieldRange, len(rngs[0]))
	// for ii := range output {
	// output[ii] = []fieldRange{rngs[0][ii]}
	// }
	// return output
	// }

	// now for the big guns...
	result := make([][]fieldRange, 0)

FOREVERYFIRST:
	for _, cur := range rngs[0] {
		// fmt.Println("cur is", cur, "and length of rngs is", len(rngs))
		// rest is rngs[1:] without cur
		rest := make([][]fieldRange, len(rngs)-1)
		for jj := range rest {
			rest[jj] = make([]fieldRange, 0)
			for _, rr := range rngs[jj+1] {
				if !cur.eq(rr) {
					rest[jj] = append(rest[jj], rr)
				}
			}

			// If there is some of rest that's empty, there's no point in continuing
			if len(rest[jj]) == 0 {
				continue FOREVERYFIRST
			}
		}

		for _, res := range deducePossible(rest) {
			curRes := make([]fieldRange, len(rngs))
			curRes[0] = cur
			copy(curRes[1:], res)
			result = append(result, curRes)
		}
	}
	return result
}

// ticketNotes represents the notes you have taken for your high-speed train trip
type ticketNotes struct {
	fields     []fieldRange
	yourTicket []int
	allTickets [][]int
}

// generateTicketNotes generates a ticketNotes object using a scanner
func generateTicketNotes(scanner *bufio.Scanner) (ticketNotes, error) {
	result := ticketNotes{}
	result.fields = make([]fieldRange, 0)
	result.allTickets = make([][]int, 0)

	state := 0 // 0 for scanning fields; 1 for scanning your ticket; 2 for scanning nearby tickets
	for scanner.Scan() {
		if scanner.Text() == "" {
			state++
			continue
		}

		switch state {
		case 0:
			ff, err := newFieldRange(scanner.Text())
			if err != nil {
				return result, errors.Wrapf(err, "could not read %s from scanner", scanner.Text())
			}
			result.fields = append(result.fields, ff)

		case 1:
			if scanner.Text() == "your ticket:" {
				continue
			}
			splitVals := strings.Split(scanner.Text(), ",")
			if len(splitVals) != len(result.fields) {
				return result, errors.Errorf("expected number of values for %v to be %d, got %d instead", scanner.Text(), len(result.fields), len(splitVals))
			}
			result.yourTicket = make([]int, len(splitVals))
			for ii, vv := range splitVals {
				val, err := strconv.Atoi(vv)
				if err != nil {
					return result, errors.Errorf("cannot convert %v from %v to int", vv, scanner.Text())
				}
				result.yourTicket[ii] = val
			}

		case 2:
			if scanner.Text() == "nearby tickets:" {
				continue
			}
			splitVals := strings.Split(scanner.Text(), ",")
			if len(splitVals) != len(result.fields) {
				return result, errors.Errorf("expected number of values for %v to be %d, got %d instead", scanner.Text(), len(result.fields), len(splitVals))
			}
			currentTicket := make([]int, len(splitVals))
			for ii, vv := range splitVals {
				val, err := strconv.Atoi(vv)
				if err != nil {
					return result, errors.Errorf("cannot convert %v from %v to int", vv, scanner.Text())
				}
				currentTicket[ii] = val
			}
			result.allTickets = append(result.allTickets, currentTicket)

		default:
			return result, errors.Errorf("scanner might be invalid: too many blank lines")
		}
	}
	return result, nil
}

// invalidValuesAt returns the sum of invalid values of ticket at some index.
// Note that it doesn't care about which value corresponds to which field.
// If ind is out of bounds, return 0.
func (notes ticketNotes) invalidValuesAt(ind int) int {
	if ind < 0 || ind >= len(notes.allTickets) {
		return 0
	}
	result := 0
	ticket := notes.allTickets[ind]
TICKETLOOP:
	for _, tv := range ticket {
		for _, field := range notes.fields {
			if field.in(tv) {
				continue TICKETLOOP
			}
		}
		result += tv
	}
	return result
}

// discardInvalids removes tickets from allTickets that contain invalid value/s.
func (notes *ticketNotes) discardInvalids() {
	newTickets := make([][]int, 0)
	for ii := range notes.allTickets {
		if vv := notes.invalidValuesAt(ii); vv == 0 {
			newTickets = append(newTickets, notes.allTickets[ii])
		}
	}
	notes.allTickets = newTickets
}

// findFieldRanges returns an array of array of fieldRanges
// whose values satisfy each index of all the tickets at notes.
func (notes ticketNotes) findFieldRanges() [][]fieldRange {
	result := make([][]fieldRange, len(notes.yourTicket))
	for ii := range result {
		result[ii] = make([]fieldRange, len(notes.fields))
		copy(result[ii], notes.fields)
	}

	// result[ii] should have all fields, but this will be truncated one by one as we progress with notes.allNotes
	for ii := range result {
		for _, tck := range notes.allTickets {
			// tck[ii] will be checked against all values of result[ii]
			toCheck := tck[ii]
			improperIndices := make([]int, 0) // indices from result[ii] that don't satisfy toCheck
			for jj, rng := range result[ii] {
				if !rng.in(toCheck) {
					improperIndices = append(improperIndices, jj)
				}
			}

			if len(improperIndices) == 0 {
				continue
			}
			if len(improperIndices) == 1 {
				toPop := improperIndices[0]
				origLen := len(result[ii])
				result[ii][toPop], result[ii][origLen-1] = result[ii][origLen-1], result[ii][toPop]
				result[ii] = result[ii][:origLen-1]
				continue
			}

			// recreate result[ii] but without elements from improperIndices
			// inImproper returns if x is in improperIndices
			inImproper := func(x int) bool {
				for _, vv := range improperIndices {
					if x == vv {
						return true
					}
				}
				return false
			}
			newResult := make([]fieldRange, 0)
			for jj := range result[ii] {
				if inImproper(jj) {
					continue
				}
				newResult = append(newResult, result[ii][jj])
			}
			result[ii] = newResult
		}
	}
	return result
}

// Day16 solves the sixteenth day puzzle "Ticket Translation"
//
// Input
//
// A text file containing the notes you have taken down. For example:
//
//   departure location: 43-237 or 251-961
//   departure station: 27-579 or 586-953
//   departure platform: 31-587 or 608-967
//   departure track: 26-773 or 784-973
//   departure date: 41-532 or 552-956
//   departure time: 33-322 or 333-972
//   arrival location: 30-165 or 178-965
//   arrival station: 31-565 or 571-968
//   arrival platform: 36-453 or 473-963
//   arrival track: 35-912 or 924-951
//   class: 39-376 or 396-968
//   duration: 31-686 or 697-974
//   price: 28-78 or 96-971
//   route: 32-929 or 943-955
//   row: 40-885 or 896-968
//   seat: 26-744 or 765-967
//   train: 46-721 or 741-969
//   type: 30-626 or 641-965
//   wagon: 48-488 or 513-971
//   zone: 34-354 or 361-973
//
//   your ticket:
//   151,71,67,113,127,163,131,59,137,103,73,139,107,101,97,149,157,53,109,61
//
//   nearby tickets:
//   680,418,202,55,792,800,896,801,312,252,721,702,24,112,608,837,98,222,797,364
//   876,910,289,816,873,280,791,313,641,15,719,587,353,366,617,710,565,419,339,621
//   869,683,645,185,121,51,670,401,308,213,54,150,813,264,330,50,444,825,837,368
//   244,131,561,851,165,270,78,573,818,869,946,370,527,866,655,816,146,792,431,431
//   626,744,334,482,185,797,770,723,532,475,441,62,837,657,412,106,404,433,577,439
//   610,514,611,681,156,854,838,126,524,710,518,193,744,857,258,642,12,273,192,221
//   215,408,404,344,319,720,403,434,885,248,707,195,672,312,809,855,118,529,673,439
//   103,402,226,678,711,212,843,806,358,436,342,257,98,401,527,831,524,401,275,709
//   429,906,336,77,146,122,179,295,251,529,487,858,901,71,417,477,458,474,154,786
//   903,313,919,564,904,853,673,373,209,408,825,516,441,818,474,773,556,848,151,229
//   96,861,791,620,929,256,587,211,250,684,420,258,365,96,104,653,368,528,784,817
//   898,866,669,867,901,618,575,210,498,265,413,556,190,376,624,818,295,294,431,649
//   50,346,807,717,850,820,791,561,486,752,414,669,927,72,205,117,137,564,744,799
//
// It is guaranteed that the input resembles the one above,
// but it is possible for one's input to use a different number of fields.
func Day16(input string) (answer1, answer2 string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	notes, err := generateTicketNotes(scanner)
	if err != nil {
		err = errors.Wrapf(err, "could not read from input scanner")
		return
	}
	scanningErrorRate := 0
	for ii := range notes.allTickets {
		scanningErrorRate += notes.invalidValuesAt(ii)
	}
	answer1 = strconv.Itoa(scanningErrorRate)

	// Now for the second part
	notes.discardInvalids()
	validRanges := notes.findFieldRanges()

	// validRanges ought to be sorted by length to make deducePossible as fast as possible.
	// However, the arrangement of validRanges is important in evaluating yourTicket.
	// Let us create a data structure that takes in validRanges...
	orderedValidRanges := make([]struct {
		order int
		rngs  []fieldRange
	}, len(validRanges))
	for ii := range validRanges {
		orderedValidRanges[ii].order = ii
		orderedValidRanges[ii].rngs = validRanges[ii]
	}
	sort.Slice(orderedValidRanges, func(i, j int) bool {
		return len(orderedValidRanges[i].rngs) < len(orderedValidRanges[j].rngs)
	})

	// And let us reassign validRanges, but take note of the orders that we have placed on it
	for ii := range validRanges {
		validRanges[ii] = orderedValidRanges[ii].rngs
	}

	deduced := deducePossible(validRanges)
	if len(deduced) == 0 {
		err = errors.Errorf("invalid input: could not find a possible configuration for fields")
		return
	} else if len(deduced) > 1 {
		err = errors.Errorf("invalid input: could not find an inambiguous configuration for fields, got %v", deduced)
	}

	actualConfig := deduced[0]

	departureProduct := 1
	for ii, rng := range actualConfig {
		if strings.HasPrefix(rng.name, "departure") {
			departureProduct *= notes.yourTicket[orderedValidRanges[ii].order]
		}
	}

	answer2 = strconv.Itoa(departureProduct)

	return
}

package aoc2020

import (
	"sort"
	"strconv"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/pkg/errors"
)

// planeSeat represents a plane seat
type planeSeat struct {
	row int // 0 to 127
	col int // 0 to 7
}

// readBoardingPass reads a boarding pass string and returns a planeSeat
func readBoardingPass(boardingPass string) (planeSeat, error) {
	if len(boardingPass) != 10 {
		return planeSeat{}, errors.Errorf("%s is not length 10", boardingPass)
	}

	seat := planeSeat{}
	for ii := 0; ii < 7; ii++ {
		switch boardingPass[ii] {
		case 'F':
			continue
		case 'B':
			seat.row += (1 << (6 - ii)) // jank arithmetic
		default:
			return seat, errors.Errorf("couldn't parse character %c in %s", boardingPass[ii], boardingPass)
		}
	}

	for ii := 0; ii < 3; ii++ {
		switch boardingPass[7+ii] {
		case 'L':
			continue
		case 'R':
			seat.col += (1 << (2 - ii)) // more jank arithmetic
		default:
			return seat, errors.Errorf("couldn't parse character %c in %s", boardingPass[ii], boardingPass)
		}
	}

	return seat, nil
}

// seatID returns the seat ID of a plane seat
func (seat planeSeat) seatID() int {
	return seat.row*8 + seat.col
}

// Day05 solves the fifth day puzzle "Binary Boarding"
//
// Input
//
// A list of boarding passes, represented by 10-character strings. For example:
//
//    FFBFBFFRLR
//    BBBFFBFLLL
//    BFFBBBBRRL
//    BBFBFFFRRR
//    FBFFFFFRLL
//    BFBFBBFLRL
//    BFBFBFBLRL
//    BFFFBFFLLL
//    BBFBBBFRRL
//    FBFFBBFRRR
//
// It is guaranteed that the first 7 characters for each boarding pass
// consist of either `B` or `F`, and the last 3 consist of either `R` or `L`.
func Day05(input string) (answer1, answer2 string, err error) {

	allSeats := make([]planeSeat, 0)
	highestID := 0 // all IDs are at least 0

	for _, line := range aoc.SplitLines(input) {
		seat, e := readBoardingPass(line)
		if e != nil {
			err = errors.Wrapf(e, "couldn't read string %s", line)
			return
		}
		allSeats = append(allSeats, seat)

		if id := seat.seatID(); id > highestID {
			highestID = id
		}
	}
	answer1 = strconv.Itoa(highestID)

	// For part2, we are looking for a boarding pass whose ID isn't in allSeats
	// but isn't lowest-k nor highest+k for any k>=1.
	// That is, your boarding pass is somewhere in the middle of the list.
	//
	// The algorithm is as follows:
	// 1. Sort allSeats via seatID. Set expected to the lowest seatID
	// 2. For each pass in sorted allSeats
	// 2a. Check if pass.seatID is greater than expected.
	// 2b. If so, return expected. Otherwise, increment expected and continue.
	sort.Slice(allSeats, func(i, j int) bool {
		return allSeats[i].seatID() < allSeats[j].seatID()
	})
	expected := allSeats[0].seatID()
	for _, seat := range allSeats {
		if seat.seatID() > expected { // i.e., expected does not exist
			break
		}
		expected = seat.seatID() + 1
	}
	answer2 = strconv.Itoa(expected)
	return
}

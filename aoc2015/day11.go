package aoc2015

import (
	"bufio"
	"strings"
)

// password is a collection of eight bytes
// whose contents are 'a' to 'z'.
// password[7] will be the least significant byte
// and password[0] the most significant.
type password [8]byte

// newPassword returns a password from the first eight bytes of an input.
// All bytes in input which are lexicographically less than 'a' will be turned into 'a'
// and all bytes greater than 'z' will be turned into 'z'.
// This means that newPassword does not support uppercase input.
func newPassword(input string) password {
	var ps [8]byte
	if len(input) > 8 {
		// get the first 8 bytes
		input = input[:8]
	}
	if len(input) < 8 {
		// if input == "pqrs" change it to "aaaapqrs"
		addendLength := 8 - len(input)
		var sb strings.Builder
		for ii := 0; ii < addendLength; ii++ {
			sb.WriteByte('a')
		}
		sb.WriteString(input)
		input = sb.String()
	}
	for ii := range ps {
		if input[ii] < 'a' {
			ps[ii] = 'a'
			continue
		}
		if input[ii] > 'z' {
			ps[ii] = 'z'
			continue
		}
		ps[ii] = input[ii]
	}
	return ps
}

func (ps password) string() string {
	return string(ps[:])
}

func (ps password) copy() password {
	var result password
	for ii := range ps {
		result[ii] = ps[ii]
	}
	return result
}

// hasIncreasing returns true if it includes three characters that are
// continuously increasing e.g., "abc", "bcd", and so on.
func (ps password) hasIncreasing() bool {
	for ii := 0; ii < 5; ii++ {
		if ps[ii] == ps[ii+1]-1 && ps[ii+1] == ps[ii+2]-1 {
			return true
		}
	}
	return false
}

// noForbidden returns true if it does not contain the letters "i", "o", or "l".
func (ps password) noForbidden() bool {
	return !strings.ContainsAny(ps.string(), "iol")
}

// hasTwoPairs returns true if it contains at least two
// different, non-overlapping pairs of letters.
func (ps password) hasTwoPairs() bool {
	ind := 0
	for ; ind < 7; ind++ {
		if ps[ind] == ps[ind+1] {
			break
		}
	}
	for ind += 2; ind < 7; ind++ {
		if ps[ind] == ps[ind+1] {
			return true
		}
	}
	return false
}

// increment increments the password by one.
func (ps *password) increment() {
	ind := 7
	for ind >= 0 {
		ps[ind]++
		if ps[ind] <= 'z' {
			break
		}
		ps[ind] = 'a'
		ind--
	}
	return
}

// Day11 solves the eleventh day puzzle "Corporate Policy".
//
// Input
//
// A single line containing the password, as described in the puzzle:
// eight lowercase letters. For example:
//
//	hepxcrrq
func Day11(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	// do things
	scanner.Scan()
	passwd := newPassword(scanner.Text())
	for !passwd.noForbidden() || !passwd.hasIncreasing() || !passwd.hasTwoPairs() {
		passwd.increment()
	}
	answer1 = passwd.string()
	passwd.increment()
	for !passwd.noForbidden() || !passwd.hasIncreasing() || !passwd.hasTwoPairs() {
		passwd.increment()
	}
	answer2 = passwd.string()
	return
}

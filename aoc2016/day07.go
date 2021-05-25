package aoc2016

import (
	"strconv"
	"strings"

	aoc "github.com/janreggie/aoc/internal"
)

// ipv7Address is a list of either bracketed or unbracketed strings
type ipv7Address []struct {
	raw       string
	bracketed bool
}

// newIpv7Address creates a new ipv7 address
func newIpv7Address(input string) ipv7Address {
	result := make(ipv7Address, 0)
	var sb strings.Builder
	bracketed := false
	for _, vv := range input {
		if vv == '[' || vv == ']' {
			result = append(result, struct {
				raw       string
				bracketed bool
			}{sb.String(), bracketed})
			bracketed = !bracketed
			sb.Reset()
			continue
		}
		sb.WriteRune(vv)
	}
	result = append(result, struct {
		raw       string
		bracketed bool
	}{sb.String(), bracketed})

	return result
}

// supportTLS returns true if an ipv7Address supports TLS (Part 1)
func (ipv7 ipv7Address) supportTLS() bool {
	checkABBA := func(input string) bool {
		for ii := 0; ii < len(input)-3; ii++ {
			if input[ii] != input[ii+1] &&
				input[ii] == input[ii+3] &&
				input[ii+1] == input[ii+2] {
				return true
			}
		}
		return false
	}

	found := false // if we've found something

	for _, segment := range ipv7 {
		if segment.bracketed {
			if checkABBA(segment.raw) {
				return false
			}
			continue
		}
		if !segment.bracketed && !found {
			found = checkABBA(segment.raw)
		}
	}
	return found
}

// supportSSL returns true if an ipv7Address supports SSL (Part 2)
func (ipv7 ipv7Address) supportSSL() bool {
	// create two maps that contain all 3-letter combinations...
	unbracketedStrings, bracketedStrings := make(map[string]struct{}), make(map[string]struct{})
	// reverse: "aba" -> "bab"; only for length 3
	reverse := func(input string) string {
		if len(input) != 3 {
			return ""
		}
		var sb strings.Builder
		sb.WriteByte(input[1])
		sb.WriteByte(input[0])
		sb.WriteByte(input[1])
		return sb.String()
	}
	// checkFlip: "aba" is true, "abb" is false; only for length 3
	checkFlip := func(input string) bool {
		if len(input) != 3 {
			return false
		}
		return input[0] == input[2] && input[0] != input[1]
	}

	for _, segment := range ipv7 {
		for ii := 0; ii < len(segment.raw)-2; ii++ {
			if flip := segment.raw[ii : ii+3]; checkFlip(flip) {
				if segment.bracketed {
					bracketedStrings[flip] = struct{}{}
					if _, ok := unbracketedStrings[reverse(flip)]; ok {
						return true
					}
				} else {
					unbracketedStrings[flip] = struct{}{}
					if _, ok := bracketedStrings[reverse(flip)]; ok {
						return true
					}
				}
			}
		}
	}
	return false
}

// Day07 solves the seventh day puzzle "Internet Protocol Version 7".
//
// Input
//
// A file containing at most 2000 lines, each line representing an IPv7 address
// containing the lowercase letters `a` to `z` and the bracket characters `[` and `]`,
// each line of length at most 150.
//
// For example:
//     abba[mnop]qrst
//     ioxxoj[asdfgh]zxcvbn
//     wvwhsvattcfkxsjbif[bkesznrpxrlcnsmhbxv]rdycrgrtwazfqlx[genmoutcoedshrn]ucufntphttrvzmgt
func Day07(input string) (answer1, answer2 string, err error) {

	a1, a2 := 0, 0
	for _, line := range aoc.SplitLines(input) {
		ipv7 := newIpv7Address(line)
		if ipv7.supportTLS() {
			a1++
		}
		if ipv7.supportSSL() {
			a2++
		}
	}

	answer1 = strconv.Itoa(a1)
	answer2 = strconv.Itoa(a2)
	return
}

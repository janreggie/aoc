package aoc2016

import (
	"bufio"
	"strings"
)

// repertoire represents an array of dictionaries containing bytes
// and their frequency in a given column.
type repertoire []map[byte]uint

// newRepertoire creates a repertoire.
// Note that repertoire is a dynamic structure
// so length could be set to zero but reallocate will just add new elements.
func newRepertoire(length uint) repertoire {
	rp := make([]map[byte]uint, length)
	for ii := range rp {
		rp[ii] = make(map[byte]uint)
	}
	return rp
}

// length returns the length of the repertoire.
func (rp repertoire) length() uint {
	return uint(len(rp))
}

// reallocate reallocates the length of repertoire to a new length.
// If length is no more than rp.length, does nothing.
func (rp *repertoire) reallocate(length uint) {
	if length <= rp.length() {
		return
	}
	nrp := make(repertoire, length)
	copy(nrp, *rp)
	for ii := rp.length(); ii < length; ii++ {
		nrp[ii] = make(map[byte]uint)
	}

	*rp = nrp // for some reason rp = &nrp doesn't work...
}

// add adds some string to the repertoire.
// For example, if input is "abc", then
// rp[0]['a']++, rp[1]['b']++, and rp[2]['c']++.
// Runs reallocate first.
func (rp *repertoire) add(input string) {
	rp.reallocate(uint(len(input)))
	for ii := range input {
		(*rp)[ii][input[ii]]++
	}
}

// mostCommon returns the error-corrected version of the message
// as described by Year 2016 Day 6 Part 1
// by checking the most common byte for each map.
func (rp repertoire) mostCommon() string {
	var sb strings.Builder

	for ii := range rp {
		var record struct {
			count  uint
			holder byte
		}
		for kk, vv := range rp[ii] {
			if vv > record.count {
				record.count = vv
				record.holder = kk
			}
		}
		sb.WriteByte(record.holder)
	}

	return sb.String()
}

// leastCommon returns the error-corrected version of the message
// as described by Year 2016 Day 6 Part 2
// by checking the least common byte for each map.
func (rp repertoire) leastCommon() string {
	var sb strings.Builder

	for ii := range rp {
		var record struct {
			count  uint
			holder byte
		}
		record.count = 1<<32 - 1 // highest possible
		for kk, vv := range rp[ii] {
			if vv < record.count {
				record.count = vv
				record.holder = kk
			}
		}
		sb.WriteByte(record.holder)
	}

	return sb.String()
}

// Day06 solves the sixth day puzzle "Signals and Noise".
//
// Input
//
// A file containing some number of lines each having the same length
// as one another. For example:
//
//	jsgoijzv
//	iwgxjyxi
//	yzeeuwoi
//	gmgisfmd
//	vdtezvan
//	secfljup
//	dngzexve
//	xzanwmgd
//	ziobunnv
//
// It is guaranteed that all lines have the same length
// and that there are no more than 600 lines.
func Day06(input string) (answer1, answer2 string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	rp := newRepertoire(0)
	for scanner.Scan() {
		rp.add(scanner.Text())
	}
	answer1 = rp.mostCommon()
	answer2 = rp.leastCommon()
	return
}

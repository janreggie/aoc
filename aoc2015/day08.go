package aoc2015

import (
	"bufio"
	"strconv"

	"github.com/golang/glog"
)

// Day08 solves the eighth day puzzle "Matchsticks".
//
// Input
//
// A file containing 300 lines where each line represents a "string literal",
// which contains combinations such as '\"' and '\x27'. For example:
//
//	"ubgxxcvnltzaucrzg\\xcez"
//	"pnocjvo\\yt"
//	"fcabrtqog\"a\"zj"
//	"o\\bha\\mzxmrfltnflv\xea"
//	"isos\x6azbnkojhxoopzetbj\xe1yd"
//
// Each line is observed to start and end with a double-quotation mark '"'.
func Day08(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	var representedDiff uint64 // character difference (representation and actual) for answer 1
	var encodedDiff uint64     // character difference (literal and encoded) for answer 2

	// literalLength is length of a strength
	literalLength := func(str string) int {
		return len(str)
	}

	// stringLength is length of a "string" while considering "escaped" characters
	// e.g., `"aaa\"aaa"` is 6.
	stringLength := func(str string) int {
		// assume that str[0] and str[len(str)-1] is `"`
		if str[0] == '"' {
			str = str[1:]
		}
		if str[len(str)-1] == '"' {
			str = str[:len(str)-1]
		}
		// create an iterative function for this...
		// apparently you can't create recursive closures :sad:
		totalLength := 0
		glog.Infof("String to be used is %v\n", str)
		for ind := 0; ind < len(str); ind++ {
			totalLength++ // either way it'll iterate
			if str[ind] == '\\' {
				if str[ind+1] == 'x' {
					ind += 3
				} else {
					ind++
				}
				continue
			}
		}
		return totalLength
	}

	// encodedLength determines the length of the "encoded" string
	// e.g., `"aaa\"aaa"` -> `"\"aaa\\\"aaa\""` (16)
	encodedLength := func(str string) int {
		totalLength := 2 // air quotes in final answer!
		for _, letter := range str {
			if letter == '"' || letter == '\\' {
				totalLength += 2
			} else {
				totalLength++
			}
		}
		return totalLength
	}

	for scanner.Scan() {
		eachLine := scanner.Text()
		representedDiff += uint64(literalLength(eachLine) - stringLength(eachLine))
		encodedDiff += uint64(encodedLength(eachLine) - literalLength(eachLine))
	}

	answer1 = strconv.FormatUint(representedDiff, 10)
	answer2 = strconv.FormatUint(encodedDiff, 10)
	return
}

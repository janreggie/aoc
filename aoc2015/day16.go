package aoc2015

import (
	"bufio"
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

// unknownAuntParameter is an unknown Aunt Sue parameter
const unknownAuntParameter uint = 4294967295

// auntSue is a representation of an Aunt named Sue.
type auntSue struct {
	number      uint // "Sue 1" to "Sue 500"
	children    uint // by human DNA age analysis
	cats        uint // It doesn't differentiate individual breeds.
	samoyeds    uint // Several seemingly random breeds of dog
	pomeranians uint
	akitas      uint
	vizslas     uint
	goldfish    uint // No other kinds of fish.
	trees       uint // all in one group.
	cars        uint // presumably by exhaust or gasoline or something
	perfumes    uint // which is handy, since many of your Aunts Sue wear a few kinds
}

// newAuntSue returns a new Aunt Sue object from a string.
// The string would look like the following:
//
//	Sue 22: perfumes: 7, children: 1, pomeranians: 7
//
// Note that there are some variables that are undefined.
// Undefined variables will get a value of unknownAuntParameter,
// the largest possible uint.
func newAuntSue(input string) (auntSue, error) {
	output := auntSue{number: 0,
		children:    unknownAuntParameter,
		cats:        unknownAuntParameter,
		samoyeds:    unknownAuntParameter,
		pomeranians: unknownAuntParameter,
		akitas:      unknownAuntParameter,
		vizslas:     unknownAuntParameter,
		goldfish:    unknownAuntParameter,
		trees:       unknownAuntParameter,
		cars:        unknownAuntParameter,
		perfumes:    unknownAuntParameter}

	splitInput := strings.Fields(input)
	// assume that splitInput's length is always even
	if len(splitInput)%2 != 0 {
		return output, errors.Errorf("input %v has odd number of fields", input)
	}

	// check the first two values
	if len(splitInput) <= 2 {
		return output, errors.Errorf("input %v is too short", input)
	}
	if splitInput[0] != "Sue" {
		return output, errors.Errorf("input %v is not Sue", input)
	}
	splitInput[1] = strings.TrimSuffix(splitInput[1], ":")
	number, err := strconv.ParseUint(splitInput[1], 10, 32)
	if err != nil {
		return output, errors.Wrapf(err, "invalid number for Sue %v", splitInput[1])
	}
	output.number = uint(number)

	// now for the remaining fields
	for ii := 2; (ii + 1) < len(splitInput); ii += 2 {
		// splitInput[ii]: key
		// splitInput[ii+1]: value
		splitInput[ii] = strings.TrimSuffix(splitInput[ii], ":")
		splitInput[ii+1] = strings.TrimSuffix(splitInput[ii+1], ",")
		value, err := strconv.ParseUint(splitInput[ii+1], 10, 32)
		if err != nil {
			return output, errors.Wrapf(err, "invalid value %v", splitInput[ii+1])
		}

		// now where do we place value..?
		switch splitInput[ii] {
		case "children":
			output.children = uint(value)
		case "cats":
			output.cats = uint(value)
		case "samoyeds":
			output.samoyeds = uint(value)
		case "pomeranians":
			output.pomeranians = uint(value)
		case "akitas":
			output.akitas = uint(value)
		case "vizslas":
			output.vizslas = uint(value)
		case "goldfish":
			output.goldfish = uint(value)
		case "trees":
			output.trees = uint(value)
		case "cars":
			output.cars = uint(value)
		case "perfumes":
			output.perfumes = uint(value)
		default:
			return output, errors.Errorf("unknown key %v", splitInput[ii])
		}
	}

	return output, nil
}

// howSimilar returns a score to determine how many fields aunt is similar to
// another aunt Sue
func (aunt auntSue) howSimilar(another auntSue) uint {
	var score uint
	// oh my god...
	if aunt.children == another.children &&
		aunt.children != unknownAuntParameter &&
		another.children != unknownAuntParameter {
		score++
	}
	if aunt.cats == another.cats &&
		aunt.cats != unknownAuntParameter &&
		another.cats != unknownAuntParameter {
		score++
	}
	if aunt.samoyeds == another.samoyeds &&
		aunt.samoyeds != unknownAuntParameter &&
		another.samoyeds != unknownAuntParameter {
		score++
	}
	if aunt.pomeranians == another.pomeranians &&
		aunt.pomeranians != unknownAuntParameter &&
		another.pomeranians != unknownAuntParameter {
		score++
	}
	if aunt.akitas == another.akitas &&
		aunt.akitas != unknownAuntParameter &&
		another.akitas != unknownAuntParameter {
		score++
	}
	if aunt.vizslas == another.vizslas &&
		aunt.vizslas != unknownAuntParameter &&
		another.vizslas != unknownAuntParameter {
		score++
	}
	if aunt.goldfish == another.goldfish &&
		aunt.goldfish != unknownAuntParameter &&
		another.goldfish != unknownAuntParameter {
		score++
	}
	if aunt.trees == another.trees &&
		aunt.trees != unknownAuntParameter &&
		another.trees != unknownAuntParameter {
		score++
	}
	if aunt.cars == another.cars &&
		aunt.cars != unknownAuntParameter &&
		another.cars != unknownAuntParameter {
		score++
	}
	if aunt.perfumes == another.perfumes &&
		aunt.perfumes != unknownAuntParameter &&
		another.perfumes != unknownAuntParameter {
		score++
	}
	return score
}

// howSimilarRecalibrated is like howSimilar but is now calibrated
// such that it checks if another's number of cats and trees
// is greater than aunt's cats and trees
// and if another's pomeranians and goldfish are fewer than aunt's.
func (aunt auntSue) howSimilarRecalibrated(another auntSue) uint {
	var score uint
	// oh my god...
	if aunt.children == another.children &&
		aunt.children != unknownAuntParameter &&
		another.children != unknownAuntParameter {
		score++
	}
	if aunt.cats < another.cats &&
		aunt.cats != unknownAuntParameter &&
		another.cats != unknownAuntParameter {
		score++
	}
	if aunt.samoyeds == another.samoyeds &&
		aunt.samoyeds != unknownAuntParameter &&
		another.samoyeds != unknownAuntParameter {
		score++
	}
	if aunt.pomeranians > another.pomeranians &&
		aunt.pomeranians != unknownAuntParameter &&
		another.pomeranians != unknownAuntParameter {
		score++
	}
	if aunt.akitas == another.akitas &&
		aunt.akitas != unknownAuntParameter &&
		another.akitas != unknownAuntParameter {
		score++
	}
	if aunt.vizslas == another.vizslas &&
		aunt.vizslas != unknownAuntParameter &&
		another.vizslas != unknownAuntParameter {
		score++
	}
	if aunt.goldfish > another.goldfish &&
		aunt.goldfish != unknownAuntParameter &&
		another.goldfish != unknownAuntParameter {
		score++
	}
	if aunt.trees < another.trees &&
		aunt.trees != unknownAuntParameter &&
		another.trees != unknownAuntParameter {
		score++
	}
	if aunt.cars == another.cars &&
		aunt.cars != unknownAuntParameter &&
		another.cars != unknownAuntParameter {
		score++
	}
	if aunt.perfumes == another.perfumes &&
		aunt.perfumes != unknownAuntParameter &&
		another.perfumes != unknownAuntParameter {
		score++
	}
	return score
}

// Day16 solves the sixteenth day puzzle "Aunt Sue".
//
// Input
//
// A file containing 500 lines, each of which describes an Aunt Sue
// of a certain number, as well as three things that the
// person knows about a particular Sue. For example:
//
//	Sue 436: samoyeds: 3, trees: 2, cars: 6
//	Sue 437: children: 9, akitas: 0, pomeranians: 3
//	Sue 438: perfumes: 10, akitas: 2, cars: 7
//	Sue 439: perfumes: 10, samoyeds: 6, akitas: 10
//	Sue 440: vizslas: 10, trees: 2, akitas: 8
//
// It is guaranteed that the keys will be three of the following:
// "children", "cats", "samoyeds", "pomerians", "akitas",
// "vizslas", "goldfish", "trees", "cars", and "perfumes",
// and that the values are non-negative integers no more than 10.
func Day16(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	// basis...
	var basis = auntSue{
		children:    3,
		cats:        7,
		samoyeds:    2,
		pomeranians: 3,
		akitas:      0,
		vizslas:     0,
		goldfish:    5,
		trees:       3,
		cars:        2,
		perfumes:    1,
	}

	// construct a list of aunts
	allAunts := make([]auntSue, 0, 500)
	for scanner.Scan() {
		aunt, e := newAuntSue(scanner.Text())
		if e != nil {
			err = errors.Wrapf(e, "could not properly create: %v", aunt)
			return
		}
		allAunts = append(allAunts, aunt)
	}

	// Note:
	// The computation performed here is trivial and inexpensive,
	// and implementing concurrency and goroutines in solving this puzzle is unnecessary,
	// and even causes a dip in performance.
	// Nevertheless, I have implemented goroutines for the sake of demonstration,
	// to show that these problems *can* be solved concurrently.

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		// now check which one has the same value...
		// check for each Aunt...
		for _, aunt := range allAunts {
			if basis.howSimilar(aunt) == 3 {
				answer1 = strconv.FormatUint(uint64(aunt.number), 10)
				break
			}
		}
		wg.Done()
	}()

	go func() {
		// now for Part two...
		for _, aunt := range allAunts {
			if basis.howSimilarRecalibrated(aunt) == 3 {
				answer2 = strconv.FormatUint(uint64(aunt.number), 10)
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()
	return
}

package aoc2020

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// passPolicy represents a password policy
type passPolicy struct {
	a int
	b int
	c rune
}

// newPolicy creates a policy object from a string formatted `m-M x`.
// If invalid, throws an error.
func newPolicy(policy string) (*passPolicy, error) {
	result := &passPolicy{}
	_, err := fmt.Sscanf(policy, "%d-%d %c", &result.a, &result.b, &result.c)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't parse policy string `%v` properly", policy)
	}

	return result, nil
}

// checkPasswordCount returns true if password is valid for the following policy:
// The number of instances `.c` occurs in `password` is no less than `.a` and no more than `.b`
func (policy *passPolicy) checkPasswordCount(password string) bool {
	countChr := 0
	for _, char := range password {
		if char == policy.c {
			countChr++
		}
	}

	return policy.a <= countChr && countChr <= policy.b
}

// checkPasswordPos returns true if password is valid for the following policy:
// Either `password` at position `.a` or `.b` (but not both!) equals `.c`.
// Note that `.a` and `.b` index from 1.
// If `password` is too long for `.a` or `.b`, return false.
func (policy *passPolicy) checkPasswordPos(password string) bool {
	passSplit := []rune(password)
	if len(passSplit) < policy.a || len(passSplit) < policy.b {
		return false
	}

	// make sure both have different truth values!!
	return (passSplit[policy.a-1] == policy.c) != (passSplit[policy.b-1] == policy.c)
}

// Day02 solves the second day puzzle "Password Philosophy".
//
// Input
//
// A list of lines in the format `m-M x: PASS`. For example:
//
//  1-3 a: abcde
//  1-3 b: cdefg
//  2-9 c: ccccccccc
//
// `m` and `M` are guaranteed to be numbers no less than 1 and no more than 20.
// `x` is guaranteed to be a single character.
// `PASS` is guaranteed to be of length no more than 20.
func Day02(input string) (answer1, answer2 string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	validFirst, validSecond := 0, 0

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), ": ")
		if len(splitLine) != 2 {
			err = errors.Errorf("`%v` could not be split properly", scanner.Text())
			return
		}

		policy, e := newPolicy(splitLine[0])
		if e != nil {
			err = errors.Wrapf(e, "couldn't parse line `%v`", scanner.Text())
			return
		}

		password := splitLine[1]
		if policy.checkPasswordCount(password) {
			validFirst++
		}
		if policy.checkPasswordPos(password) {
			validSecond++
		}
	}

	answer1 = strconv.Itoa(validFirst)
	answer2 = strconv.Itoa(validSecond)
	return
}

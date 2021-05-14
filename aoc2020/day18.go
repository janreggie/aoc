package aoc2020

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// parseEquation returns a sequence of symbols in the form of a string slice from a given string output.
func parseEquation(eqn string) []string {
	// add spaces to parentheses so that they may be parsed more easily
	eqn = strings.ReplaceAll(eqn, "(", "( ")
	eqn = strings.ReplaceAll(eqn, ")", " )")
	return strings.Fields(eqn)
}

// simpleMath solves an equation represented by a parsed sequence of symbols
// from left to right.
func simpleMath(eqn []string) (int, error) {
	if len(eqn) == 0 {
		return 0, errors.Errorf("can't solve an empty equation")
	}

	if len(eqn) == 1 {
		ans, err := strconv.Atoi(eqn[0])
		if err != nil {
			return 0, errors.Errorf("could not parse lone variable %s", eqn[0])
		}
		return ans, nil
	}

	if len(eqn) == 2 {
		return 0, errors.Errorf("could not solve an equation with two parameters %v", eqn)
	}

	// Now, look for RHS and then LHS...
	rhsStart, rhsEnd, rhsVal := len(eqn)-1, len(eqn), 0
	if eqn[rhsStart] != ")" {
		var e error
		rhsVal, e = simpleMath([]string{eqn[rhsStart]})
		if e != nil {
			return 0, errors.Wrapf(e, "could not parse RHS value %v", eqn[rhsStart])
		}
	} else {
		parenLevel := 1
		rhsStart--
		for ; rhsStart > 0 && parenLevel != 0; rhsStart-- {
			if eqn[rhsStart] == "(" {
				parenLevel--
			} else if eqn[rhsStart] == ")" {
				parenLevel++
			}
		}
		if rhsStart == 0 { // i.e., ended up at the start of the equation
			return simpleMath(eqn[1 : rhsEnd-1])
		}
		rhsStart++ // start with parenthesis
		var e error
		rhsVal, e = simpleMath(eqn[rhsStart:rhsEnd])
		if e != nil {
			return 0, errors.Wrapf(e, "could not solve for RHS %v", eqn[rhsStart:rhsEnd])
		}
	}
	op := eqn[rhsStart-1]

	lhsVal, err := simpleMath(eqn[:rhsStart-1])
	if err != nil {
		return 0, errors.Wrapf(err, "could not solve for LHS %v", eqn[:rhsStart-1])
	}

	if op == "+" {
		return lhsVal + rhsVal, nil
	}
	if op == "*" {
		return lhsVal * rhsVal, nil
	}
	return 0, errors.Errorf("could not recognize operator %v", op)
}

// advancedMath solves an equation using advanced math principles with addition preceding multiplication.
func advancedMath(eqn []string) (int, error) {
	if len(eqn) == 0 || len(eqn) == 2 {
		return 0, errors.Errorf("equation %v invalid length %d", eqn, len(eqn))
	}
	if len(eqn) == 1 {
		ans, err := strconv.Atoi(eqn[0])
		if err != nil {
			return 0, errors.Wrapf(err, "could not parse number %v", eqn[0])
		}
		return ans, nil
	}

	// check if the entire thing is enclosed in one parenthesis
	if eqn[0] == "(" {
		parenLevel := 1
		ii := 1
		for ; ii < len(eqn) && parenLevel != 0; ii++ {
			if eqn[ii] == "(" {
				parenLevel++
			} else if eqn[ii] == ")" {
				parenLevel--
			}
		}
		if ii == len(eqn) {
			return advancedMath(eqn[1 : ii-1])
		}
	}

	getLeftmost := func(eqn []string) []string {
		if eqn[0] != "(" {
			return eqn[0:1]
		}
		parenLvl := 1
		ind := 1
		for ; ind < len(eqn) && parenLvl != 0; ind++ {
			if eqn[ind] == "(" {
				parenLvl++
			} else if eqn[ind] == ")" {
				parenLvl--
			}
		}
		return eqn[0:ind]
	}

	leftMost := getLeftmost(eqn)
	opInd := len(leftMost)
	op := eqn[opInd]
	lhs, err := advancedMath(leftMost)
	if err != nil {
		return 0, errors.Wrapf(err, "could not handle LHS %v", leftMost)
	}

	if op == "*" {
		rest, err := advancedMath(eqn[opInd+1:])
		if err != nil {
			return 0, errors.Wrapf(err, "could not resolve rest of the equation %v", eqn[opInd+1:])
		}
		return lhs * rest, nil
	}

	if op == "+" {
		next := getLeftmost(eqn[opInd+1:])
		last := eqn[opInd+len(next)+1:]
		rhs, err := advancedMath(next)
		if err != nil {
			return 0, errors.Wrapf(err, "could not handle RHS %v", next)
		}
		newEqn := make([]string, len(last)+1)
		newEqn[0] = strconv.Itoa(lhs + rhs)
		copy(newEqn[1:], last)
		return advancedMath(newEqn)
	}

	return 0, errors.Errorf("could not parse equation %v", op)
}

// Day18 solves the eighteenth day puzzle "Operation Order"
//
// Input
//
// A list of valid equations to be solved. For example:
//
//    1 + 2 * 3 + 4 * 5 + 6
//    1 + (2 * 3) + (4 * (5 + 6))
//    2 * 3 + (4 * 5)
//    5 + (8 * 3 + 9 + 3 * 4 * 3)
//
// It is guaranteed that the equations only contain numerical integers, the operations `+` and `*`, and balanced parentheses.
func Day18(input string) (answer1, answer2 string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	equations := make([][]string, 0)
	for scanner.Scan() {
		equations = append(equations, parseEquation(scanner.Text()))
	}
	result1, result2 := 0, 0
	for _, eqn := range equations {
		simpleAnswer, e := simpleMath(eqn)
		if e != nil {
			err = errors.Wrapf(e, "could not solve equation %v", eqn)
			return
		}
		result1 += simpleAnswer
		advancedAnswer, e := advancedMath(eqn)
		if e != nil {
			err = errors.Wrapf(e, "could not solve equation %v in advanced mode", eqn)
			return
		}
		result2 += advancedAnswer
	}

	answer1 = strconv.Itoa(result1)
	answer2 = strconv.Itoa(result2)

	return
}

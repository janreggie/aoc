package aoc2020

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// messageRuleset is a set of rules for which messages ought to be valid
type messageRuleset []struct {
	subRules [][]int // subRules[0] is a subrule, where a message ought to be valid
	pattern  string  // rule must follow pattern if subrule does not exist
}

// newMessageRuleset returns a new messageRuleset object
func newMessageRuleset() messageRuleset {
	return make(messageRuleset, 0)
}

// addRule adds a rule to the ruleset. If the parsed string is "invalid", return an errro.
func (ruleset *messageRuleset) addRule(rule string) error {
	splitRule := strings.Split(rule, ": ")
	if len(splitRule) != 2 {
		return errors.Errorf("could not parse rule %v: cannot split via `: ` properly", rule)
	}

	ruleNumber, err := strconv.Atoi(splitRule[0])
	if err != nil {
		return errors.Wrapf(err, "could not parse rule number %s", splitRule[0])
	}
	if ruleNumber >= len(*ruleset) {
		*ruleset = append(*ruleset, make(messageRuleset, ruleNumber-len(*ruleset)+1)...)
	}

	ruleMessage := splitRule[1]
	if strings.HasPrefix(ruleMessage, `"`) && strings.HasSuffix(ruleMessage, `"`) {
		ruleMessage = strings.TrimPrefix(ruleMessage, `"`)
		ruleMessage = strings.TrimSuffix(ruleMessage, `"`)
		(*ruleset)[ruleNumber].subRules = nil
		(*ruleset)[ruleNumber].pattern = ruleMessage
	} else {
		(*ruleset)[ruleNumber].pattern = ""
		subRules := strings.Split(ruleMessage, " | ")
		(*ruleset)[ruleNumber].subRules = make([][]int, len(subRules))
		for ii := range subRules {
			subRule := strings.Split(subRules[ii], " ")
			(*ruleset)[ruleNumber].subRules[ii] = make([]int, len(subRule))
			for jj, nn := range subRule {
				reference, err := strconv.Atoi(nn)
				if err != nil {
					return errors.Wrapf(err, "could not turn %s into an integer from subrule %s", nn, subRules[ii])
				}
				(*ruleset)[ruleNumber].subRules[ii][jj] = reference
			}
		}
	}

	return nil
}

// check returns if a string matches the rule defined at ind.
// Note that if rule defined at ind does not exist, it may return true if str is empty.
// IF ind is out of bounds, return false.
func (ruleset messageRuleset) check(ind int, str string) bool {
	// recurse is the recursive function we'd want.
	// It returns if any of strs has a prefix matching ind's ruleset and, if there is a match, return an appropriate list of strings.
	// If there is no match, return an empty slice.
	var recurse func(ind int, strs []string) []string
	recurse = func(ind int, strs []string) []string {
		if ind < 0 || ind >= len(ruleset) || len(strs) == 0 {
			return []string{}
		}
		rule := ruleset[ind]

		if rule.subRules == nil {
			potentialMatches := []string{}
			for _, str := range strs {
				hasPrefix, trimmed := strings.HasPrefix(str, rule.pattern), strings.TrimPrefix(str, rule.pattern)
				if hasPrefix {
					potentialMatches = append(potentialMatches, trimmed)
				}
			}
			return potentialMatches
		}

		result := []string{}
	SUBRULES:
		for _, subRule := range rule.subRules {
			toCheck := make([]string, len(strs))
			copy(toCheck, strs)
			for _, field := range subRule {
				toCheck = recurse(field, toCheck)
				if len(toCheck) == 0 {
					continue SUBRULES
				}
			}
			if len(toCheck) > 0 {
				result = append(result, toCheck...)
			}
		}
		return result
	}

	result := recurse(ind, []string{str})
	for _, res := range result {
		// An empty string means that all values in the string has been exhausted
		if res == "" {
			return true
		}
	}
	return false
}

// Day19 solves the nineteenth day puzzle "Monster Messages"
//
// Input
//
// A list of "rules", followed by a list of messages. For example:
//
//   0: 4 1 5
//   1: 2 3 | 3 2
//   2: 4 4 | 5 5
//   3: 4 5 | 5 4
//   4: "a"
//   5: "b"
//
//   ababbb
//   bababa
//   abbbab
//   aaabbb
//   aaaabbb
//
// It is guaranteed that the rules are numbers which reference to numbers or strings,
// and the messages only consist of `a` and `b`.
// It is also guaranteed that, if there are N rules, they are numbered 0 to N-1,
// although they are not necessarily in order.
func Day19(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	ruleset := newMessageRuleset()
	messages := []string{}
	writeMessages := false
	for scanner.Scan() {
		if scanner.Text() == "" {
			writeMessages = true
			continue
		}

		if !writeMessages {
			if err = ruleset.addRule(scanner.Text()); err != nil {
				err = errors.Wrapf(err, "could not parse rule %s", scanner.Text())
				return
			}
		} else {
			messages = append(messages, scanner.Text())
		}
	}

	validMessages := 0
	for _, msg := range messages {
		if ruleset.check(0, msg) {
			validMessages++
		}
	}
	answer1 = strconv.Itoa(validMessages)

	// now for part 2
	newRules := []string{
		"8: 42 | 42 8",
		"11: 42 31 | 42 11 31",
	}
	for _, rule := range newRules {
		if err = ruleset.addRule(rule); err != nil {
			err = errors.Wrapf(err, "could not parse rule %s", rule)
			return
		}
	}

	validMessages = 0
	for _, msg := range messages {
		if ruleset.check(0, msg) {
			validMessages++
		}
	}
	answer2 = strconv.Itoa(validMessages)

	return
}

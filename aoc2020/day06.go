package aoc2020

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// customsResponse represents a person's response from a customs form.
type customsResponse [26]bool

func (response customsResponse) String() string {
	var sb strings.Builder
	for _, rr := range response {
		if rr {
			sb.WriteRune('X')
		} else {
			sb.WriteRune('.')
		}
	}
	return sb.String()
}

// newCustomsResponse generates a customsResponse from a string representing his answers (e.g., `qebagdfvhr`).
// If there are invalid responses in customsResponse, return an error.
func newCustomsResponse(answers string) (customsResponse, error) {
	result := customsResponse{}
	for _, cc := range answers {
		if cc < 'a' || cc > 'z' {
			err := errors.Errorf("could not parse `%s`: invalid character `%c`", answers, cc)
			return customsResponse{}, err
		}
		result[cc-'a'] = true
	}
	return result, nil
}

// countYes returns how many yes's (true's) are there in customsResponse
func (response customsResponse) countYes() int {
	result := 0
	for _, rr := range response {
		if rr {
			result++
		}
	}
	return result
}

// customsResponseGroup represents a group of people whose responses are recorded
type customsResponseGroup []customsResponse

// atLeastOne returns a customsResponse for a customsResponseGroup where at least one in the group has answered yes.
func (group customsResponseGroup) atLeastOne() customsResponse {
	var result customsResponse
	for _, response := range group {
		for ii, rr := range response {
			if rr {
				result[ii] = true
			}
		}
	}
	return result
}

// everyone returns a customsResponse for a customsResponseGroup where everyone in the group has answered yes.
func (group customsResponseGroup) everyone() customsResponse {
	if len(group) == 0 {
		return customsResponse{} // no one has answered yes
	}

	var result customsResponse = group[0]
	for _, response := range group {
		for ii, rr := range response {
			if !rr {
				result[ii] = false
			}
		}
	}
	return result
}

// generateCustomsResponseGroups creates a slice of customResponseGroups from an input
func generateCustomsResponseGroups(scanner *bufio.Scanner) ([]customsResponseGroup, error) {
	result := make([]customsResponseGroup, 0)
	current := make(customsResponseGroup, 0)
	for scanner.Scan() {
		if scanner.Text() == "" && len(current) != 0 {
			result = append(result, current)
			current = make(customsResponseGroup, 0)
			continue
		}

		response, err := newCustomsResponse(scanner.Text())
		if err != nil {
			return nil, errors.Wrapf(err, "could not generate response groups because of %s", scanner.Text())
		}
		current = append(current, response)
	}

	if len(current) != 0 { // we might have forgotten to append it...
		result = append(result, current)
	}

	return result, nil
}

// Day06 solves the sixth day puzzle "Custom Customs"
//
// Input
//
// A file containing characters that indicate whether a person in a group has answered yes to a particular question. For example:
//
//   mxuwh
//   hwuxm
//   uhxmw
//   hwumx
//   hwuxm
//
//   k
//   k
//   tl
//   k
//
//   qebagdfvhr
//   alvkif
//   yufaovwi
//   fivsa
//   nwifazovu
//
// The file should only contain newlines and the characters `a` to `z`.
func Day06(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	responseGroups, err := generateCustomsResponseGroups(scanner)
	if err != nil {
		err = errors.Wrapf(err, "could not read from scanner")
		return
	}

	sumAtLeastOne := 0
	sumEveryone := 0
	for _, group := range responseGroups {
		sumAtLeastOne += group.atLeastOne().countYes()
		sumEveryone += group.everyone().countYes()
	}
	answer1 = strconv.Itoa(sumAtLeastOne)
	answer2 = strconv.Itoa(sumEveryone)

	return
}

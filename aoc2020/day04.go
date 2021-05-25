package aoc2020

import (
	"strconv"
	"strings"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/pkg/errors"
)

// passport represents passport data
type passport struct {
	byr string // Birth Year
	iyr string // Issue Year
	eyr string // Expiration Year
	hgt string // Height
	hcl string // Hair Color
	ecl string // Eye Color
	pid string // Passport ID
	cid string // Country ID
}

// valid checks if passport is "valid" (part1)
func (pp passport) valid() bool {
	return pp.byr != "" &&
		pp.iyr != "" &&
		pp.eyr != "" &&
		pp.hgt != "" &&
		pp.hcl != "" &&
		pp.ecl != "" &&
		pp.pid != ""
	// pp.cid == "" // not needed since the system ignores cid
}

// isInBounds returns true if num is an integer within given bounds
func isInBounds(raw string, min, max int) bool {
	if num, err := strconv.Atoi(raw); err != nil || num < min || num > max {
		return false
	}
	return true
}

func (pp passport) checkBYR() bool {
	return isInBounds(pp.byr, 1920, 2002)
}

func (pp passport) checkIYR() bool {
	return isInBounds(pp.iyr, 2010, 2020)
}

func (pp passport) checkEYR() bool {
	return isInBounds(pp.eyr, 2020, 2030)
}

func (pp passport) checkHGT() bool {
	if strings.HasSuffix(pp.hgt, "cm") {
		rawHgt := strings.TrimSuffix(pp.hgt, "cm")
		return isInBounds(rawHgt, 150, 193)
	}

	if strings.HasSuffix(pp.hgt, "in") {
		rawHgt := strings.TrimSuffix(pp.hgt, "in")
		return isInBounds(rawHgt, 59, 76)
	}

	return false
}

func (pp passport) checkHCL() bool {
	if len(pp.hcl) != 7 || pp.hcl[0] != '#' {
		return false
	}
	hexColor := pp.hcl[1:]

	// check if valid
	for _, rr := range hexColor {
		if (rr < '0' || rr > '9') && (rr < 'a' || rr > 'f') {
			return false
		}
	}
	return true
}

func (pp passport) checkECL() bool {
	validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, validColor := range validColors {
		if pp.ecl == validColor {
			return true
		}
	}
	return false
}

func (pp passport) checkPID() bool {
	if len(pp.pid) != 9 {
		return false
	}
	if _, err := strconv.Atoi(pp.pid); err != nil {
		return false
	}
	return true
}

// strictValid checks if passport is strictly valid (part2)
func (pp passport) strictValid() bool {
	return pp.checkBYR() &&
		pp.checkIYR() &&
		pp.checkEYR() &&
		pp.checkHGT() &&
		pp.checkHCL() &&
		pp.checkECL() &&
		pp.checkPID()
}

func (pp *passport) readKeyVal(key, val string) error {
	switch key {
	case "byr":
		pp.byr = val
	case "iyr":
		pp.iyr = val
	case "eyr":
		pp.eyr = val
	case "hgt":
		pp.hgt = val
	case "hcl":
		pp.hcl = val
	case "ecl":
		pp.ecl = val
	case "pid":
		pp.pid = val
	case "cid":
		pp.cid = val
	default:
		return errors.Errorf("could not handle key %s", key)
	}

	return nil
}

// readLine reads from a line of text and updates the current passport.
func (pp *passport) readLine(line string) error {
	keyVals := strings.Fields(line)
	for _, keyValRaw := range keyVals {
		keyVal := strings.Split(keyValRaw, ":")
		if len(keyVal) != 2 {
			return errors.Errorf("could not parse keyval pair %s properly", keyValRaw)
		}
		if err := pp.readKeyVal(keyVal[0], keyVal[1]); err != nil {
			return errors.Wrapf(err, "could not parse keyval pair %s properly", keyValRaw)
		}
	}

	return nil
}

// readPassportFile reads from a bufio.Scanner and returns a slice of passport data
func readPassportFile(input string) ([]passport, error) {
	result := make([]passport, 0)

	for _, passportData := range strings.Split(input, "\n\n") {
		current := passport{}
		for _, line := range aoc.SplitLines(passportData) {
			if err := (&current).readLine(line); err != nil {
				return nil, errors.Wrapf(err, "could not read line %v properly", line)
			}
		}
		result = append(result, current)
	}

	return result, nil
}

// Day04 solves the fourth day puzzle "Passport Processing"
//
// Input
//
// A list of passport data in key:value pairs, where each passport is separated by a newline. For example:
//
//  ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
//  byr:1937 iyr:2017 cid:147 hgt:183cm
//
//  iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
//  hcl:#cfa07d byr:1929
//
//  hcl:#ae17e1 iyr:2013
//  eyr:2024
//  ecl:brn pid:760753108 byr:1931
//  hgt:179cm
//
//  hcl:#cfa07d eyr:2025 pid:166559648
//  iyr:2011 ecl:brn hgt:59in
//
// All key:value pairs are strings that do not contain spaces, and it is guaranteed that the keys are as follows:
//
//  byr (Birth Year)
//  iyr (Issue Year)
//  eyr (Expiration Year)
//  hgt (Height)
//  hcl (Hair Color)
//  ecl (Eye Color)
//  pid (Passport ID)
//  cid (Country ID)
func Day04(input string) (answer1, answer2 string, err error) {

	allPassports, err := readPassportFile(input)
	if err != nil {
		errors.Wrapf(err, "could not read file properly")
		return
	}

	countValid := 0
	countStrictValid := 0
	for _, passport := range allPassports {
		if passport.valid() {
			countValid++
		}
		if passport.strictValid() {
			countStrictValid++
		}
	}

	answer1 = strconv.Itoa(countValid)
	answer2 = strconv.Itoa(countStrictValid)
	return
}

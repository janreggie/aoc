package aoc2020

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_readPassportFile(t *testing.T) {
	assert := assert.New(t)
	scanned, err := readPassportFile(day04sampleInput)
	assert.NoError(err)
	assert.ElementsMatch(
		scanned,
		[]passport{
			{ecl: "gry", pid: "860033327", eyr: "2020", hcl: "#fffffd", byr: "1937", iyr: "2017", cid: "147", hgt: "183cm"},
			{iyr: "2013", ecl: "amb", cid: "350", eyr: "2023", pid: "028048884", hcl: "#cfa07d", byr: "1929"},
			{hcl: "#ae17e1", iyr: "2013", eyr: "2024", ecl: "brn", pid: "760753108", byr: "1931", hgt: "179cm"},
			{hcl: "#cfa07d", eyr: "2025", pid: "166559648", iyr: "2011", ecl: "brn", hgt: "59in"},
		})
}

func Test_strictValid(t *testing.T) {
	assert := assert.New(t)
	allInvalid := `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`

	allValid := `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`

	testCases := []struct {
		input string
		total int // number of passports
		valid int // number of strictly valid passports
	}{
		{allInvalid, 4, 0},
		{allValid, 4, 4},
	}

	for _, tt := range testCases {
		passports, err := readPassportFile(tt.input)
		assert.NoError(err)
		assert.Len(passports, tt.total)
		countValid := 0
		for _, pp := range passports {
			if pp.strictValid() {
				countValid++
			}
		}
		assert.Equal(tt.valid, countValid)
	}
}

func TestDay04(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{
			Details: "Y2020D04 sample input",
			Input:   day04sampleInput,
			Result1: "2",
			Result2: "",
		},
		{
			Details: "Y2020D04 my input",
			Input:   day04myInput,
			Result1: "230",
			Result2: "156",
		},
	}
	for _, tt := range testCases {
		tt.Test(Day04, assert)
	}
}

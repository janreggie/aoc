package aoc2016

import (
	"sort"
	"strconv"
	"strings"
	"unicode"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/pkg/errors"
)

// room represents a room (Year 2016 Day 4).
//
// Syntax
//
// RAW is represented by ENCRYPTEDNAME-SECTORID[CHECKSUM],
// where ENCRYPTEDNAME is a string of length at least 1 contains lowercase letters separated by dashes,
// SECTORID is a three-digit integer,
// and CHECKSUM is a string of length five consisting solely of lowercase letters.
type room struct {
	raw      string          // raw representation
	letters  map[byte]uint16 // letters in ENCRYPTEDNAME
	sectorID uint16
	checksum string
}

// newRoom creates a room struct reflecting the syntax of room.
// Will return an error if syntax is invalid.
func newRoom(raw string) (room, error) {
	out := room{}
	out.raw = raw

	// Minimum string: a-100[abcde] is length 12
	if len(raw) < 12 {
		return out, errors.Errorf("%v is too short", raw)
	}

	// split everything into dashes
	splitRaw := strings.Split(raw, "-")

	// parse the letters
	out.letters = make(map[byte]uint16)
	for ii := 0; ii < len(splitRaw)-1; ii++ {
		for _, letter := range splitRaw[ii] {
			if !unicode.IsLower(letter) {
				return out, errors.Errorf("%v in %v not lowercase letter", letter, splitRaw[ii])
			}
			out.letters[byte(letter)]++
		}
	}

	// check last item in splitRaw
	lastString := splitRaw[len(splitRaw)-1]
	if len(lastString) != 10 || lastString[3] != '[' || lastString[9] != ']' {
		return out, errors.Errorf("%v cannot be parsed", lastString)
	}

	parsedInt, err := strconv.ParseUint(lastString[0:3], 10, 16)
	if err != nil {
		return out, errors.Wrapf(err, "could not pase %v", lastString[0:3])
	}
	out.sectorID = uint16(parsedInt)

	for _, letter := range lastString[4:9] {
		if !unicode.IsLower(letter) {
			return out, errors.Errorf("%v in %v not lowercase letter", letter, lastString[4:9])
		}
	}
	out.checksum = lastString[4:9]

	return out, nil
}

// isReal returns true if the room is "real" (Part 1)
// by checking its checksum and letters.
func (rm room) isReal() bool {
	if len(rm.checksum) < 5 {
		// too short
		return false
	}

	// sort all rm.letters...
	type letterCount struct {
		letter byte
		count  uint16
	}
	allLetters := make([]letterCount, 0, len(rm.letters))
	for kk, vv := range rm.letters {
		allLetters = append(allLetters, letterCount{kk, vv})
	}
	sort.Slice(allLetters, func(i, j int) bool {
		// check counts
		if allLetters[i].count > allLetters[j].count {
			return true
		}
		// check letters
		if allLetters[i].count == allLetters[j].count &&
			allLetters[i].letter < allLetters[j].letter {
			return true
		}
		return false
	})

	for ii := range rm.checksum {
		if allLetters[ii].letter != rm.checksum[ii] {
			return false
		}
	}
	return true
}

// decrypt decrypts the message found
func (rm room) decrypt() string {
	input := strings.Split(strings.Split(rm.raw, "[")[0], "-")
	input = input[:len(input)-1]
	key := byte(rm.sectorID % 26)
	parse := func(input byte) byte {
		return 0x61 + (input-0x61+key)%26
	}
	output := make([]string, 0, len(input))
	for _, istr := range input {
		var sb strings.Builder
		for ii := range istr {
			sb.WriteByte(parse(istr[ii]))
		}
		output = append(output, sb.String())
	}
	return strings.Join(output, " ")
}

// Day04 solves the fourth day puzzle "Security Through Obscurity".
//
// Input
//
// A file where each line represents a room, which is of the format
// ENCRYPTEDNAME-SECTORID[CHECKSUM], where ENCRYPTED NAME consists of lowercase
// letters separated by dashes, SECTORID is an integer between 100 and 999,
// and CHECKSUM is a string containing five lowercase letters. For example:
//
//	ipvohghykvbz-ihzrla-jbzavtly-zlycpjl-253[lzhvy]
//	cybyjqho-whqtu-rqiauj-fkhsxqiydw-322[syzwi]
//	tipfxvezt-sleep-tljkfdvi-jvimztv-425[tveif]
//	ktiaaqnqml-xtiabqk-oziaa-xczkpiaqvo-616[aiqko]
//	ckgvutofkj-xghhoz-gtgreyoy-306[nyhpz]
//
// The input contains at most 1000 lines,
// and it is guaranteed that ENCRYPTEDNAME is at most of length 60.
func Day04(input string) (answer1, answer2 string, err error) {
	var a1 uint64

	for _, line := range aoc.SplitLines(input) {
		rm, e := newRoom(line)
		if e != nil {
			err = errors.Wrapf(e, "could not parse %v", line)
			return
		}
		if !rm.isReal() {
			continue
		}
		a1 += uint64(rm.sectorID)
		if rm.decrypt() == "northpole object storage" {
			answer2 = strconv.FormatUint(uint64(rm.sectorID), 10)
		}
	}

	answer1 = strconv.FormatUint(a1, 10)
	return
}

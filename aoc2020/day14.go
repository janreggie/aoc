package aoc2020

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// mask represents a bitmask that can force bits to be zero or one.
type mask struct {
	ones   uint64
	zeroes uint64
}

// newMask generates a mask object from a string of length 32
func newMask(mm string) (mask, error) {
	if len(mm) != 36 {
		return mask{}, errors.Errorf("could not parse %v since not length 32", mm)
	}
	var ones, zeroes uint64
	for ii := range mm {
		switch mm[ii] {
		case '1':
			ones += 1 << (35 - ii)
		case '0':
			zeroes += 1 << (35 - ii)
		}
	}
	return mask{ones: ones, zeroes: zeroes}, nil
}

func (mm mask) String() string {
	var sb strings.Builder
	for ii := 0; ii < 36; ii++ {
		switch mm.at(ii) {
		case 1:
			sb.WriteByte('1')
		case 0:
			sb.WriteByte('0')
		default:
			sb.WriteByte('X')
		}
	}
	return sb.String()
}

// at checks the value of mm at some index.
// Returns 0, 1, or -1 (X).
// If index is oob return -2.
func (mm mask) at(ind int) int {
	if ind < 0 || ind >= 36 {
		return -2
	}
	if mm.ones|1<<(35-ind) == mm.ones {
		return 1
	}
	if mm.zeroes|1<<(35-ind) == mm.zeroes {
		return 0
	}
	return -1
}

// apply maps a value using the mask
func (mm mask) apply(val uint64) uint64 {
	return (val | mm.ones) &^ mm.zeroes
}

// v2floaters returns a slice of uint64's that represent the possible addresses when addr is masked by mm
func (mm mask) v2floaters(addr uint64) []uint64 {
	// addrAt returns addr at some index. Note that ind=0 means leftmost
	addrAt := func(ind int) uint64 {
		return addr & (1 << (35 - ind))
	}

	// recurse recurses. Note that ind=0 means leftmost
	var recurse func(int) []uint64
	recurse = func(ind int) []uint64 {
		if ind >= 36 {
			return []uint64{}
		}
		if ind == 35 {
			switch mm.at(ind) {
			case 1:
				return []uint64{1}
			case 0:
				return []uint64{addrAt(35)}
			case -1:
				return []uint64{0, 1}
			}
		}

		result := make([]uint64, 0)
		switch mm.at(ind) {
		case 1:
			result = recurse(ind + 1)
			for ii := range result {
				result[ii] += 1 << (35 - ind)
			}
		case 0:
			result = recurse(ind + 1)
			aa := addrAt(ind)
			for ii := range result {
				result[ii] += aa
			}
		case -1:
			result = recurse(ind + 1)
			// this will not loop indefinitely
			for ii := range result {
				result = append(result, result[ii]+1<<(35-ind))
			}
		}
		return result
	}
	return recurse(0)
}

// ferryInstruction refers to an instruction for the ferry
type ferryInstruction struct {
	setMask   bool // set the mask or just write to memory?
	mask      mask
	writeAddr uint64
	writeVal  uint64
}

// newFerryInstruction reads from a string and evaluates the appropriate ferryInstruciton
func newFerryInstruction(instr string) (ferryInstruction, error) {
	var rawMask string
	if _, e := fmt.Sscanf(instr, "mask = %36s", &rawMask); e == nil {
		mask, err := newMask(rawMask)
		if err != nil {
			return ferryInstruction{}, errors.Wrapf(err, "couldn't read mask val %s", rawMask)
		}
		return ferryInstruction{setMask: true, mask: mask}, nil
	}

	var d1, d2 uint64
	if _, e := fmt.Sscanf(instr, "mem[%d] = %d", &d1, &d2); e == nil {
		return ferryInstruction{
			setMask:   false,
			writeAddr: d1,
			writeVal:  d2,
		}, nil
	}

	return ferryInstruction{}, errors.Errorf("couldn't read instruction `%s`", instr)
}

// readFerryInstructions reads a slice of ferryInstruction and returns the state of the memory as a map
func readFerryInstructions(instrs []ferryInstruction) map[uint64]uint64 {
	mem := make(map[uint64]uint64)
	msk := mask{ones: 0, zeroes: 0} // blank mask

	for _, instr := range instrs {
		if instr.setMask {
			msk = instr.mask
		} else {
			mem[instr.writeAddr] = msk.apply(instr.writeVal)
		}
	}

	return mem
}

// readFerryV2 reads a slice of ferryInstruction in accordance to version 2.0 and returns the state of the memory as a map
func readFerryV2(instrs []ferryInstruction) map[uint64]uint64 {
	mem := make(map[uint64]uint64)
	msk := mask{ones: 0, zeroes: 0}

	for _, instr := range instrs {
		if instr.setMask {
			msk = instr.mask
		} else {
			for _, addr := range msk.v2floaters(instr.writeAddr) {
				mem[addr] = instr.writeVal
			}
		}
	}
	return mem
}

// Day14 solves the fourteenth day puzzle "Docking Data"
//
// Input
//
// An initialization program that updates the bitmask or inserts a value into memory.
// For example:
//
//   mask = 1100X10X01001X111001X00010X00100X011
//   mem[24821] = 349
//   mem[34917] = 13006
//   mem[53529] = 733
//   mem[50289] = 245744
//   mem[23082] = 6267
//   mask = 011X1X00X100100XXXX11100X0000100X010
//   mem[21316] = 14188354
//   mem[53283] = 7137
//   mem[57344] = 62358
//   mem[63867] = 9443
//
// It is guaranteed that the mask represents a 36-bit number.
func Day14(input string) (answer1, answer2 string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	instrs := make([]ferryInstruction, 0)
	for scanner.Scan() {
		instr, e := newFerryInstruction(scanner.Text())
		if e != nil {
			err = errors.Wrapf(e, "couldn't read ferryInstructions from scanner")
			return
		}
		instrs = append(instrs, instr)
	}

	mem := readFerryInstructions(instrs)
	var sum uint64
	for _, vv := range mem {
		sum += vv
	}
	answer1 = strconv.FormatUint(sum, 10)

	mem = readFerryV2(instrs)
	sum = 0
	for _, vv := range mem {
		sum += vv
	}
	answer2 = strconv.FormatUint(sum, 10)

	return
}

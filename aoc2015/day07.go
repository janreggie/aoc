package aoc2015

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/pkg/errors"
)

// wireComputer represents the "computer" in the Day07 puzzle
type computer struct {
	wireMemory   map[identifier]signal
	instructions map[identifier]*instruction
}

// newComputer generates a new computer
func newComputer() *computer {
	return &computer{
		wireMemory:   make(map[identifier]signal),
		instructions: make(map[identifier]*instruction),
	}
}

// addInstruction adds an instruction to the computer
func (com *computer) addInstruction(instr string) error {

	actual, err := parseInstruction(instr)
	if err != nil {
		return errors.Wrapf(err, "couldn't parse instruction %v", instr)
	}

	com.instructions[actual.id] = actual
	return nil
}

// clear clears the computer memory and forces instructions to be "undone"
func (com *computer) clear() {

	for id := range com.wireMemory {
		com.wireMemory[id] = 0
	}

	for id := range com.instructions {
		com.instructions[id].done = false
	}
}

// get gets the signal at some id and will run some instructions if they aren't done yet
func (com *computer) get(id identifier) (signal, error) {

	instr, ok := com.instructions[id]
	if !ok {
		return 0, errors.Errorf("could not find instruction with id %v", id)
	}
	if instr.done {
		return com.wireMemory[id], nil
	}

	for _, child := range instr.childrenID {
		_, err := com.get(child)
		if err != nil {
			return 0, errors.Wrapf(err, "error at evaluating child %v of %v", child, id)
		}
	}

	err := instr.function(com.wireMemory, instr.id, instr.parameters)
	if err != nil {
		return 0, errors.Wrapf(err, "could not evaluate at id %v", id)
	}
	instr.done = true
	return com.wireMemory[id], nil
}

// identifier is used to identify a memory location
type identifier string

// toSignal converts an identifier to a signal and true if applicable
// or 0 and false if not.
func (id identifier) toSignal() (signal, bool) {
	val, err := strconv.Atoi(string(id))
	return signal(val), err == nil
}

// signal is an alias for an unsigned 16-bit integer
type signal uint16 // 0 to 65535

// instructionFunc is a function that takes in memory state, address, and parameters
// and returns an applicable error
type instructionFunc func(mem map[identifier]signal, address identifier, params []identifier) error

// parseParams parses the parameters ([]identifier)
// and, using a map[identifier]signal,
// returns the appropriate values
//
func parseParams(mem map[identifier]signal, params []identifier) ([]signal, error) {
	answer := make([]signal, len(params))
	for ind, param := range params {
		val, isInt := param.toSignal()
		if isInt {
			answer[ind] = signal(val)
		} else {
			val, found := mem[param]
			if !found {
				return nil, fmt.Errorf("%v not found in memory", param)
			}
			answer[ind] = val
		}
	}
	return answer, nil
}

// removeNumbers removes all integers in params and returns a resulting []identifier
func removeNumbers(params []identifier) []identifier {
	answer := make([]identifier, 0)
	for _, item := range params {
		if _, isInt := item.toSignal(); !isInt {
			answer = append(answer, item)
		}
	}
	return answer
}

// The following instruction functions are for the Day 07 puzzle
//
var (
	// assign assigns a value
	//	"44430 -> b"
	//	params = {"44430"}
	//	mem["b"] = 44430
	// Note that the paramter can either be a signal or an identifier
	// e.g., "lx -> a".
	assign instructionFunc = func(mem map[identifier]signal, address identifier, params []identifier) error {
		if len(params) != 1 {
			return fmt.Errorf("assign error: params %v not length 1", params)
		}
		p, err := parseParams(mem, params)
		if err != nil {
			return fmt.Errorf("and error: %v", err)
		}
		mem[address] = p[0]
		return nil
	}

	// and performs bitwise and
	//	"eg AND 3 -> ej" turns into
	//	params = {"eg","3"}
	//	mem["ej"] = mem["eg"] & 3
	// Note that either parameter can be a signal or an identifier
	and instructionFunc = func(mem map[identifier]signal, address identifier, params []identifier) error {
		if len(params) != 2 {
			return fmt.Errorf("and error: params %v not length 2", params)
		}
		p, err := parseParams(mem, params)
		if err != nil {
			return fmt.Errorf("and error: %v", err)
		}
		mem[address] = p[0] & p[1]
		return nil
	}

	// or performs bitwise or
	//	"eg OR 3 -> ej" turns into
	//	params = {"eg","3"}
	//	mem["ej"] = mem["eg"] | 3
	// Note that either parameter can be a signal or an identifier
	or instructionFunc = func(mem map[identifier]signal, address identifier, params []identifier) error {
		// bitwise and e.g., "eg AND ei -> ej"; params={"eg","ei"}
		// NOTE THAT EITHER PARAMETER CAN BE A NUMBER e.g., "1 OR ed"
		if len(params) != 2 {
			return fmt.Errorf("or error: params %v not length 2", params)
		}
		p, err := parseParams(mem, params)
		if err != nil {
			return fmt.Errorf("or error: %v", err)
		}
		mem[address] = p[0] | p[1]
		return nil
	}

	// not performs bitwise not
	//	"NOT h -> i" turns into
	//	params = {"h"}
	//	mem["i"] = ^mem["h"]
	// Note that the parameter can be a signal or an identifier
	not instructionFunc = func(mem map[identifier]signal, address identifier, params []identifier) error {
		// bitwise NOT e.g., "NOT h -> i"; params = {"h"}
		if len(params) != 1 {
			return fmt.Errorf("not error: params %v not length 1", params[0])
		}
		p, err := parseParams(mem, params)
		if err != nil {
			return fmt.Errorf("not error: %v", err)
		}
		mem[address] = ^p[0]
		return nil
	}

	// lshift performs bitwise left shift
	//	"eg LSHIFT 3 -> ej" turns into
	//	params = {"eg","3"}
	//	mem["ej"] = mem["eg"] << 3
	// Note that either parameter can be a signal or an identifier
	lshift instructionFunc = func(mem map[identifier]signal, address identifier, params []identifier) error {
		// bitwise shift e.g., "lv LSHIFT 15 -> lz"; params={"lv","15"}
		// NOTE THAT EITHER VALUE CAN BE AN INTEGER
		if len(params) != 2 {
			return fmt.Errorf("lshift error: params %v not length 2", params[0])
		}
		p, err := parseParams(mem, params)
		if err != nil {
			return fmt.Errorf("lshift error: %v", err)
		}
		mem[address] = p[0] << p[1]
		return nil
	}

	// rshift performs bitwise right shift
	//	"eg RSHIFT 3 -> ej" turns into
	//	params = {"eg","3"}
	//	mem["ej"] = mem["eg"] >> 3
	// Note that either parameter can be a signal or an identifier
	rshift instructionFunc = func(mem map[identifier]signal, address identifier, params []identifier) error {
		// bitwise shift e.g., "lv RSHIFT 15 -> lz"; params={"lv","15"}
		// NOTE THAT EITHER VALUE CAN BE AN INTEGER
		if len(params) != 2 {
			return fmt.Errorf("lshift error: params %v not length 2", params[0])
		}
		p, err := parseParams(mem, params)
		if err != nil {
			return fmt.Errorf("lshift error: %v", err)
		}
		mem[address] = p[0] >> p[1]
		return nil
	}
)

// instruction represents an instruction in Day07.
//
type instruction struct {
	id         identifier      // the identifier
	parameters []identifier    // addresses it is dependent on ("a AND b -> c" will have parameters={"a","b"})
	childrenID []identifier    // identifiers that will need to be evaluated to evaluate this->id.
	function   instructionFunc // the function to be done to our memory map
	done       bool            // has instruction been executed yet?
}

// parseInstruction parses a string and returns an instruction or an error in parsing.
//
func parseInstruction(raw string) (*instruction, error) {
	nd := &instruction{}
	query := strings.Split(raw, " -> ")
	if len(query) != 2 {
		return nil, fmt.Errorf("unknown instruction: %q", raw)
	}
	nd.id = identifier(query[1])
	// now determine what instruction should be undertaken
	splitLHS := strings.Fields(query[0])
	switch len(splitLHS) {
	case 1: // assign
		nd.parameters = []identifier{identifier(splitLHS[0])}
		nd.function = assign
	case 2: // most likely a NOT
		if splitLHS[0] != "NOT" {
			return nil, fmt.Errorf("unknown instruction: %v is not not", splitLHS[0])
		}
		nd.parameters = []identifier{identifier(splitLHS[1])}
		nd.function = not
	case 3: // AND, OR, LSHIFT, RSHIFT
		nd.parameters = []identifier{identifier(splitLHS[0]), identifier(splitLHS[2])}
		switch splitLHS[1] {
		case "AND":
			nd.function = and
		case "OR":
			nd.function = or
		case "LSHIFT":
			nd.function = lshift
		case "RSHIFT":
			nd.function = rshift
		default:
			return nil, fmt.Errorf("unknown function: %v", splitLHS[1])
		}
	default:
		return nil, fmt.Errorf("%v too long", splitLHS)
	}
	nd.childrenID = removeNumbers(nd.parameters)
	nd.done = false
	return nd, nil
}

func (instr *instruction) String() string {
	b := new(strings.Builder)
	b.WriteString(fmt.Sprintf("instruction %v\n", instr.id))
	b.WriteString(fmt.Sprintf("\tparameters: %v\n", instr.parameters))
	b.WriteString(fmt.Sprintf("\tfunction name: %v\n", instr.operationName()))
	b.WriteString(fmt.Sprintf("\tchildren: %v\n", instr.childrenID))
	b.WriteRune('\n')
	b.WriteString(fmt.Sprintf("\timplemented: %v\n", instr.done))
	return b.String()
}

// operationName returns the name of the function that instr contains
func (instr *instruction) operationName() string {
	s := strings.Split(runtime.FuncForPC(reflect.ValueOf(instr.function).Pointer()).Name(), ".")
	return s[len(s)-1]
}

// Day07 solves the seventh day puzzle "Some Assembly Required".
//
// Input
//
// A file containing several lines where each line represents an instruction.
// An instruction can be any of the following:
//
// Assignment instruction
//	VALUE -> IDENT
//	123 -> x
//
// AND instruction
//	IDENT AND IDENT -> IDENT
//	x AND y -> z
//
// OR instruction
//	IDENT OR IDENT -> IDENT
//	bj OR bi -> bk
//
// NOT instruction
//	NOT IDENT -> IDENT
//	NOT e -> f
//
// LSHIFT instruction
//	IDENT LSHIFT VALUE -> IDENT
//	p LSHIFT 2 -> q
//
// RSHIFT instruction
//	IDENT RSHIFT VALUE -> IDENT
//	p RSHIFT 2 -> q
//
// Note that IDENT is any alphabetic string at most length 2
// that represents a Signal, a 16-bit unsigned integer,
// and VALUE represents a raw Signal.
// Also note that VALUE is only a parameter for the LSHIFT, RSHIFT, and
// Assignment instructions.
//
func Day07(input string) (answer1, answer2 string, err error) {

	computer := newComputer()
	for _, rawQuery := range aoc.SplitLines(input) {
		if e := computer.addInstruction(rawQuery); e != nil {
			err = errors.Wrapf(e, "could not parse instruction %#v", rawQuery)
			return
		}
	}

	// Part 1
	//
	val, err := computer.get("a")
	if err != nil {
		err = errors.Wrapf(err, "could not get value at wire a in part1")
		return
	}
	answer1 = strconv.FormatUint(uint64(val), 10)

	// Part 2
	//
	computer.clear()
	computer.addInstruction(answer1 + " -> b") // Override instruction
	val, err = computer.get("a")
	if err != nil {
		err = errors.Wrapf(err, "could not get value at wire a in part2")
		return
	}
	answer2 = strconv.FormatUint(uint64(val), 10)

	return
}

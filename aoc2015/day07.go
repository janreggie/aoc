package aoc2015

import (
	"bufio"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

// isInteger checks whether a string is a number
// and returns true if so, as well as the number
func isInteger(str string) (int, bool) {
	val, err := strconv.Atoi(str)
	return val, err == nil
}

// identifier is used to identify a memory location
type identifier string // one/two-character address

// signal is an alias for an unsigned 16-bit integer
type signal uint16 // 0 to 65535

// instructionFunc is a function that takes in
// memory state, address, and parameters
// and returns an applicable error
type instructionFunc func(mem map[identifier]signal, address identifier, params []identifier) error

// parseParams parses the parameters ([]identifier)
// and, using a map[identifier]signal,
// returns the appropriate values
func parseParams(mem map[identifier]signal, params []identifier) ([]signal, error) {
	answer := make([]signal, len(params))
	for ind, param := range params {
		val, isInt := isInteger(string(param))
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
		if _, isInt := isInteger(string(item)); !isInt {
			answer = append(answer, item)
		}
	}
	return answer
}

// The following instruction functions are for the Day 07 puzzle
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

// instruction represents a line of assembly code
type instruction struct {
	id         identifier      // the identifier
	parameters []identifier    // addresses it is dependent on ("a AND b -> c" will have dependsOn={"b","c"})
	childrenID []identifier    // the children's Identifiers
	children   []*instruction  // its actual children which it is dependent on
	function   instructionFunc // the function to be done to our memory map
	done       bool            // has instruction been executed yet?
}

// parseInstruction parses a string
// and returns an instruction or an error
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
	nd.children = make([]*instruction, 0)
	nd.done = false
	return nd, nil
}

func (instr *instruction) String() string {
	b := new(strings.Builder)
	b.WriteString(fmt.Sprintf("instruction %v\n", instr.id))
	b.WriteString(fmt.Sprintf("\tparameters: %v\n", instr.parameters))
	b.WriteString(fmt.Sprintf("\tfunction name: %v\n", instr.operationName()))
	b.WriteString(fmt.Sprintf("\tchildren: %v\n", instr.childrenID))
	b.WriteString(fmt.Sprintf("\tactual children: "))
	for _, child := range instr.children {
		b.WriteString(fmt.Sprintf("%v,", child.id))
	}
	b.WriteRune('\n')
	b.WriteString(fmt.Sprintf("\timplemented: %v\n", instr.done))
	return b.String()
}

// operationName returns the name of the function
// that instr contains
func (instr *instruction) operationName() string {
	s := strings.Split(runtime.FuncForPC(reflect.ValueOf(instr.function).Pointer()).Name(), ".")
	return s[len(s)-1]
}

// checkChildren returns true if children is complete
func (instr *instruction) checkChildren() bool {
	actualChildren := make([]identifier, len(instr.children))
	for ind, child := range instr.children {
		actualChildren[ind] = child.id
	}
	if len(actualChildren) != len(instr.childrenID) {
		return false
	}
	// create a map of identifier -> int
	diff := make(map[identifier]int, len(actualChildren))
	for _, child := range actualChildren {
		// 0 value for int is 0, so just increment a counter for the identifier
		diff[child]++
	}
	for _, child := range instr.childrenID {
		// If the identifier child is not in diff bail out early
		if _, ok := diff[child]; !ok {
			return false
		}
		diff[child]--
		if diff[child] == 0 {
			delete(diff, child)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}

// findChildren looks for an instruction's children using an imap
func (instr *instruction) findChildren(imap *instructionMap) error {
	if instr.checkChildren() {
		return nil
	}
	isInSlice := func(slice []*instruction, val identifier) bool {
		for _, child := range slice {
			if child.id == val {
				return true
			}
		}
		return false
	}
	for _, child := range instr.childrenID {
		ptr, err := imap.lookup(child)
		if err != nil {
			return fmt.Errorf("cannot find children of %v: %v", instr.id, err)
		}
		if !isInSlice(instr.children, child) {
			instr.children = append(instr.children, ptr)
		}
	}
	return nil
}

// populateChildren populates all children of a certain node
func (instr *instruction) populateChildren(imap *instructionMap) error {
	// if it already has the children we skip
	if instr.checkChildren() {
		return nil
	}
	glog.Infof("populating children of %v which are %v", instr.id, instr.childrenID)
	if err := instr.findChildren(imap); err != nil {
		return fmt.Errorf("cannot populate children of %v: %v", instr.id, err)
	}
	for _, child := range instr.children {
		if err := child.populateChildren(imap); err != nil {
			return fmt.Errorf("cannot populate %v child of %v: %v", child.id, instr.id, err)
		}
	}
	return nil
}

// performInstruction performs the instruction of a node using some memory
func (instr *instruction) performInstruction(mem map[identifier]signal) error {
	if instr.done == true {
		return nil
	}
	if err := instr.function(mem, instr.id, instr.parameters); err != nil {
		return fmt.Errorf("error in performing %v's instruction: %v", instr.id, err)
	}
	glog.Infof("\t%v: %v(%v) ; %v = %v",
		instr.id, instr.operationName(), instr.parameters,
		instr.id, mem[instr.id])
	instr.done = true
	return nil
}

// performAll performs instructions of a node and its children
func (instr *instruction) performAll(mem map[identifier]signal) error {
	// remember that children need to be performed first
	if instr.done == true {
		return nil
	}
	glog.Infof("Perform all children of %v which are %v", instr.id, instr.childrenID)
	for _, child := range instr.children {
		if err := child.performAll(mem); err != nil {
			return err
		}
	}
	if err := instr.performInstruction(mem); err != nil {
		return err
	}
	instr.done = true
	return nil
}

// instructionMap is a way
// to store all instructions neatly
type instructionMap struct {
	m map[identifier]*instruction
}

// newInstructionMap creates an instruction map
func newInstructionMap() *instructionMap {
	m := make(map[identifier]*instruction)
	return &instructionMap{m: m}
}

func (imap *instructionMap) String() string {
	b := new(strings.Builder)
	for _, instr := range imap.m {
		b.WriteString(fmt.Sprintln(instr))
	}
	return b.String()
}

// Append appends to an instruction map
// and will log to warning if instruction already exists.
func (imap *instructionMap) append(is *instruction) {
	if _, found := imap.m[is.id]; found {
		glog.Warningf("append warning: %v already exists in map", is.id)
	}
	imap.m[is.id] = is
}

// Delete deletes an identifier from the instruction map
func (imap *instructionMap) delete(id identifier) {
	delete(imap.m, id)
}

// lookup looks from the instruction map
func (imap *instructionMap) lookup(id identifier) (*instruction, error) {
	ptr, found := imap.m[id]
	if !found {
		return nil, fmt.Errorf("cannot find instruction with id %v", id)
	}
	return ptr, nil
}

// traverse performs a function func(id identifier) through each key of imap
func (imap *instructionMap) traverse(f func(id identifier)) {
	for key := range imap.m {
		f(key)
	}
}

// Day07 solves the seventh day puzzle
// "Some Assembly Required"
func Day07(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	wires := make(map[identifier]signal) // map of all wires and their signal values
	imap1 := newInstructionMap()
	imap2 := newInstructionMap()

	// place everything in a instructionMap
	for scanner.Scan() {
		rawQuery := scanner.Text()
		instr1, e := parseInstruction(rawQuery)
		if e != nil {
			err = fmt.Errorf("cannot parse %v: ", e)
			return
		}
		instr2, e := parseInstruction(rawQuery)
		if e != nil {
			err = fmt.Errorf("cannot parse %v: ", e)
			return
		}
		imap1.append(instr1)
		imap2.append(instr2)
	}
	// make sure everyone found their children
	parent1, e := imap1.lookup("a") // parent1 of our "tree"
	if e != nil {
		err = fmt.Errorf("cannot find parent1 a: %v", err)
		return
	}
	glog.Infof("Populating parent1...\n")
	if err = parent1.populateChildren(imap1); err != nil {
		return
	}
	// fmt.Println(imap1)
	// now do all children operations...
	glog.Infof("Performing parent1's operations...\n")
	if err = parent1.performAll(wires); err != nil {
		return
	}
	for key, val := range wires {
		glog.Infof("%v: %v\n", key, val)
	}
	answer1 = strconv.Itoa(int(wires[parent1.id]))

	// reset wires and hijack b
	wires = make(map[identifier]signal)
	nodeB, e := imap1.lookup("b")
	if e != nil {
		err = e
		return
	}
	nodeB.parameters = []identifier{identifier(answer1)}
	nodeB.childrenID = []identifier{}
	nodeB.children = []*instruction{}
	nodeB.function = assign
	// "undo" all operations
	imap1.traverse(func(id identifier) {
		instr, _ := imap1.lookup(id)
		instr.done = false
	})
	glog.Infof("Performing parent1's operations...\n")
	if err = parent1.performAll(wires); err != nil {
		return
	}
	for key, val := range wires {
		glog.Infof("%v: %v\n", key, val)
	}
	answer2 = strconv.Itoa(int(wires[parent1.id]))

	return
}

package aoc2015

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/golang/glog"
)

// Day07 solves the seventh day problem
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
	fmt.Println(imap1)
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

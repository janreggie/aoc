package aoc2015

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/golang/glog"
)

// Day07 solves the seventh day problem
// "Some Assembly Required"
func Day07(scanner *bufio.Scanner) (string, string, error) {
	answer1, answer2 := "", ""
	wires := make(map[identifier]signal) // map of all wires and their signal values
	imap1 := newInstructionMap()
	imap2 := newInstructionMap()

	// place everything in a instructionMap
	for scanner.Scan() {
		rawQuery := scanner.Text()
		instr1, err := parseInstruction(rawQuery)
		if err != nil {
			return "", "", fmt.Errorf("cannot parse %v: ", err)
		}
		instr2, err := parseInstruction(rawQuery)
		if err != nil {
			return "", "", fmt.Errorf("cannot parse %v: ", err)
		}
		imap1.append(instr1)
		imap2.append(instr2)
	}
	// make sure everyone found their children
	parent1, err := imap1.lookup("a") // parent1 of our "tree"
	if err != nil {
		return "", "", fmt.Errorf("cannot find parent1 a: %v", err)
	}
	glog.Infof("Populating parent1...\n")
	if err := parent1.populateChildren(imap1); err != nil {
		return "", "", err
	}
	fmt.Println(imap1)
	// now do all children operations...
	glog.Infof("Performing parent1's operations...\n")
	if err := parent1.performAll(wires); err != nil {
		return "", "", err
	}
	for key, val := range wires {
		glog.Infof("%v: %v\n", key, val)
	}
	answer1 = strconv.Itoa(int(wires[parent1.id]))

	// reset wires and hijack b
	wires = make(map[identifier]signal)
	nodeB, err := imap1.lookup("b")
	if err != nil {
		return "", "", err
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
	if err := parent1.performAll(wires); err != nil {
		return "", "", err
	}
	for key, val := range wires {
		glog.Infof("%v: %v\n", key, val)
	}
	answer2 = strconv.Itoa(int(wires[parent1.id]))

	return answer1, answer2, nil
}

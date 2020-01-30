package aoc2019

import (
	"bufio"
	"strconv"

	"github.com/golang/glog"
	"github.com/janreggie/AdventOfCode/tools/intcode"
)

// Day09 solves the ninth day puzzle
// "Sensor Boost"
func Day09(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	ic, err := intcode.NewFromScanner(scanner)
	if err != nil {
		return
	}
	memory := ic.Snapshot()
	intcode.InstallAdderMultiplier(ic)
	intcode.InstallJumpers(ic)
	ic.Install(intcode.Inputter)
	ic.Install(intcode.Outputter)
	ic.Install(intcode.ChangeRelativeBase)
	// now run it.
	ic.PushToInput(1)
	if err = ic.Operate(); !intcode.IsHalt(err) {
		// there's a deeper reason here...
		if serr, isOpErr := err.(*intcode.OperationError); isOpErr {
			glog.Errorf("ERROR DETAILS")
			glog.Errorf("Memory: %v", serr.Memory)
			glog.Errorf("PC: %v (%v)", serr.PC, serr.Opcode)
			glog.Errorf("Inputs: %v", serr.Input)
			glog.Errorf("Outputs: %v", serr.Output)
			glog.Errorf("Underlying cause: %+v", serr.Child)
		}
		return
	}
	out, err := ic.GetOutput()
	if err != nil {
		return
	}
	answer1 = strconv.FormatInt(out, 10)

	// now for part 2
	ic.Format(memory)
	ic.PushToInput(2)
	if err = ic.Operate(); !intcode.IsHalt(err) {
		// there's a deeper reason here...
		if serr, isOpErr := err.(*intcode.OperationError); isOpErr {
			glog.Errorf("ERROR DETAILS")
			glog.Errorf("Memory: %v", serr.Memory)
			glog.Errorf("PC: %v (%v)", serr.PC, serr.Opcode)
			glog.Errorf("Inputs: %v", serr.Input)
			glog.Errorf("Outputs: %v", serr.Output)
			glog.Errorf("Underlying cause: %+v", serr.Child)
		}
		return
	}
	out, err = ic.GetOutput()
	if err != nil {
		return
	}
	answer2 = strconv.FormatInt(out, 10)

	return
}

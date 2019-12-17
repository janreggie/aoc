package intcode

import "fmt"

// OutOfBoundsError returns when accessing a memory location
// that is outside IntCode.mem
type OutOfBoundsError struct {
	msg          string
	LookingFor   int // what memory location is out of bounds?
	ActualLength int // how long is memory?
}

// NewOutOfBoundsError generates an OutOfBoundsError
func NewOutOfBoundsError(lookingFor, actualLength int) error {
	return &OutOfBoundsError{
		msg: fmt.Sprintf("wants to access %v, mem length %v",
			lookingFor, actualLength),
		LookingFor:   lookingFor,
		ActualLength: actualLength,
	}
}

func (e *OutOfBoundsError) Error() string {
	return e.msg
}

// InvalidOpcodeError returns when the opcode being read
// is not valid.
type InvalidOpcodeError struct {
	msg      string
	opcode   int    // what opcode did it see?
	Wants    int    // what does a Module want?
	mnemonic string // name of a Module
}

// NewInvalidOpcodeError returns an InvalidOpcodeError
// with message "invalid opcode opcode, wants Wants (mnemonic)"
func NewInvalidOpcodeError(opcode int, module *Module) error {
	var msg string
	if module.opcode != 0 {
		msg = fmt.Sprintf("invalid opcode %v, wants %v (%v)", opcode, module.opcode, module.mnemonic)
	} else {
		msg = fmt.Sprintf("invalid opcode %v, wants %v", opcode, module.mnemonic)
	}
	return &InvalidOpcodeError{
		msg:      msg,
		opcode:   opcode,
		Wants:    module.opcode,
		mnemonic: module.mnemonic,
	}
}

func (e *InvalidOpcodeError) Error() string {
	return e.msg
}

// HaltError is an error that essentially stops the IntCode from running
type HaltError struct {
	From string
}

// NewHaltError returns a halt error where "from" could be any reason
// e.g., "program halted from HALT (99)"
func NewHaltError(from string) error {
	return &HaltError{From: from}
}

func (e *HaltError) Error() string {
	return fmt.Sprintf("program halted from %v", e.From)
}

// IsHalt returns true if the error returned is a Halt statement
func IsHalt(err error) bool {
	_, result := err.(*HaltError)
	return result
}

package intcode

import "fmt"

// OutOfBoundsError returns when accessing a memory location
// that is outside Intcode.mem
type OutOfBoundsError struct {
	msg          string
	LookingFor   int64 // what memory location is out of bounds?
	ActualLength int64 // how long is memory?
}

// NewOutOfBoundsError generates an OutOfBoundsError
func NewOutOfBoundsError(lookingFor, actualLength int64) error {
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
	msg    string
	Opcode int64 // what opcode did it see?
	At     int64 // program counter?
}

// NewInvalidOpcodeError returns an InvalidOpcodeError
// with message "invalid opcode OPCODE at position POSITION"
func NewInvalidOpcodeError(opcode, position int64) error {
	var msg string
	msg = fmt.Sprintf("invalid opcode %d at position %d", opcode, position)
	return &InvalidOpcodeError{
		msg:    msg,
		Opcode: opcode,
		At:     position,
	}
}

func (e *InvalidOpcodeError) Error() string {
	return e.msg
}

// HaltError is an error that essentially stops the Intcode from running
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

// OperationError is an error that occurs during the operation of the Intcode computer
type OperationError struct {
	msg    string
	Child  error   // a "deeper" error
	PC     int64   // program counter
	Opcode int64   // the current opcode
	Memory []int64 // memory snapshot
	Input  []int64
	Output []int64
}

// NewOperationError creates an Error object that occurred due to some other error
func NewOperationError(err error, ic *Intcode) *OperationError {
	return &OperationError{
		msg:    fmt.Sprintf("operation error due to %v", err),
		Child:  err,
		Memory: ic.Snapshot(),
		PC:     ic.PC(),
		Opcode: ic.Current(),
		Input:  ic.Input(),
		Output: ic.Output(),
	}
}

func (e *OperationError) Error() string {
	return e.msg
}

package aoc2016

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// assembunnyComputer is a computer that could run an assembunny program.
// The zero value is ready to use.
type assembunnyComputer struct {
	registers struct {
		a int
		b int
		c int
		d int
	}

	// instructionPointer represents which instruction to use
	instructionPointer int
}

// Get gets the value of some register and some error
func (asm assembunnyComputer) Get(register string) (int, error) {
	switch register {
	case "a":
		return asm.registers.a, nil
	case "b":
		return asm.registers.b, nil
	case "c":
		return asm.registers.c, nil
	case "d":
		return asm.registers.d, nil
	}
	return 0, fmt.Errorf("could not get invalid register %v", register)
}

// Set sets the value of some register and returns an error
func (asm *assembunnyComputer) Set(register string, value int) error {
	switch register {
	case "a":
		asm.registers.a = value
		return nil
	case "b":
		asm.registers.b = value
		return nil
	case "c":
		asm.registers.c = value
		return nil
	case "d":
		asm.registers.d = value
		return nil
	}
	return fmt.Errorf("could not set invalid register %v", register)
}

// GetInstructionPointer gets the instruction pointer
func (asm *assembunnyComputer) GetInstructionPointer() int {
	return asm.instructionPointer
}

// SetInstructionPointer sets the instruction pointer
func (asm *assembunnyComputer) SetInstructionPointer(ip int) {
	asm.instructionPointer = ip
}

// inc increments register x by one and increments the instruction pointer.
func (asm *assembunnyComputer) inc(x string) error {
	v, err := asm.Get(x)
	if err != nil {
		return errors.Wrapf(err, "could not increment register %v", x)
	}
	asm.instructionPointer++
	return asm.Set(x, v+1)
}

// dec decrements register x by one and increments the instruction pointer.
func (asm *assembunnyComputer) dec(x string) error {
	v, err := asm.Get(x)
	if err != nil {
		return errors.Wrapf(err, "could not decrement register %v", x)
	}
	asm.instructionPointer++
	return asm.Set(x, v-1)
}

// cpy copies x into register y and increments the instruction pointer.
// x can be either an integer or the value of a register.
func (asm *assembunnyComputer) cpy(x, y string) error {
	v, err := strconv.Atoi(x)
	if err != nil {
		// it's a register
		v, err = asm.Get(x)
		if err != nil {
			return errors.Wrapf(err, "could not read register %v", x)
		}
	}
	if err = asm.Set(y, v); err != nil {
		return errors.Wrapf(err, "could not copy value %v to register %v", v, y)
	}
	asm.instructionPointer++
	return nil
}

// jnz jumps to an instruction y away only if x is not zero.
// If x is zero, iterate instruction pointer by 1.
func (asm *assembunnyComputer) jnz(x, y string) error {
	v, err := strconv.Atoi(x)
	if err != nil {
		// it's a register
		v, err = asm.Get(x)
		if err != nil {
			return errors.Wrapf(err, "could not read register %v", x)
		}
	}
	if v == 0 {
		asm.instructionPointer++
		return nil
	}
	j, err := strconv.Atoi(y)
	if err != nil {
		return errors.Wrapf(err, "could not parse %v as int", y)
	}
	asm.instructionPointer += j
	return nil
}

// readInstruction reads an instruction.
func (asm *assembunnyComputer) readInstruction(instr string) error {
	if len(instr) < 5 { // e.g., "inc a"
		return fmt.Errorf("%v too short", instr)
	}
	switch instr[0:4] {
	case "cpy ":
		args := strings.Fields(strings.TrimPrefix(instr, "cpy "))
		if len(args) != 2 {
			return fmt.Errorf("arguments %v invalid for cpy", args)
		}
		if err := asm.cpy(args[0], args[1]); err != nil {
			return errors.Wrapf(err, "could not complete instruction %v", instr)
		}
		return nil

	case "inc ":
		args := strings.Fields(strings.TrimPrefix(instr, "inc "))
		if len(args) != 1 {
			return fmt.Errorf("arguments %v invalid for inc", args)
		}
		if err := asm.inc(args[0]); err != nil {
			return errors.Wrapf(err, "could not complete instruction %v", instr)
		}
		return nil

	case "dec ":
		args := strings.Fields(strings.TrimPrefix(instr, "dec "))
		if len(args) != 1 {
			return fmt.Errorf("arguments %v invalid for dec", args)
		}
		if err := asm.dec(args[0]); err != nil {
			return errors.Wrapf(err, "could not complete instruction %v", instr)
		}
		return nil

	case "jnz ":
		args := strings.Fields(strings.TrimPrefix(instr, "jnz "))
		if len(args) != 2 {
			return fmt.Errorf("arguments %v invalid for jnz", args)
		}
		if err := asm.jnz(args[0], args[1]); err != nil {
			return errors.Wrapf(err, "could not complete instruction %v", instr)
		}
		return nil
	}

	return fmt.Errorf("unknown instruction %v", instr)
}

// Reset resets the registers and instruction pointer to zero.
func (asm *assembunnyComputer) Reset() {
	asm.registers = struct {
		a int
		b int
		c int
		d int
	}{0, 0, 0, 0}
	asm.SetInstructionPointer(0)
}

// ReadMemory reads from a tape of memory represented by a slice of instructions.
// It does not reset the registers or the instruction pointer.
func (asm *assembunnyComputer) ReadMemory(memory []string) error {
	for asm.GetInstructionPointer() < len(memory) && asm.GetInstructionPointer() >= 0 {
		instr := memory[asm.GetInstructionPointer()]
		if err := asm.readInstruction(instr); err != nil {
			return errors.Wrapf(err, "could not parse instruction %v at pointer %v", instr, asm.GetInstructionPointer())
		}
	}
	return nil
}

// Day12 solves the twelvth day puzzle "Leonardo's Monorail".
//
// Input
//
// A list of instructions written in assembunny.
// The number of instructions do not exceed 50.
// For example:
//
// 		cpy 41 a
// 		inc a
// 		inc a
// 		dec a
// 		jnz a 2
// 		dec a
func Day12(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	var asm assembunnyComputer
	memory := make([]string, 0)
	for scanner.Scan() {
		memory = append(memory, scanner.Text())
	}
	if err = asm.ReadMemory(memory); err != nil {
		err = errors.Wrapf(err, "could not read memory %v", memory)
		return
	}
	valA, err := asm.Get("a")
	if err != nil {
		err = errors.Wrapf(err, "could not get reg a")
		return
	}
	answer1 = strconv.Itoa(valA)

	// but something's up
	asm.Reset()
	if err = asm.Set("c", 1); err != nil {
		err = errors.Wrapf(err, "could not set reg a")
		return
	}
	if err = asm.ReadMemory(memory); err != nil {
		err = errors.Wrapf(err, "could not read memory %v", memory)
		return
	}
	if valA, err = asm.Get("a"); err != nil {
		err = errors.Wrapf(err, "could not get reg a")
		return
	}
	answer2 = strconv.Itoa(valA)
	return
}

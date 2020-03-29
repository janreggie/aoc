package aoc2016

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_assembunnyComputer_ReadMemory(t *testing.T) {
	assert := assert.New(t)
	mem := []string{"cpy 41 a",
		"inc a",
		"inc a",
		"dec a",
		"jnz a 2",
		"dec a",
	}
	var asm assembunnyComputer
	err := asm.ReadMemory(mem)
	assert.NoError(err, mem, err)
	assert.Equal(asm.registers.a, 42)
}

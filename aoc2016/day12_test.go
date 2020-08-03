package aoc2016

import (
	"testing"

	"github.com/janreggie/AdventOfCode/internal"
	"github.com/stretchr/testify/assert"
)

const day12myInput = `cpy 1 a
cpy 1 b
cpy 26 d
jnz c 2
jnz 1 5
cpy 7 c
inc d
dec c
jnz c -2
cpy a c
inc a
dec b
jnz b -2
cpy c b
dec d
jnz d -6
cpy 17 c
cpy 18 d
inc a
dec d
jnz d -2
dec c
jnz c -5
`

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

func TestDay12(t *testing.T) {
	assert := assert.New(t)
	testCase := internal.TestCase{
		Details: "Y2016D12 my input",
		Input:   day12myInput,
		Result1: `318117`,
		Result2: `9227771`,
	}
	testCase.Test(Day12, assert)
}

func BenchmarkDay12(b *testing.B) {
	internal.Benchmark(Day12, b, day12myInput)
}

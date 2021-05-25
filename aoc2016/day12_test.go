package aoc2016

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
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

func TestDay12(t *testing.T) {
	assert := assert.New(t)
	testCase := aoc.TestCase{
		Details: "Y2016D12 my input",
		Input:   day12myInput,
		Result1: `318117`,
		Result2: `9227771`,
	}
	testCase.Test(Day12, assert)
}

func BenchmarkDay12(b *testing.B) {
	aoc.Benchmark(Day12, b, day12myInput)
}

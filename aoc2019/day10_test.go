package aoc2019

import (
	"testing"

	"github.com/janreggie/AdventOfCode/structs"
	"github.com/stretchr/testify/assert"
)

func TestDay10(t *testing.T) {
	assert := assert.New(t)
	testCases := []structs.TestCase{
		{Input: ".#..##.###...#######\n" +
			"##.############..##.\n" +
			".#.######.########.#\n" +
			".###.#######.####.#.\n" +
			"#####.##.#.##.###.##\n" +
			"..#####..#.#########\n" +
			"####################\n" +
			"#.####....###.#.#.##\n" +
			"##.#################\n" +
			"#####.##.###..####..\n" +
			"..######..##.#######\n" +
			"####.##.####...##..#\n" +
			".#####..#.######.###\n" +
			"##...#.##########...\n" +
			"#.##########.#######\n" +
			".####.#.###.###.#.##\n" +
			"....##.##.###..#####\n" +
			".#.#.###########.###\n" +
			"#.#.#.#####.####.###\n" +
			"###.##.####.##.#..##\n",
			Result1: "210",
			Result2: "802"},
	}
	for _, eachCase := range testCases {
		eachCase.Test(Day10, assert)
	}
}

func Benchmark_generateSisonPoints(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		channel := generateSisonPoints(10, 15)
		for range channel {
			// do nothing
		}
	}
}
func Benchmark_generateSisonPointsSlow(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		channel := generateSisonPointsSlow(10, 15)
		for range channel {
			// do nothing
		}
	}
}

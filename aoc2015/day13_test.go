package aoc2015

import (
	"bufio"
	"strings"
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

const day13myInput = `Alice would gain 2 happiness units by sitting next to Bob.
Alice would gain 26 happiness units by sitting next to Carol.
Alice would lose 82 happiness units by sitting next to David.
Alice would lose 75 happiness units by sitting next to Eric.
Alice would gain 42 happiness units by sitting next to Frank.
Alice would gain 38 happiness units by sitting next to George.
Alice would gain 39 happiness units by sitting next to Mallory.
Bob would gain 40 happiness units by sitting next to Alice.
Bob would lose 61 happiness units by sitting next to Carol.
Bob would lose 15 happiness units by sitting next to David.
Bob would gain 63 happiness units by sitting next to Eric.
Bob would gain 41 happiness units by sitting next to Frank.
Bob would gain 30 happiness units by sitting next to George.
Bob would gain 87 happiness units by sitting next to Mallory.
Carol would lose 35 happiness units by sitting next to Alice.
Carol would lose 99 happiness units by sitting next to Bob.
Carol would lose 51 happiness units by sitting next to David.
Carol would gain 95 happiness units by sitting next to Eric.
Carol would gain 90 happiness units by sitting next to Frank.
Carol would lose 16 happiness units by sitting next to George.
Carol would gain 94 happiness units by sitting next to Mallory.
David would gain 36 happiness units by sitting next to Alice.
David would lose 18 happiness units by sitting next to Bob.
David would lose 65 happiness units by sitting next to Carol.
David would lose 18 happiness units by sitting next to Eric.
David would lose 22 happiness units by sitting next to Frank.
David would gain 2 happiness units by sitting next to George.
David would gain 42 happiness units by sitting next to Mallory.
Eric would lose 65 happiness units by sitting next to Alice.
Eric would gain 24 happiness units by sitting next to Bob.
Eric would gain 100 happiness units by sitting next to Carol.
Eric would gain 51 happiness units by sitting next to David.
Eric would gain 21 happiness units by sitting next to Frank.
Eric would gain 55 happiness units by sitting next to George.
Eric would lose 44 happiness units by sitting next to Mallory.
Frank would lose 48 happiness units by sitting next to Alice.
Frank would gain 91 happiness units by sitting next to Bob.
Frank would gain 8 happiness units by sitting next to Carol.
Frank would lose 66 happiness units by sitting next to David.
Frank would gain 97 happiness units by sitting next to Eric.
Frank would lose 9 happiness units by sitting next to George.
Frank would lose 92 happiness units by sitting next to Mallory.
George would lose 44 happiness units by sitting next to Alice.
George would lose 25 happiness units by sitting next to Bob.
George would gain 17 happiness units by sitting next to Carol.
George would gain 92 happiness units by sitting next to David.
George would lose 92 happiness units by sitting next to Eric.
George would gain 18 happiness units by sitting next to Frank.
George would gain 97 happiness units by sitting next to Mallory.
Mallory would gain 92 happiness units by sitting next to Alice.
Mallory would lose 96 happiness units by sitting next to Bob.
Mallory would lose 51 happiness units by sitting next to Carol.
Mallory would lose 81 happiness units by sitting next to David.
Mallory would gain 31 happiness units by sitting next to Eric.
Mallory would lose 73 happiness units by sitting next to Frank.
Mallory would lose 89 happiness units by sitting next to George.
`

const day13sampleInput = `Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.
`

func day13sampleScenario() *tableScenario {
	// guaranteed not to error
	scenario, _ := newTableScenario(bufio.NewScanner(strings.NewReader(day13sampleInput)))
	return scenario
}

func day13myScenario() *tableScenario {
	scenario, _ := newTableScenario(bufio.NewScanner(strings.NewReader(day13myInput)))
	return scenario
}

func Test_excludeVisitors(t *testing.T) {
	assert := assert.New(t)
	oldVisitors := []visitor{"Alice", "Bob", "Carol", "David", "Eric"}
	newVisitors := excludeVisitors(oldVisitors, "Carol")
	assert.ElementsMatch([]visitor{"Alice", "Bob", "Carol", "David", "Eric"}, oldVisitors)
	assert.ElementsMatch([]visitor{"Alice", "Bob", "David", "Eric"}, newVisitors)
}

func Test_newTableScenario(t *testing.T) {
	scenario := day13sampleScenario()
	assert := assert.New(t)
	assert.ElementsMatch([]visitor{"Alice", "Bob", "Carol", "David"},
		scenario.visitors)
	assert.Equal(map[visitorPair]happiness{
		{"Alice", "Bob"}:   137,
		{"Alice", "Carol"}: -141,
		{"Alice", "David"}: 44,
		{"Bob", "Carol"}:   53,
		{"Bob", "David"}:   -70,
		{"Carol", "David"}: 96,
	}, scenario.potentialHappiness)
}

func Test_tableScenario_happiestExhaustive(t *testing.T) {
	assert := assert.New(t)
	scenario := day13sampleScenario()
	assert.EqualValues(330, scenario.happiestExhaustive())
}

func Test_tableScenario_happiestPermutative(t *testing.T) {
	assert := assert.New(t)
	scenario := day13sampleScenario()
	assert.EqualValues(330, scenario.happiestPermutative())
}

func Test_seatingArrangement(t *testing.T) {
	assert := assert.New(t)
	scenario := day13myScenario()
	arrangement, err := newSeatingArrangement("Alice", "Bob", scenario)
	assert.NoError(err)
	assert.Equal(seatingArrangement{
		basis:     scenario,
		raw:       []visitor{"Alice", "Bob"},
		remaining: []visitor{"Carol", "David", "Eric", "Frank", "George", "Mallory"},
		happiness: 42,
	}, arrangement, "added Alice and Bob")

	arrangement, err = arrangement.add("David")
	assert.NoError(err)
	assert.Equal(seatingArrangement{
		basis:     scenario,
		raw:       []visitor{"Alice", "Bob", "David"},
		remaining: []visitor{"Carol", "Eric", "Frank", "George", "Mallory"},
		happiness: 9,
	}, arrangement, "added David")

	_, err = arrangement.add("Earl")
	assert.Error(err, "added nonexistent Earl")
	_, err = arrangement.add("David")
	assert.Error(err, "added already present David")

	arrangement, err = arrangement.add("Mallory")
	assert.NoError(err)
	assert.Equal(seatingArrangement{
		basis:     scenario,
		raw:       []visitor{"Alice", "Bob", "David", "Mallory"},
		remaining: []visitor{"Carol", "Eric", "Frank", "George"},
		happiness: -30,
	}, arrangement, "added Mallory")

	_, err = arrangement.add("Mallory")
	assert.Error(err, "added already present Mallory")

	arrangement, err = arrangement.add("George")
	assert.NoError(err)
	assert.Equal(seatingArrangement{
		basis:     scenario,
		raw:       []visitor{"Alice", "Bob", "David", "Mallory", "George"},
		remaining: []visitor{"Carol", "Eric", "Frank"},
		happiness: -22,
	}, arrangement, "added George")

	arrangement, err = arrangement.add("Eric")
	assert.NoError(err)
	assert.Equal(seatingArrangement{
		basis:     scenario,
		raw:       []visitor{"Alice", "Bob", "David", "Mallory", "George", "Eric"},
		remaining: []visitor{"Carol", "Frank"},
		happiness: -59,
	}, arrangement, "added Eric")

	_, err = arrangement.add("Eric")
	assert.Error(err, "added already present Eric")
	_, err = arrangement.add("Alice")
	assert.Error(err, "added already present Alice")
	_, err = arrangement.add("Bob")
	assert.Error(err, "added already present Bob")
	_, err = arrangement.add("David")
	assert.Error(err, "added already present David")
	_, err = arrangement.add("Mallory")
	assert.Error(err, "added already present Mallory")

	arrangement, err = arrangement.add("Carol")
	assert.NoError(err)
	assert.Equal(seatingArrangement{
		basis:     scenario,
		raw:       []visitor{"Alice", "Bob", "David", "Mallory", "George", "Eric", "Carol"},
		remaining: []visitor{"Frank"},
		happiness: 136,
	}, arrangement, "added Carol")

	arrangement, err = arrangement.add("Frank")
	assert.NoError(err)
	assert.Equal(seatingArrangement{
		basis:     scenario,
		raw:       []visitor{"Alice", "Bob", "David", "Mallory", "George", "Eric", "Carol", "Frank"},
		remaining: []visitor{},
		happiness: 228,
	}, arrangement, "added Frank and went full circle")

}

func Test_seatingArrangementQueue(t *testing.T) {
	assert := assert.New(t)
	scenario := day13sampleScenario()
	queue := newSeatingArrangementQueue()

	// suppose we start from Alice.
	// should not error!
	arrangement, err := newSeatingArrangement("Alice", "Alice", scenario)
	assert.Error(err)
	arrangement, err = newSeatingArrangement("Alice", "Bob", scenario)
	assert.NoError(err)
	queue.push(arrangement)
	arrangement, err = newSeatingArrangement("Alice", "Carol", scenario)
	assert.NoError(err)
	queue.push(arrangement)
	arrangement, err = newSeatingArrangement("Alice", "David", scenario)
	assert.NoError(err)
	queue.push(arrangement)
	assert.ElementsMatch([]seatingArrangement{
		{scenario, []visitor{"Alice", "Carol"}, []visitor{"Bob", "David"}, -141},
		{scenario, []visitor{"Alice", "David"}, []visitor{"Bob", "Carol"}, 44},
		{scenario, []visitor{"Alice", "Bob"}, []visitor{"Carol", "David"}, 137},
	}, queue)
}

func TestDay13(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Details: "Y2015D13 sample input",
			Input:   day13sampleInput,
			Result1: "330",
			Result2: "286"},
		{Details: "Y2015D13 my input",
			Input:   day13myInput,
			Result1: "733",
			Result2: "725"},
	}
	for _, tt := range testCases {
		tt.Test(Day13, assert)
	}
}

func Benchmark_tableScenario_happiestExhaustive(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		day13myScenario().happiestExhaustive()
	}
}

func Benchmark_tableScenario_happiestPermutative(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		day13myScenario().happiestPermutative()
	}
}

func BenchmarkDay13(b *testing.B) {
	internal.Benchmark(Day13, b, day13myInput)
}

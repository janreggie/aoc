package aoc2015

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func day13sampleScenario() *tableScenario {
	// guaranteed not to error
	scenario, _ := newTableScenario(day13sampleInput)
	return scenario
}

func day13myScenario() *tableScenario {
	scenario, _ := newTableScenario(day13myInput)
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
	testCases := []aoc.TestCase{
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
	aoc.Benchmark(Day13, b, day13myInput)
}

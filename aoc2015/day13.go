package aoc2015

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

// "The endgame of politics is to make everyone sad.
// Not to make some happy and some miserable,
// but to make everyone sad." - probably someone.

// happiness is represented by an int.
type happiness int

func (happiness happiness) int() int {
	return int(happiness)
}

// visitor is represented by a string
type visitor string

func (visitor visitor) String() string {
	return string(visitor)
}

// excludeVisitors returns the same slice of visitors but without some specified values.
func excludeVisitors(from []visitor, except ...visitor) []visitor {
	out := make([]visitor, 0, len(from)/2) // don't worry it'll realloc.
FORALL:
	for _, vv := range from {
		// check if vv equals any of except
		for _, exception := range except {
			if vv == exception {
				continue FORALL
			}
		}
		out = append(out, vv)
	}
	return out
}

// visitorPair is a seating arrangement of two visitors
// where a is seated next to b.
// Naturally b would be seated next to a as well.
// The happiness of a visitorPair is the happiness that a would gain from b
// plus the happiness that b would gain from a.
// a is lexographically less than b.
type visitorPair struct {
	a, b visitor
}

func (pair visitorPair) String() string {
	return fmt.Sprintf("(%v,%v)", pair.a, pair.b)
}

// tableScenario is a representation of a "scenario"
// which is represented by a list of visitorPairs and their respective happiness gains.
type tableScenario struct {
	// potentialHappiness is the potential happiness gain of a visitor pair
	potentialHappiness map[visitorPair]happiness

	// visitors are all the visitors in the scenario
	visitors []visitor
}

// newTableScenario creates a table scenario from a scanner object.
// It uses parseTableScenario to read text such as:
//
//     Alice would lose 2 happiness units by sitting next to David.
func newTableScenario(input string) (*tableScenario, error) {
	potentialHappiness := make(map[visitorPair]happiness)
	visitors := make([]visitor, 0)
	result := &tableScenario{potentialHappiness: potentialHappiness, visitors: visitors}

	// now parse each line
	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}
		if err := result.parseTableScenario(line); err != nil {
			return result, errors.Wrapf(err, "could not parse line %v", line)
		}
	}

	return result, nil
}

func (scenario *tableScenario) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Visitors: %v\n", scenario.visitors))
	sb.WriteString("Happiness:\n")
	for kk, vv := range scenario.potentialHappiness {
		sb.WriteString(fmt.Sprintf("%v:%v\n", kk, vv))
	}

	return sb.String()
}

func (scenario *tableScenario) copy() tableScenario {
	result := tableScenario{}
	result.potentialHappiness = make(map[visitorPair]happiness)
	for kk, vv := range scenario.potentialHappiness {
		result.potentialHappiness[kk] = vv
	}
	result.visitors = make([]visitor, len(scenario.visitors))
	copy(result.visitors, scenario.visitors)
	return result
}

// parseTableScenario parses each line of Day 13's input
// and modifies the scenario in-place accordingly.
// A line should look like the following two:
//
//	Alice would gain 26 happiness units by sitting next to Carol.
//	Alice would lose 82 happiness units by sitting next to David.
//
// The names are a combination of letters, and they would either gain or lose
// some amount of happiness by sitting next to someone who is not them.
// The line ends with a period.
func (scenario *tableScenario) parseTableScenario(line string) error {
	// check if line ends in a period
	if line[len(line)-1] != '.' {
		return errors.Errorf("could not parse %v: does not end with period", line)
	}
	line = line[:len(line)-1]

	// use fields
	splitLine := strings.Fields(line)
	if len(splitLine) != 11 {
		return errors.Errorf("could not parse %v: wrong number of elements", line)
	}

	vOne, vTwo := splitLine[0], splitLine[10]
	if vOne == vTwo {
		return errors.Errorf("%v is seated with himself!: %v", vOne, line)
	}
	if vOne > vTwo {
		vOne, vTwo = vTwo, vOne
	}

	gain, err := strconv.Atoi(splitLine[3])
	if err != nil {
		return errors.Wrapf(err, "%v cannot be parsed", splitLine[3])
	}
	if splitLine[2] == "lose" {
		gain = -gain
	}

	return scenario.influence(visitorPair{visitor(vOne), visitor(vTwo)}, happiness(gain))
}

// addToVisitors will add a pair's names to the visitor list
func (scenario *tableScenario) addToVisitors(pair visitorPair) {
	writeA, writeB := true, true
	for _, vv := range scenario.visitors {
		if pair.a == vv {
			writeA = false
		}
		if pair.b == vv {
			writeB = false
		}
	}

	if writeA {
		scenario.visitors = append(scenario.visitors, pair.a)
	}
	if writeB {
		scenario.visitors = append(scenario.visitors, pair.b)
	}
}

// influence adds happiness to the happiness that visitorPair would have.
// Will give off an error if pair refers to only one person
// or if one of pair is an empty string.
func (scenario *tableScenario) influence(pair visitorPair, h happiness) error {
	if pair.a == "" || pair.b == "" {
		return errors.New("influence: empty name")
	}
	if pair.a > pair.b {
		pair.a, pair.b = pair.b, pair.a
	}

	scenario.potentialHappiness[pair] += h
	scenario.addToVisitors(pair)

	return nil
}

// permutations creates a channel that contains all the possible circular seating arrangements
// and will close said channel.
func (scenario *tableScenario) permutations() <-chan []visitor {
	permutate := func(c chan []visitor, inputs []visitor) {
		output := make([]visitor, len(inputs))
		copy(output, inputs)
		c <- output

		size := len(inputs)
		p := make([]int, size+1)
		for i := 0; i < size+1; i++ {
			p[i] = i
		}
		for i := 1; i < size; {
			p[i]--
			j := 0
			if i%2 == 1 {
				j = p[i]
			}
			tmp := inputs[j]
			inputs[j] = inputs[i]
			inputs[i] = tmp
			output := make([]visitor, len(inputs))
			copy(output, inputs)
			c <- output
			for i = 1; p[i] == 0; i++ {
				p[i] = i
			}
		}
	}

	c := make(chan []visitor)
	go func(c chan []visitor) {
		defer close(c)
		visitors := make([]visitor, len(scenario.visitors))
		copy(visitors, scenario.visitors)
		permutate(c, visitors)
	}(c)
	return c
}

// get determines the happiness that two people seated together have.
// Return an error if the two visitors aren't in the visitor list
// or if the visitors are equal to each other.
// If the happiness between these visitors does not exist return zero
// as if they are neutral towards each other.
func (scenario *tableScenario) get(a, b visitor) (happiness, error) {
	if a == b {
		return 0, fmt.Errorf("could not evaluate happiness of %v and themselves", a)
	}
	if a > b {
		a, b = b, a // reverse them
	}

	aIn, bIn := false, false
	for _, vv := range scenario.visitors {
		if a == vv {
			aIn = true
		}
		if b == vv {
			bIn = true
		}
		if aIn && bIn {
			break
		}
	}
	if !aIn {
		return 0, errors.Errorf("%v not in scenario", a)
	}
	if !bIn {
		return 0, errors.Errorf("%v not in scenario", b)
	}

	return scenario.potentialHappiness[visitorPair{a, b}], nil
}

// happiestFrom determines the happiest seating arrangement starting from some person.
// This uses the concept of seatingArrangements and a seatingArrangementQueue to record all arrangements.
func (scenario *tableScenario) happiestFrom(visitor visitor) happiness {
	queue := newSeatingArrangementQueue()

	// check all visitors in the list
	for _, vv := range scenario.visitors {
		arrangement, err := newSeatingArrangement(visitor, vv, scenario)
		if err != nil {
			continue // maybe same person
		}
		queue.push(arrangement)
	}

	var record happiness
	for len(queue) > 0 {
		arrangement, _ := queue.pop() // guaranteed no error
		if len(arrangement.remaining) == 0 {
			// exhaust!
			if arrangement.happiness > record {
				record = arrangement.happiness
			}
		}

		// otherwise consider all remaining in arrangement
		for _, next := range arrangement.remaining {
			nextArr, err := arrangement.add(next)
			if err != nil {
				continue // next would exist already in the arrangement
			}
			queue.push(nextArr)
		}
	}

	return record
}

// happiestExhaustive determines the happiest seating arrangement
// by exhausting all visitors in scenario.
// This simply gets the happiest arrangement from the first visitor in the scenario
// as the list is circular and it doesn't really matter who the first visitor is.
func (scenario *tableScenario) happiestExhaustive() happiness {
	return scenario.happiestFrom(scenario.visitors[0])
}

// happiestPermutative determines the happiest arrangement
// by checking all possible circular seating arrangements.
func (scenario *tableScenario) happiestPermutative() happiness {
	arrangements := scenario.permutations()
	var record happiness

	for arrangement := range arrangements {
		var thisHappiness happiness
		for ii := range arrangement {
			// the modulo does the head-tail for us
			iiH, _ := scenario.get(arrangement[ii], arrangement[(ii+1)%len(arrangement)])
			thisHappiness += iiH
		}
		if thisHappiness > record {
			record = thisHappiness
		}
	}
	return record
}

// seatingArrangement is a seating arrangement around a table in Day 13.
type seatingArrangement struct {
	basis     *tableScenario
	raw       []visitor
	remaining []visitor
	happiness happiness
}

// newSeatingArrangement creates a new seating arrangement that starts with a and then b.
// Will return an error if the two visitors are the same,
// or if the pair's names could not be found in the visitor list.
func newSeatingArrangement(a, b visitor, scenario *tableScenario) (seatingArrangement, error) {
	// do things...
	result := seatingArrangement{}
	result.basis = scenario
	var err error
	if result.happiness, err = result.basis.get(a, b); err != nil {
		return result, errors.Wrapf(err, "could not arrange %v and %v together", a, b)
	}
	result.raw = []visitor{a, b}
	result.remaining = excludeVisitors(result.basis.visitors, a, b)

	return result, nil
}

func (arrangement seatingArrangement) String() string {
	var sb strings.Builder
	sb.WriteString("Visitors: ")
	for _, vv := range arrangement.raw {
		sb.WriteString(vv.String())
		sb.WriteString(", ")
	}
	sb.WriteString(fmt.Sprintf("; Happiness: %v", arrangement.happiness))
	return sb.String()
}

// copy creates a copy of the seating arrangement
func (arrangement seatingArrangement) copy() seatingArrangement {
	result := seatingArrangement{}
	result.basis = arrangement.basis
	result.raw = make([]visitor, len(arrangement.raw))
	copy(result.raw, arrangement.raw)
	result.remaining = make([]visitor, len(arrangement.remaining))
	copy(result.remaining, arrangement.remaining)
	result.happiness = arrangement.happiness
	return result
}

// add adds a visitor to a seating arrangement
// and returns a new seating arrangement with the next visitor
// and an error if the visitor is already in the arrangement or is not in remaining.
// If the visitor is the last in remaining, happiness between the last visitor
// and the first one will be added.
func (arrangement seatingArrangement) add(next visitor) (seatingArrangement, error) {
	result := arrangement.copy()

	// check if visitor is in remaining
	notInRemaining := true
	for _, vv := range arrangement.remaining {
		if vv == next {
			notInRemaining = false
		}
	}
	if notInRemaining {
		return result, errors.Errorf("%v not in remaining", next)
	}

	result.raw = append(result.raw, next)
	result.remaining = excludeVisitors(result.remaining, next)
	// following should not return error
	nextHappiness, _ := result.basis.get(arrangement.raw[len(arrangement.raw)-1], next)
	result.happiness += nextHappiness

	// check if remaining is empty
	if len(result.remaining) == 0 {
		headTailHappiness, _ := result.basis.get(result.raw[0], next)
		result.happiness += headTailHappiness
	}

	return result, nil
}

// seatingArrangementQueue is a queue of seatingArrangements.
// The values in the queue are arranged by their happiness,
// where the leading element is the happiest scenario,
// used to perform a breadth-first-search throughout the various seating arrangements
// that can be performed in a scenario.
type seatingArrangementQueue []seatingArrangement

// newSeatingArrangementQueue creates a seatingArrangementQueue construct.
// As the type's "zero value" is a nil slice, that could not be used as-is.
func newSeatingArrangementQueue() seatingArrangementQueue {
	return make(seatingArrangementQueue, 0)
}

// push pushes a seatingArrangement to the queue.
func (queue *seatingArrangementQueue) push(arrangement seatingArrangement) {
	if len(*queue) == 0 {
		*queue = append(*queue, arrangement.copy())
		return
	}

	ind := 0
	for ind < len(*queue) && (*queue)[ind].happiness < arrangement.happiness {
		ind++
	}
	// append a value to queue. Just any value..
	*queue = append(*queue, seatingArrangement{})
	// move all elements...
	for ii := len(*queue) - 1; ii > ind; ii-- {
		(*queue)[ii] = (*queue)[ii-1]
	}
	(*queue)[ind] = arrangement
}

// pop pops a value from the queue.
// Will return an error if the queue is empty.
func (queue *seatingArrangementQueue) pop() (seatingArrangement, error) {
	if len(*queue) == 0 {
		return seatingArrangement{}, errors.New("queue is empty")
	}
	out := (*queue)[0]
	*queue = (*queue)[1:]
	return out, nil
}

// Day13 solves the thirteenth day puzzle "Knights of the Dinner Table".
//
// Input
//
// A file describing the happiness loss or gain between two visitors
// in a group of eight visitors. For example:
//
//	Alice would gain 2 happiness units by sitting next to Bob.
//	Alice would gain 26 happiness units by sitting next to Carol.
//	Alice would lose 82 happiness units by sitting next to David.
//	Alice would lose 75 happiness units by sitting next to Eric.
//	Alice would gain 42 happiness units by sitting next to Frank.
//	Alice would gain 38 happiness units by sitting next to George.
//	Alice would gain 39 happiness units by sitting next to Mallory.
//	Bob would gain 40 happiness units by sitting next to Alice.
//	Bob would lose 61 happiness units by sitting next to Carol.
//	Bob would lose 15 happiness units by sitting next to David.
//	(this line continues)
//
// If the happiness between one person and another is not provided
// it is assumed that the former would gain zero happiness
// units by sitting next to the latter.
//
// It is guaranteed that the gain or loss of happiness between any two
// people is no more than 100.
func Day13(input string) (answer1, answer2 string, err error) {

	scenario, err := newTableScenario(input)
	if err != nil {
		err = errors.Wrap(err, "could not parse input")
		return
	}

	// use waitGroup
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		answer1 = strconv.Itoa(scenario.happiestPermutative().int())
		wg.Done()
	}()

	go func() {
		// now how about yourself?
		// This should have no error
		scenarioWithYou := scenario.copy()
		scenarioWithYou.addToVisitors(visitorPair{"Alice", "You"})
		answer2 = strconv.Itoa(scenarioWithYou.happiestPermutative().int())
		wg.Done()
	}()

	wg.Wait()
	return
}

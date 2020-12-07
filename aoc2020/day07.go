package aoc2020

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// bag describes a bag (its quality and color)
type bag struct {
	quality string
	color   string
}

// bagRule describes how many bags of some quality and color can exist
type bagRule struct {
	count int
	bag   bag
}

// bagRuleset is a hashmap which describes the contents and bags some bag can have
type bagRuleset struct {
	// rr denotes representation
	rr map[bag][]bagRule
}

// generateBagRuleset generates a bagRuleset from a string containing rules
func generateBagRuleset(rules *bufio.Scanner) (*bagRuleset, error) {
	ruleset := &bagRuleset{
		rr: make(map[bag][]bagRule),
	}

	for rules.Scan() {
		rule := rules.Text()
		if err := ruleset.addRule(rule); err != nil {
			return nil, errors.Wrapf(err, "could not parse rule %s", rule)
		}
	}
	return ruleset, nil
}

// addRule adds a rule to bagRuleset. Returns an error if the string cannot be parsed.
// For instance, rule can be "light red bags contain 1 bright white bag, 2 muted yellow bags.".
func (ruleset *bagRuleset) addRule(rule string) error {
	splitRaw := strings.Split(rule, " bags contain ")
	if len(splitRaw) != 2 {
		return errors.Errorf("couldn't split `%s` using ` bags contain `", rule)
	}

	// rawBag should be []string{desc, color}
	rawBag := strings.Split(splitRaw[0], " ")
	if len(rawBag) != 2 {
		return errors.Errorf("couldn't split `%s` to two", splitRaw[0])
	}
	readBag := bag{quality: rawBag[0], color: rawBag[1]}

	contents := make([]bagRule, 0)
	// rawContents should be []string{"5 faded blue bags", "6 dotted blue bags"} or whatever
	if splitRaw[1] != "no other bags." {
		rawContents := strings.Split(splitRaw[1][:len(splitRaw[1])-1], ", ")
		for _, content := range rawContents {
			var cnt int
			var qty, clr string
			// content should match format in Sscanf
			_, err := fmt.Sscanf(content, "%d %s %s bag", &cnt, &qty, &clr)
			if err != nil {
				return errors.Wrapf(err, "could not parse content `%s`", content)
			}
			contents = append(contents, bagRule{cnt, bag{qty, clr}})
		}
	}

	ruleset.rr[readBag] = contents
	return nil
}

// findBagRules returns a list of bagRules which pertain to a particular bag.
// If findBagRules cannot be found return an empty list.
func (ruleset *bagRuleset) findBagRules(bag bag) []bagRule {
	rules, found := ruleset.rr[bag]
	if !found {
		return []bagRule{}
	}
	result := make([]bagRule, len(rules))
	copy(result, rules)
	return result
}

// findContaining returns a list of bags in which some bag can be found (directly)
func (ruleset *bagRuleset) findContaining(child bag) []bag {
	result := make([]bag, 0)
	for bb, rs := range ruleset.rr {
		for _, rr := range rs {
			if child == rr.bag {
				result = append(result, bb)
			}
		}
	}
	return result
}

// findAllContaining returns a list of bags in which some bag can be found (indirectly)
func (ruleset *bagRuleset) findAllContaining(child bag) []bag {
	primary := make(map[bag]struct{})
	for _, found := range ruleset.findContaining(child) {
		primary[found] = struct{}{}
	}
	secondary := make(map[bag]struct{})

	for {
		for bb := range primary {
			secondary[bb] = struct{}{}
			for _, found := range ruleset.findContaining(bb) {
				secondary[found] = struct{}{}
			}
		}

		// check if there are any new elements in secondary that aren't in primary
		if len(primary) == len(secondary) {
			break
		}

		// and if there are, reassign primary and secondary
		primary = secondary
		secondary = make(map[bag]struct{})
	}

	result := make([]bag, 0)
	for bb := range primary {
		result = append(result, bb)
	}
	return result
}

// countBags counts how many bags are in parent recursively.
// If such a bag is "terminal", return 0.
func (ruleset *bagRuleset) countBags(parent bag) int {
	rules := ruleset.findBagRules(parent)
	if len(rules) == 0 {
		return 0
	}

	result := 0
	for _, rr := range rules {
		result += rr.count * (1 + ruleset.countBags(rr.bag))
	}
	return result
}

// Day07 solves the seventh day puzzle "Handy Haversacks"
//
// Input
//
// A list of entries describing the contents of bags with a particular color and quality. For example:
//
//   light red bags contain 1 bright white bag, 2 muted yellow bags.
//   dark orange bags contain 3 bright white bags, 4 muted yellow bags.
//   bright white bags contain 1 shiny gold bag.
//   muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
//   shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
//   dark olive bags contain 3 faded blue bags, 4 dotted black bags.
//   vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
//   faded blue bags contain no other bags.
//   dotted black bags contain no other bags.
//
// The input format should be standard:
//
//   QUALITY COLOR bags contain INTERNAL_1, INTERNAL_2, ..., INTERNAL_N.
//   INTERNAL_k = no other bags | 1 QUALITY COLOR bag | i QUALITY COLOR bags
//
// There should be no loops (i.e., Bag A contains B, which contains C, which contains A, which...).
func Day07(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	ruleset, err := generateBagRuleset(scanner)
	if err != nil {
		err = errors.Wrapf(err, "could not read from scanner")
		return
	}

	answer1 = strconv.Itoa(len(ruleset.findAllContaining(bag{"shiny", "gold"})))
	answer2 = strconv.Itoa(ruleset.countBags(bag{"shiny", "gold"}))
	return
}

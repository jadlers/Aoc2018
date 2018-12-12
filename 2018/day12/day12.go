package main

import (
	"fmt"
	util "github.com/jadlers/advent-of-code/util"
	"strings"
)

func main() {
	lines := util.ReadLines()
	initialState := strings.Fields(lines[0])[2]
	rules := lines[2:]
	p1, p2 := Day12(initialState, rules)

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

type Rule struct {
	pattern    string
	givesPlant bool
}

func Day12(initialState string, rulesInput []string) (p1, p2 int) {
	const genLimit = 20
	plants := "..." + initialState

	// Find all rules which gives plants
	rules := []Rule{}
	for _, rule := range rulesInput {
		fields := strings.Fields(rule)
		pattern := fields[0]
		if fields[2] == "#" {
			rules = append(rules, Rule{pattern, true})
		}
	}

	// fmt.Printf("[%v] \t%v\n", 0, plants)
	for generation := 1; generation <= genLimit; generation++ {
		plants = EvolveOneGeneration(plants, rules)
		// fmt.Printf("[%v] \t%v\n", generation, plants)
	}

	// Calculate value
	for i := range plants {
		if plants[i] == '#' {
			p1 += (i - 3)
		}
	}

	return
}

func EvolveOneGeneration(plants string, rules []Rule) string {
	willBePlants := map[int]bool{}

	dots := "....."
	for i := -2; i < len(plants); i++ {
		// go over all rules to see if any one match
		var currentPart string
		if i < 0 {
			currentPart = dots[:util.AbsInt(i)] + plants[:5+i]
		} else if i+4 >= len(plants) {
			currentPart = plants[i:] + dots[:i-len(plants)+5]
		} else {
			currentPart = plants[i : i+5]
		}
		for _, rule := range rules {
			if currentPart == rule.pattern {
				willBePlants[i+2] = true
			}
		}
	}

	// Update plants
	var b strings.Builder
	i := 0
	for len(willBePlants) > 0 {
		if willBePlants[i] {
			b.WriteRune('#')
			delete(willBePlants, i)
		} else {
			b.WriteRune('.')
		}
		i++
	}

	return b.String()
}

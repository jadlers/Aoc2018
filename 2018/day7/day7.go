package main

import (
	"fmt"
	util "github.com/jadlers/advent-of-code/util"
	"sort"
	"strings"
)

func main() {
	input := util.ReadLines()
	p1, p2 := Day7(input)

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

type Instruction struct {
	name   string
	before string
}

type Block struct {
	blocking string
	node     string
	active   bool
}

type BlockSlice []Block

func Day7(lines []string) (p1 string, p2 int) {
	allNames := map[string]bool{}
	blocking := BlockSlice{}
	for _, line := range lines {
		// Step Y must be finished before step D can begin.
		words := strings.Fields(line)
		allNames[words[1]] = true
		allNames[words[7]] = true
		blocking = append(blocking, Block{words[1], words[7], true})
	}

	for {
		chosen := nextStep(allNames, blocking)

		// Find all blocking by chosen
		for i, b := range blocking {
			if chosen == b.blocking {
				blocking[i] = Block{b.blocking, b.node, false}
			}

		}
		delete(allNames, chosen)

		p1 = p1 + chosen // Add to res
		if len(allNames) == 0 {
			break
		}
	}

	return
}

func nextStep(names map[string]bool, blocking BlockSlice) string {
	// Find next step
	candidates := []string{}
	for name, _ := range names {
		if !blocking.isBlocked(name) {
			candidates = append(candidates, name)
		}
	}

	return findFirst(candidates)
}

func (blocks BlockSlice) isBlocked(node string) bool {
	for _, b := range blocks {
		if b.node == node && b.active {
			return true
		}
	}
	return false
}

func findFirst(sli []string) string {
	sort.Slice(sli, func(i, j int) bool {
		return sli[i] < sli[j]
	})
	return sli[0]
}

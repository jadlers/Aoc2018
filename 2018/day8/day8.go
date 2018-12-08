package main

import (
	"fmt"
	util "github.com/jadlers/advent-of-code/util"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadLines()
	p1, p2 := Day8(input[0])

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

func Day8(line string) (p1, p2 int) {
	parts := []int{}
	for _, part := range strings.Fields(line) {
		i, _ := strconv.Atoi(part)
		parts = append(parts, i)
	}

	_, totalMetaData := nodeMetadata(parts, 0)
	p1 = totalMetaData

	return
}

func nodeMetadata(parts []int, total int) ([]int, int) {
	numChildren := parts[0]
	numMetadata := parts[1]

	parts = parts[2:] // Since we read numChild and numMetadata
	for i := 0; i < numChildren; i++ {
		parts, total = nodeMetadata(parts, total)
	}

	for i := 0; i < numMetadata; i++ {
		total += parts[i]
	}

	return parts[numMetadata:], total
}

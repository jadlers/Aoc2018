package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/jadlers/advent-of-code/util"
)

// Part 2: 2505
func main() {
	originalInput := readInput()
	// p1, p2 := Day1(input)

	fmt.Printf("Part 1: %v\n", p1(originalInput))
	// fmt.Printf("Part 2: %2d%02d\n", p2[0], p2[1])
}

// Part 1: 9581917
func p1(originalInput []int) int {
	var input = make([]int, len(originalInput))
	copy(input, originalInput)
	input[1], input[2] = 12, 2
	m := machine{input, 0}
	m.Run()

	return m.memory[0]
}

func readInput() []int {
	lines := util.ReadLines()
	input := []int{}
	for _, line := range lines {
		strVals := strings.Split(line, ",")
		for _, num := range strVals {
			val, _ := strconv.Atoi(num)
			input = append(input, val)
		}
	}

	return input
}

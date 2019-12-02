package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/jadlers/advent-of-code/util"
)

func main() {
	originalInput := readInput()

	fmt.Printf("Part 1: %v\n", p1(originalInput))
	fmt.Printf("Part 2: %v\n", p2(originalInput))
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

// Part 2: 2505
func p2(originalInput []int) int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			var input = make([]int, len(originalInput))
			copy(input, originalInput)

			input[1], input[2] = noun, verb
			m := machine{input, 0}
			ok := m.Run()

			if ok && m.memory[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}

	return -1
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

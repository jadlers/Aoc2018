package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/jadlers/advent-of-code/util"
)

// Part 1: 9581917
// Part 2: 2505
func main() {
	lines := util.ReadLines()
	input := []int{}
	for _, line := range lines {
		strVals := strings.Split(line, ",")
		for _, num := range strVals {
			val, _ := strconv.Atoi(num)
			input = append(input, val)
		}

	}
	p1, p2 := Day1(input)

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %2d%02d\n", p2[0], p2[1])
}

func Day1(input []int) (fstAns int, sndAns []int) {
	initialInput := make([]int, len(input))
	copy(initialInput, input)

	input[1] = 12
	input[2] = 2
	afterRun, _ := runProgram(input)
	fstAns = afterRun[0]

	// Loop over all possible combinations of a tuple between 0, 99 (inclusive)
	// Find the combination which gives: 19690720
	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(input, initialInput)
			input[1] = noun
			input[2] = verb

			newMemory, ok := runProgram(input)
			if ok == nil && newMemory[0] == 19690720 {
				sndAns = []int{noun, verb}
				return
			}
		}
	}

	return
}

func runProgram(input []int) ([]int, error) {
	opCode := -1
	instruction := 0
	for {
		position := instruction * 4
		opCode = input[position]
		if opCode == 99 {
			return input, nil
		}

		if !safeReadMem(input[position+1:position+4], len(input)) {
			return input, fmt.Errorf("Invalid read")
		}

		fst := input[position+1]
		snd := input[position+2]
		target := input[position+3]

		switch opCode {
		case 1: // Add
			input[target] = input[fst] + input[snd]
		case 2: // Multiply
			input[target] = input[fst] * input[snd]
		}
		instruction++
	}
}

func safeReadMem(part []int, memoryLength int) bool {
	for _, position := range part {
		if position < 0 || position >= memoryLength {
			return false
		}
	}
	return true
}

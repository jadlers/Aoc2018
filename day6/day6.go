package main

import (
	"fmt"
	util "github.com/jadlers/advent-of-code/util"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadLines()
	p1, p2 := Day6(input[0])

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

func Day6(input string) (p1, p2 int) {
	banksStr := strings.Fields(input)
	banks := []int{}
	for _, char := range banksStr {
		i, _ := strconv.Atoi(char)
		banks = append(banks, i)
	}

	seen := map[string]bool{fmt.Sprintf("%v", banks): true}
	for {
		nextBlocks := Redistribute(banks)
		key := fmt.Sprintf("%v", nextBlocks)
		// fmt.Printf("key: %v, seen: %v\n", key, seen)
		if seen[key] {
			p1++
			break
		}
		seen[key] = true
		p1++
	}

	p2 = 0
	return
}

func Redistribute(banks []int) []int {
	maxVal, maxIdx := banks[0], 0
	for idx, val := range banks {
		if val > maxVal {
			maxVal = val
			maxIdx = idx
		}
	}
	// Reset the chosen one
	banks[maxIdx] = 0
	for j := maxVal; j > 0; j-- {
		maxIdx += 1
		currentIdx := maxIdx % len(banks)
		banks[currentIdx] += 1
	}

	return banks
}

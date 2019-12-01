package main

import (
	"fmt"
	"strconv"

	util "github.com/jadlers/advent-of-code/util"
)

func main() {
	lines := util.ReadLines()
	p1, p2 := Day1(lines)

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

func calcFuel(weight int) int {
	return (weight / 3) - 2
}

func Day1(inputLines []string) (fstAns, sndAns int) {
	fstAns = p1(inputLines)
	sndAns = p2(inputLines)
	return
}

func p1(inputLines []string) int {
	total := 0
	for _, line := range inputLines {
		val, _ := strconv.Atoi(line)
		total += calcFuel(val)
	}
	return total
}

func calcFuel2(weight int) int {
	res := 0
	for weight > 0 {
		weight = calcFuel(weight)
		if weight > 0 {
			res += weight
		}
	}

	return res
}

// Guess 5169065 is too low
func p2(inputLines []string) int {
	total := 0
	for _, line := range inputLines {
		val, _ := strconv.Atoi(line)
		total += calcFuel2(val)
	}
	return total
}

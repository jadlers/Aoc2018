package main

import (
	"fmt"
	util "github.com/jadlers/AoC2018/util"
	"strings"
)

func main() {
	lines := util.ReadLines()

	fmt.Printf("Part 1: %v\n", Part1(lines))
	// fmt.Printf("Part 2: %v\n", Part2())
}

func Part1(lines []string) int {
	var res [2]int // index 0: Twos found, 1: Threes found
	for _, line := range lines {
		chars := strings.Split(line, "")

		letterOccurances := make(map[string]int)
		for _, char := range chars {
			letterOccurances[char] += 1
		}

		foundTwo, foundThree := false, false
		for _, val := range letterOccurances {
			if !foundTwo && val == 2 {
				foundTwo = true
				res[0] += 1
			} else if !foundThree && val == 3 {
				foundThree = true
				res[1] += 1
			}
		}
	}

	checksum := res[0] * res[1]
	return checksum
}

// func Part2() int {
// 	return 0
// }

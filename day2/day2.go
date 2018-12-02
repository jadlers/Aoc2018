package main

import (
	"errors"
	"fmt"
	util "github.com/jadlers/AoC2018/util"
	"strings"
)

func main() {
	lines := util.ReadLines()

	fmt.Printf("Part 1: %v\n", Part1(lines))
	fmt.Printf("Part 2: %v\n", Part2(lines))
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

func Part2(lines []string) string {
	var chars [][]string
	for _, line := range lines {
		chars = append(chars, strings.Split(line, ""))
	}

	for i, line := range chars {
		for j := i + 1; j < len(chars); j++ {
			idx, err := findDifferingChars(line, chars[j])

			if err == nil { // We've found the two sequences
				sndPart := strings.Join(line[idx+1:], "")
				res := append(line[:idx], sndPart)
				return strings.Join(res, "")
			}
		}
	}
	return ""
}

func findDifferingChars(fst, snd []string) (index int, err error) {
	index = -1
	err = nil
	for i, c1 := range fst {
		c2 := snd[i]
		if c1 != c2 {
			if index == -1 {
				index = i
			} else {
				err = errors.New("More than one matching char")
				return
			}
		}
	}

	if index == -1 {
		err = errors.New("No chars where matching")
	}

	return
}

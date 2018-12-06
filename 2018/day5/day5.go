package main

import (
	"fmt"
	util "github.com/jadlers/advent-of-code/util"
	"strings"
)

func main() {
	input := util.ReadLines()
	p1, p2 := Day5(input[0])

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

func Day5(input string) (p1, p2 int) {
	res := fullyReact(input)
	p1 = len(res)

	// Part 2
	alphabet := strings.Split("abcdefghijklmnopqrstuvxyz", "")

	finalLengths := map[string]int{}
	for _, letter := range alphabet {
		r := strings.NewReplacer(letter, "", strings.ToUpper(letter), "")
		curr := r.Replace(input)
		if len(curr) < len(input) {
			currReacted := fullyReact(curr)
			finalLengths[letter] = len(currReacted)
		}
	}

	p2 = len(input)
	for _, val := range finalLengths {
		if val < p2 {
			p2 = val
		}
	}

	return
}

func fullyReact(str string) string {
	var removed int
	for {
		str, removed = removeReacting(str)
		if removed == 0 {
			return str
		}
	}
}

func removeReacting(polymer string) (remainder string, removed int) {
	chars := strings.Split(polymer, "")
	for i, char := range chars {
		if i+1 == len(polymer) {
			break
		}

		if IsReacting(char, chars[i+1]) {
			chars[i] = " "
			chars[i+1] = " "
			removed += 2
		}
	}
	remainder = strings.Join(chars, "")
	remainder = util.TrimAllWhitespace(remainder)

	return
}

func IsReacting(a, b string) bool {
	if strings.ToUpper(a) == strings.ToUpper(b) {
		if a != b {
			return true
		}
	}
	return false
}

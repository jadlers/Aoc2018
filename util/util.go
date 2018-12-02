package util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLines() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func StringSliceToInt(from []string) []int {
	to := make([]int, len(from))
	for i, val := range from {
		integer, _ := strconv.Atoi(val)
		to[i] = integer
	}
	return to
}

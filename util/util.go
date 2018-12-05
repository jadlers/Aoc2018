package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func TrimAllWhitespace(str string) string {
	// Taken from: https://stackoverflow.com/a/32081891
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

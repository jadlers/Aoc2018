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

func ReadLinesFromFile(path string) []string {
	if file, err := os.Open(path); err == nil {
		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		return lines
	} else {
		panic(err)
	}
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

func AbsInt(no int) int {
	if no < 0 {
		return -no
	}
	return no
}

func IncludesInt(val int, slice []int) bool {
	for _, cur := range slice {
		if val == cur {
			return true
		}
	}
	return false
}

package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  int
	}{{"Example", []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.lines); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  string
	}{{"Example", []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}, "fgij"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.lines); got != tt.want {
				t.Errorf("Part2() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

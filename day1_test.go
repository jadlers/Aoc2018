package main

import "testing"

func TestDay1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"Example 1", []string{"+1", "-1"}, 0},
		{"Example 2", []string{"+3", "+3", "+4", "-2", "-4"}, 10},
		{"Example 3", []string{"-6", "+3", "+8", "+5", "-6"}, 5},
		{"Example 4", []string{"+7", "+7", "-2", "-7", "-4"}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day1(tt.input); got != tt.want {
				t.Errorf("Day1() = %v, want %v", got, tt.want)
			}
		})
	}
}

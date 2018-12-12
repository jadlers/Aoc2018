package main

import (
	"testing"
)

func TestDay12(t *testing.T) {
	type args struct {
		initialState string
		rules        []string
	}
	tests := []struct {
		name   string
		args   args
		wantP1 int
		wantP2 int
	}{
		// TODO: Add test cases.
		{"example 1", args{"#..#.#..##......###...###", []string{
			"...## => #",
			"..#.. => #",
			".#... => #",
			".#.#. => #",
			".#.## => #",
			".##.. => #",
			".#### => #",
			"#.#.# => #",
			"#.### => #",
			"##.#. => #",
			"##.## => #",
			"###.. => #",
			"###.# => #",
			"####. => #",
		}}, 325, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP1, gotP2 := Day12(tt.args.initialState, tt.args.rules)
			if gotP1 != tt.wantP1 {
				t.Errorf("Day12() gotP1 = %v, want %v", gotP1, tt.wantP1)
			}
			if gotP2 != tt.wantP2 {
				t.Errorf("Day12() gotP2 = %v, want %v", gotP2, tt.wantP2)
			}
		})
	}
}

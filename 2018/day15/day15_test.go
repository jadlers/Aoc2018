package main

import "testing"

func TestDay15(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name   string
		args   args
		wantP1 int
		wantP2 int
	}{
		{"Reddit 1", args{[]string{ // Should take 67 rounds
			"####",
			"##E#",
			"#GG#",
			"####",
		}}, 13400, 0},
		{"Reddit 2", args{[]string{ // Should take 71 rounds
			"#####",
			"#GG##",
			"#.###",
			"#..E#",
			"#.#G#",
			"#.E##",
			"#####",
		}}, 13987, 0},
		{"Reddit 3", args{[]string{
			"#######",
			"#.E..G#",
			"#.#####",
			"#G#####",
			"#######",
		}}, 10234, 0},
		{"Example 1", args{[]string{ // Should take 47 rounds
			"#######",
			"#.G...#",
			"#...EG#",
			"#.#.#G#",
			"#..G#E#",
			"#.....#",
			"#######",
		}}, 27730, 0},
		{"Example 2", args{[]string{ // Should take 37 rounds
			"#######",
			"#G..#E#",
			"#E#E.E#",
			"#G.##.#",
			"#...#E#",
			"#...E.#",
			"#######",
		}}, 36334, 0},
		{"Example 3", args{[]string{ // Should take 46 rounds
			"#######",
			"#E..EG#",
			"#.#G.E#",
			"#E.##E#",
			"#G..#.#",
			"#..E#.#",
			"#######",
		}}, 39514, 0},
		{"Example 4", args{[]string{ // Should take 35 rounds
			"#######",
			"#E.G#.#",
			"#.#G..#",
			"#G.#.G#",
			"#G..#.#",
			"#...E.#",
			"#######",
		}}, 27755, 0},
		{"Example 5", args{[]string{ // Should take 54 rounds
			"#######",
			"#.E...#",
			"#.#..G#",
			"#.###.#",
			"#E#G#G#",
			"#...#G#",
			"#######",
		}}, 28944, 0},
		{"Example 5", args{[]string{ // Should take 20 rounds
			"#########",
			"#G......#",
			"#.E.#...#",
			"#..##..G#",
			"#...##..#",
			"#...#...#",
			"#.G...G.#",
			"#.....G.#",
			"#########",
		}}, 18740, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP1, gotP2 := Day15(tt.args.lines)
			if gotP1 != tt.wantP1 {
				t.Errorf("Day15() gotP1 = %v, want %v", gotP1, tt.wantP1)
			}
			if gotP2 != tt.wantP2 {
				t.Errorf("Day15() gotP2 = %v, want %v", gotP2, tt.wantP2)
			}
		})
	}
}

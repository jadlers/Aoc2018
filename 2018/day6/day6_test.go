package main

import "testing"

func TestDay6(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name   string
		args   args
		wantP1 int
		wantP2 int
	}{
		{"Example", args{[]string{"1, 1", "1, 6", "8, 3", "3, 4", "5, 5", "8, 9",}}, 17, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP1, gotP2 := Day6(tt.args.lines)
			if gotP1 != tt.wantP1 {
				t.Errorf("Day6() gotP1 = %v, want %v", gotP1, tt.wantP1)
			}
			if gotP2 != tt.wantP2 {
				t.Errorf("Day6() gotP2 = %v, want %v", gotP2, tt.wantP2)
			}
		})
	}
}

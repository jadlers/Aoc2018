package main

import "testing"

func TestDay9(t *testing.T) {
	type args struct {
		players         int
		lastMarbleValue int
	}
	tests := []struct {
		name   string
		args   args
		wantP1 int
		wantP2 int
	}{
		{"Example 1", args{9, 25}, 32, 22563},
		{"Example 2", args{10, 1618}, 8317, 74765078},
		{"Example 3", args{13, 7999}, 146373, 1406506154},
		{"Example 4", args{17, 1104}, 2764, 20548882},
		{"Example 5", args{21, 6111}, 54718, 507583214},
		{"Example 6", args{30, 5807}, 37305, 320997431},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP1, gotP2 := Day9(tt.args.players, tt.args.lastMarbleValue)
			if gotP1 != tt.wantP1 {
				t.Errorf("Day9() gotP1 = %v, want %v", gotP1, tt.wantP1)
			}
			if gotP2 != tt.wantP2 {
				t.Errorf("Day9() gotP2 = %v, want %v", gotP2, tt.wantP2)
			}
		})
	}
}

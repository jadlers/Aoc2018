package main

import (
	"testing"
)

func TestDay8(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		args   args
		wantP1 int
		wantP2 int
	}{
		{"Example", args{"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"}, 138, 66},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP1, gotP2 := Day8(tt.args.line)
			if gotP1 != tt.wantP1 {
				t.Errorf("Day7() gotP1 = %v, want %v", gotP1, tt.wantP1)
			}
			if gotP2 != tt.wantP2 {
				t.Errorf("Day7() gotP2 = %v, want %v", gotP2, tt.wantP2)
			}
		})
	}
}

package main

import "testing"

func TestDay7(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name   string
		args   args
		wantP1 string
		wantP2 int
	}{
		{"Example", args{[]string{
			"Step C must be finished before step A can begin.",
			"Step C must be finished before step F can begin.",
			"Step A must be finished before step B can begin.",
			"Step A must be finished before step D can begin.",
			"Step B must be finished before step E can begin.",
			"Step D must be finished before step E can begin.",
			"Step F must be finished before step E can begin.",
		}}, "CABDFE", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP1, gotP2 := Day7(tt.args.lines)
			if gotP1 != tt.wantP1 {
				t.Errorf("Day7() gotP1 = '%v', want '%v'", gotP1, tt.wantP1)
			}
			if gotP2 != tt.wantP2 {
				t.Errorf("Day7() gotP2 = %v, want %v", gotP2, tt.wantP2)
			}
		})
	}
}

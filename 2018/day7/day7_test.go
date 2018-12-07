package main

import (
	"testing"
	"fmt"
)

func TestDay7(t *testing.T) {
	const visualRepresentation = `
  -->A--->B--
 /    \      \
C      -->D----->E
 \           /
  ---->F-----`
	fmt.Println(visualRepresentation)

	type args struct {
		lines            []string
		Workers          int
		ExtraTimePerStep int
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
		}, 2, 0}, "CABDFE", 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP1, gotP2 := Day7(tt.args.lines, tt.args.Workers, tt.args.ExtraTimePerStep)
			if gotP1 != tt.wantP1 {
				t.Errorf("Day7() gotP1 = '%v', want '%v'", gotP1, tt.wantP1)
			}
			if gotP2 != tt.wantP2 {
				t.Errorf("Day7() gotP2 = %v, want %v", gotP2, tt.wantP2)
			}
		})
	}
}

func TestInstructionTime(t *testing.T) {
	testChars := "ABF"
	type args struct {
		letter uint8
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"A", args{testChars[0]}, 1},
		{"B", args{testChars[1]}, 2},
		{"F", args{testChars[2]}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InstructionTime(tt.args.letter); got != tt.want {
				t.Errorf("Instructiontime() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func TestDay5(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		args   args
		wantP1 int
		wantP2 int
	}{
		{"Example", args{"dabAcCaCBAcCcaDA"}, 10, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP1, gotP2 := Day5(tt.args.input)
			if gotP1 != tt.wantP1 {
				t.Errorf("Day5() gotP1 = %v, want %v", gotP1, tt.wantP1)
			}
			if gotP2 != tt.wantP2 {
				t.Errorf("Day5() gotP2 = %v, want %v", gotP2, tt.wantP2)
			}
		})
	}
}

func TestIsReacting(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Not same type", args{"a", "b"}, false},
		{"Same polarity", args{"a", "a"}, false},
		{"Same type, different polarity", args{"a", "A"}, true},
		{"Same type, different polarity, reverse", args{"A", "a"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsReacting(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IsReacting() = %v, want %v", got, tt.want)
			}
		})
	}
}

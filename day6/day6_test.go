package main

import (
	"reflect"
	"testing"
)

func TestDay6(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		args   args
		wantP1 int
		wantP2 int
	}{
		{"Example", args{"0 2 7 0"}, 5, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP1, gotP2 := Day6(tt.args.input)
			if gotP1 != tt.wantP1 {
				t.Errorf("Day6() gotP1 = %v, want %v", gotP1, tt.wantP1)
			}
			if gotP2 != tt.wantP2 {
				t.Errorf("Day6() gotP2 = %v, want %v", gotP2, tt.wantP2)
			}
		})
	}
}

func TestRedistribute(t *testing.T) {
	type args struct {
		banks []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{0, 2, 7, 0}}, []int{2, 4, 1, 2}},
		{"2", args{[]int{2, 4, 1, 2}}, []int{3, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Redistribute(tt.args.banks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Redistribute() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"
import "reflect"

func TestDay16_p1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name   string
		args   args
		wantP1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotP1 := Day16_p1(tt.args.input); gotP1 != tt.wantP1 {
				t.Errorf("Day16_p1() = %v, want %v", gotP1, tt.wantP1)
			}
		})
	}
}

func TestEqualOps(t *testing.T) {
	type args struct {
		dump Dump
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"Example 1", args{Dump{
			[4]int{3, 2, 1, 1},
			[]int{9, 2, 1, 2},
			[4]int{3, 2, 2, 1},
		}}, []int{0, 3, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := EqualOps(tt.args.dump); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("NumEqualOps() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

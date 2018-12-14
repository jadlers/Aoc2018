package main

import (
	"reflect"
	"testing"
)

func TestDay14(t *testing.T) {
	type args struct {
		P1noRecipes int
	}
	tests := []struct {
		name   string
		args   args
		wantP1 []int
		wantP2 int
	}{
		{"Example 1", args{9}, []int{5, 1, 5, 8, 9, 1, 6, 7, 7, 9}, 0},
		{"Example 2", args{5}, []int{0, 1, 2, 4, 5, 1, 5, 8, 9, 1}, 0},
		{"Example 3", args{18}, []int{9, 2, 5, 1, 0, 7, 1, 0, 8, 5}, 0},
		{"Example 4", args{2018}, []int{5, 9, 4, 1, 4, 2, 9, 8, 8, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP1, gotP2 := Day14(tt.args.P1noRecipes)
			if !reflect.DeepEqual(gotP1, tt.wantP1) {
				t.Errorf("Day14() gotP1 = %v, want %v", gotP1, tt.wantP1)
			}
			if gotP2 != tt.wantP2 {
				t.Errorf("Day14() gotP2 = %v, want %v", gotP2, tt.wantP2)
			}
		})
	}
}

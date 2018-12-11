package main

import (
	"reflect"
	"testing"
)

func TestDay11(t *testing.T) {
	type args struct {
		serialNumber int
	}
	tests := []struct {
		name   string
		args   args
		wantP1 []int
		wantP2 int
	}{
		{"Example 1", args{18}, []int{33, 45}, 0},
		{"Example 2", args{42}, []int{21, 61}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP1, gotP2 := Day11(tt.args.serialNumber)
			if !reflect.DeepEqual(gotP1, tt.wantP1) {
				t.Errorf("Day11() gotP1 = %v, want %v", gotP1, tt.wantP1)
			}
			if gotP2 != tt.wantP2 {
				t.Errorf("Day11() gotP2 = %v, want %v", gotP2, tt.wantP2)
			}
		})
	}
}

func TestCalcPowerLevel(t *testing.T) {
	type args struct {
		x            int
		y            int
		serialNumber int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{122, 79, 57}, -5},
		{"example 2", args{217, 196, 39}, 0},
		{"example 3", args{101, 153, 71}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcPowerLevel(tt.args.x, tt.args.y, tt.args.serialNumber); got != tt.want {
				t.Errorf("CalcPowerLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

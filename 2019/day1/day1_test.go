package main

import (
	"testing"
)

func Test_calcFuel(t *testing.T) {
	type args struct {
		weight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{12}, 2},
		{"", args{14}, 2},
		{"", args{1969}, 654},
		{"", args{100756}, 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFuel(tt.args.weight); got != tt.want {
				t.Errorf("calcFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcFuel2(t *testing.T) {
	type args struct {
		weight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{14}, 2},
		{"", args{1969}, 966},
		{"", args{100756}, 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFuel2(tt.args.weight); got != tt.want {
				t.Errorf("calcFuel2() = %v, want %v", got, tt.want)
			}
		})
	}
}

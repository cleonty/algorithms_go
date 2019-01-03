package main

import (
	"testing"
)

func Test_findSumClosestToZero(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"zero", args{[]int{-1, 0, 1}}, 0},
		{"zero again", args{[]int{-1, 0, 1, 2}}, 0},
		{"one", args{[]int{-2, 0, 1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSumClosestToZero(tt.args.data); got != tt.want {
				t.Errorf("findSumClosestToZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"positive", args{1}, 1},
		{"positive", args{11}, 11},
		{"negative", args{-1}, 1},
		{"negative", args{-11}, 11},
		{"zero", args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.a); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

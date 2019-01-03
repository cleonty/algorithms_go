package bsearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		data  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"found 1", args{[]int{1, 2, 3, 4, 5}, 1}, true},
		{"found 2", args{[]int{1, 2, 3, 4, 5}, 2}, true},
		{"found 3", args{[]int{1, 2, 3, 4, 5}, 3}, true},
		{"found 4", args{[]int{1, 2, 3, 4, 5}, 4}, true},
		{"found 5", args{[]int{1, 2, 3, 4, 5}, 5}, true},
		{"not found 6", args{[]int{1, 2, 3, 4, 5}, 6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.data, tt.args.value); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	type args struct {
		data  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"found 1", args{[]int{1, 2, 3, 4, 5}, 1}, true},
		{"found 2", args{[]int{1, 2, 3, 4, 5}, 2}, true},
		{"found 3", args{[]int{1, 2, 3, 4, 5}, 3}, true},
		{"found 4", args{[]int{1, 2, 3, 4, 5}, 4}, true},
		{"found 5", args{[]int{1, 2, 3, 4, 5}, 5}, true},
		{"not found 6", args{[]int{1, 2, 3, 4, 5}, 6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchRecursive(tt.args.data, tt.args.value); got != tt.want {
				t.Errorf("BinarySearchRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

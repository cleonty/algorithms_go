package bitoniclist

import (
	"testing"
)

func Test_findMaxima(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{"middle", args{[]int{0, 1, 0}}, 1, true},
		{"two", args{[]int{-1, 0, 1, 0}}, 2, true},
		{"three", args{[]int{-2, -1, 0, 1, 0}}, 3, true},
		{"one", args{[]int{0, 1, 0, -1}}, 1, true},
		{"one", args{[]int{0, 1, 0, -1, -2}}, 1, true},
		{"two elements", args{[]int{0, 1}}, -1, false},
		{"three equal elements", args{[]int{1, 1, 1}}, -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findMaxima(tt.args.data)
			if got != tt.want {
				t.Errorf("findMaxima() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findMaxima() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		data []int
		val  int
		inc  bool
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{"found inc 3", args{[]int{0, 1, 2, 3, 4}, 3, true}, 3, true},
		{"found inc 1", args{[]int{0, 1, 2, 3, 4}, 1, true}, 1, true},
		{"found dec 1", args{[]int{4, 3, 2, 1, 0}, 1, false}, 3, true},
		{"found dec 3", args{[]int{4, 3, 2, 1, 0}, 3, false}, 1, true},
		{"too low", args{[]int{4, 3, 2, 1, 0}, -1, false}, -1, false},
		{"too hi", args{[]int{4, 3, 2, 1, 0}, 5, false}, -1, false},
		{"empty", args{[]int{}, 3, false}, -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := binarySearch(tt.args.data, tt.args.val, tt.args.inc)
			if got != tt.want {
				t.Errorf("binarySearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("binarySearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_findElement(t *testing.T) {
	type args struct {
		data []int
		val  int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{"middle", args{[]int{0, 1, 0}, 1}, 1, true},
		{"two", args{[]int{-1, 0, 1, 0}, 1}, 2, true},
		{"three", args{[]int{-2, -1, 0, 1, 0}, 1}, 3, true},
		{"-1", args{[]int{-2, -1, 0, 1, 0, -5}, -1}, 1, true},
		{"-5", args{[]int{-2, -1, 0, 1, 0, -5}, -5}, 5, true},
		{"-6", args{[]int{-2, -1, 0, 1, 0, -5}, -6}, -1, false},
		{"one", args{[]int{0, 1, 0, -1}, 1}, 1, true},
		{"one", args{[]int{0, 1, 0, -1, -2}, 1}, 1, true},
		{"two elements", args{[]int{0, 1}, 1}, -1, false},
		{"three equal elements", args{[]int{1, 1, 1}, 1}, -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findElement(tt.args.data, tt.args.val)
			if got != tt.want {
				t.Errorf("findElement() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findElement() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
